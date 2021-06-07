//lint:file-ignore ST1003 Ignore underscores in names, this is straigh copied from the proto package to ease introducing the domain types

package types

import (
	"code.vegaprotocol.io/vega/proto"
	commandspb "code.vegaprotocol.io/vega/proto/commands/v1"
)

type LiquidityMonitoringParameters = proto.LiquidityMonitoringParameters
type LiquidityProvisionSubmission = commandspb.LiquidityProvisionSubmission
type LiquidityProvision = proto.LiquidityProvision
type LiquidityOrder = proto.LiquidityOrder
type LiquidityOrderReference = proto.LiquidityOrderReference

type LiquidityProvision_Status = proto.LiquidityProvision_Status

const (
	// The default value
	LiquidityProvision_STATUS_UNSPECIFIED LiquidityProvision_Status = 0
	// The liquidity provision is active
	LiquidityProvision_STATUS_ACTIVE LiquidityProvision_Status = 1
	// The liquidity provision was stopped by the network
	LiquidityProvision_STATUS_STOPPED LiquidityProvision_Status = 2
	// The liquidity provision was cancelled by the liquidity provider
	LiquidityProvision_STATUS_CANCELLED LiquidityProvision_Status = 3
	// The liquidity provision was invalid and got rejected
	LiquidityProvision_STATUS_REJECTED LiquidityProvision_Status = 4
	// The liquidity provision is valid and accepted by network, but orders aren't deployed
	LiquidityProvision_STATUS_UNDEPLOYED LiquidityProvision_Status = 5
	// The liquidity provision is valid and accepted by network
	// but have never been deployed. I when it's possible to deploy them for the first time
	// margin check fails, then they will be cancelled without any penalties.
	LiquidityProvision_STATUS_PENDING LiquidityProvision_Status = 6
)
