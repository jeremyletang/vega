// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.DATANODE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package sqlstore

import (
	"context"
	"fmt"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/datanode/metrics"
	v2 "code.vegaprotocol.io/vega/protos/data-node/api/v2"
	"github.com/georgysavva/scany/pgxscan"
)

type Rewards struct {
	*ConnectionSource
}

var rewardsOrdering = TableOrdering{
	ColumnOrdering{Name: "epoch_id", Sorting: ASC},
}

func NewRewards(connectionSource *ConnectionSource) *Rewards {
	r := &Rewards{
		ConnectionSource: connectionSource,
	}
	return r
}

func (rs *Rewards) Add(ctx context.Context, r entities.Reward) error {
	defer metrics.StartSQLQuery("Rewards", "Add")()
	_, err := rs.Connection.Exec(ctx,
		`INSERT INTO rewards(
			party_id,
			asset_id,
			market_id,
			reward_type,
			epoch_id,
			amount,
			percent_of_total,
			timestamp,
			tx_hash,
			vega_time,
			seq_num)
		 VALUES ($1,  $2,  $3,  $4,  $5,  $6, $7, $8, $9, $10, $11);`,
		r.PartyID, r.AssetID, r.MarketID, r.RewardType, r.EpochID, r.Amount, r.PercentOfTotal, r.Timestamp, r.TxHash,
		r.VegaTime, r.SeqNum)
	return err
}

func (rs *Rewards) GetAll(ctx context.Context) ([]entities.Reward, error) {
	defer metrics.StartSQLQuery("Rewards", "GetAll")()
	rewards := []entities.Reward{}
	err := pgxscan.Select(ctx, rs.Connection, &rewards, `
		SELECT * from rewards;`)
	return rewards, err
}

func (rs *Rewards) GetByCursor(ctx context.Context,
	partyIDHex *string,
	assetIDHex *string,
	pagination entities.CursorPagination,
) ([]entities.Reward, entities.PageInfo, error) {
	var pageInfo entities.PageInfo
	query, args := selectRewards(partyIDHex, assetIDHex)

	query, args, err := PaginateQuery[entities.RewardCursor](query, args, rewardsOrdering, pagination)
	if err != nil {
		return nil, pageInfo, err
	}

	rewards := []entities.Reward{}
	if err := pgxscan.Select(ctx, rs.Connection, &rewards, query, args...); err != nil {
		return nil, entities.PageInfo{}, fmt.Errorf("querying rewards: %w", err)
	}

	rewards, pageInfo = entities.PageEntities[*v2.RewardEdge](rewards, pagination)
	return rewards, pageInfo, nil
}

func (rs *Rewards) GetByOffset(ctx context.Context,
	partyIDHex *string,
	assetIDHex *string,
	pagination *entities.OffsetPagination,
) ([]entities.Reward, error) {
	query, args := selectRewards(partyIDHex, assetIDHex)

	if pagination != nil {
		orderCols := []string{"epoch_id", "party_id", "asset_id"}
		query, args = orderAndPaginateQuery(query, orderCols, *pagination, args...)
	}

	rewards := []entities.Reward{}
	defer metrics.StartSQLQuery("Rewards", "Get")()
	if err := pgxscan.Select(ctx, rs.Connection, &rewards, query, args...); err != nil {
		return nil, fmt.Errorf("querying rewards: %w", err)
	}
	return rewards, nil
}

func selectRewards(partyIDHex, assetIDHex *string) (string, []interface{}) {
	query := `SELECT * from rewards`
	args := []interface{}{}
	addRewardWhereClause(&query, &args, partyIDHex, assetIDHex)
	return query, args
}

func (rs *Rewards) GetSummaries(ctx context.Context,
	partyIDHex *string, assetIDHex *string,
) ([]entities.RewardSummary, error) {
	query := `SELECT party_id, asset_id, sum(amount) as amount FROM rewards`
	args := []interface{}{}
	addRewardWhereClause(&query, &args, partyIDHex, assetIDHex)
	query = fmt.Sprintf("%s GROUP BY party_id, asset_id", query)

	summaries := []entities.RewardSummary{}
	defer metrics.StartSQLQuery("Rewards", "GetSummaries")()
	err := pgxscan.Select(ctx, rs.Connection, &summaries, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying rewards: %w", err)
	}
	return summaries, nil
}

// -------------------------------------------- Utility Methods

func addRewardWhereClause(queryPtr *string, args *[]interface{}, partyIDHex, assetIDHex *string) {
	query := *queryPtr
	if partyIDHex != nil && *partyIDHex != "" {
		partyID := entities.PartyID(*partyIDHex)
		query = fmt.Sprintf("%s WHERE party_id=%s", query, nextBindVar(args, partyID))
	}

	if assetIDHex != nil && *assetIDHex != "" {
		clause := "WHERE"
		if partyIDHex != nil {
			clause = "AND"
		}

		assetID := entities.AssetID(*assetIDHex)
		query = fmt.Sprintf("%s %s asset_id=%s", query, clause, nextBindVar(args, assetID))
	}
	*queryPtr = query
}
