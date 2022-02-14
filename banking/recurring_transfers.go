package banking

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"
)

var (
	ErrStartEpochInThePast                                     = errors.New("start epoch in the past")
	ErrCannotSubmitDuplicateRecurringTransferWithSameFromAndTo = errors.New("cannot submit duplicate recurring transfer with same from an to")
)

func (e *Engine) recurringTransfer(
	ctx context.Context,
	transfer *types.RecurringTransfer,
) (err error) {
	defer func() {
		if err != nil {
			e.bss.changed[recurringTransfersKey] = true
		}
		e.broker.Send(events.NewRecurringTransferFundsEvent(ctx, transfer))
	}()

	// ensure asset exists
	a, err := e.assets.Get(transfer.Asset)
	if err != nil {
		transfer.Status = types.TransferStatusRejected
		e.log.Debug("cannot transfer funds, invalid asset", logging.Error(err))
		return fmt.Errorf("could not transfer funds: %w", err)
	}

	if err := transfer.IsValid(); err != nil {
		transfer.Status = types.TransferStatusRejected
		return err
	}

	if err := e.ensureMinimalTransferAmount(a, transfer.Amount); err != nil {
		transfer.Status = types.TransferStatusRejected
		return err
	}

	if err := e.ensureNoRecurringTransferDuplicates(transfer); err != nil {
		transfer.Status = types.TransferStatusRejected
		return err
	}

	// can't create transfer with start epoch in the past
	if transfer.StartEpoch <= e.currentEpoch {
		transfer.Status = types.TransferStatusRejected
		return ErrStartEpochInThePast
	}

	// from here all sounds OK, we can add the transfer
	// in the recurringTransfer map
	e.recurringTransfers[transfer.ID] = transfer

	return nil
}

func (e *Engine) ensureNoRecurringTransferDuplicates(
	transfer *types.RecurringTransfer,
) error {
	for _, v := range e.recurringTransfers {
		if v.From == transfer.From && v.To == transfer.To && v.FromAccountType == transfer.FromAccountType && v.ToAccountType == transfer.ToAccountType {
			return ErrCannotSubmitDuplicateRecurringTransferWithSameFromAndTo
		}
	}

	return nil
}

func (e *Engine) distributeRecurringTransfers(
	ctx context.Context,
	newEpoch uint64,
) error {
	var (
		transfersDone = []events.Event{}
		tresps        = []*types.TransferResponse{}
		currentEpoch  = num.NewUint(newEpoch).ToDecimal()
	)

	allIDs := make([]string, 0, len(e.recurringTransfers))
	for k := range e.recurringTransfers {
		allIDs = append(allIDs, k)
	}

	sort.SliceStable(allIDs, func(i, j int) bool { return allIDs[i] < allIDs[j] })

	// iterate over all transfers
	for _, k := range allIDs {
		v := e.recurringTransfers[k]
		if v.StartEpoch > newEpoch {
			// not started
			continue
		}

		var (
			startEpoch  = num.NewUint(v.StartEpoch).ToDecimal()
			startAmount = v.Amount.ToDecimal()
			amount, _   = num.UintFromDecimal(
				startAmount.Mul(
					v.Factor.Pow(currentEpoch.Sub(startEpoch)),
				),
			)
		)

		// check if the amount is still enough
		// ensure asset exists
		a, err := e.assets.Get(v.Asset)
		if err != nil {
			// this should not be possible, asset was validated at first when
			// accepting the transfer
			e.log.Panic("this should never happen", logging.Error(err))
		}

		if err := e.ensureMinimalTransferAmount(a, amount); err != nil {
			v.Status = types.TransferStatusStopped
			transfersDone = append(transfersDone,
				events.NewRecurringTransferFundsEvent(ctx, v))
			delete(e.recurringTransfers, k)
			continue
		}

		resps, err := e.processTransfer(
			ctx, v.From, v.To, v.Asset, v.FromAccountType, v.ToAccountType, amount, v.Reference, nil, // last is eventual oneoff, which this is not
		)
		if err != nil {
			v.Status = types.TransferStatusStopped
			transfersDone = append(transfersDone,
				events.NewRecurringTransferFundsEvent(ctx, v))
			delete(e.recurringTransfers, k)
			continue
		}

		tresps = append(tresps, resps...)

		// if we don't have anymore
		if v.EndEpoch != nil && *v.EndEpoch == e.currentEpoch {
			v.Status = types.TransferStatusDone
			transfersDone = append(transfersDone, events.NewRecurringTransferFundsEvent(ctx, v))
			delete(e.recurringTransfers, k)
		}
	}

	// send events
	if len(tresps) > 0 {
		e.broker.Send(events.NewTransferResponse(ctx, tresps))
	}
	if len(transfersDone) > 0 {
		// also set the state change
		e.bss.changed[recurringTransfersKey] = true
		e.broker.SendBatch(transfersDone)
	}

	return nil
}
