package query

import (
	"fmt"

	apipb "code.vegaprotocol.io/protos/vega/api/v1"

	"github.com/golang/protobuf/jsonpb"
)

type MarketsCmd struct {
	NodeAddress string `long:"node-address" description:"The address of the vega node to use" default:"0.0.0.0:3002"`
}

func (opts *MarketsCmd) Execute(_ []string) error {
	req := apipb.ListMarketsRequest{}
	return getPrintMarkets(opts.NodeAddress, &req)
}

func getPrintMarkets(nodeAddress string, req *apipb.ListMarketsRequest) error {
	clt, err := getClient(nodeAddress)
	if err != nil {
		return fmt.Errorf("could not connect to the vega node: %w", err)
	}

	ctx, cancel := timeoutContext()
	defer cancel()
	res, err := clt.ListMarkets(ctx, req)
	if err != nil {
		return fmt.Errorf("error querying the vega node: %w", err)
	}

	m := jsonpb.Marshaler{
		Indent: "  ",
	}
	buf, err := m.MarshalToString(res)
	if err != nil {
		return fmt.Errorf("invalid response from vega node: %w", err)
	}

	fmt.Printf("%v", buf)

	return nil
}