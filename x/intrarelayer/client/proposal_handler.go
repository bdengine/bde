package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/tharsis/bdeos/x/intrarelayer/client/cli"
	"github.com/tharsis/bdeos/x/intrarelayer/client/rest"
)

var (
	RegisterCoinProposalHandler         = govclient.NewProposalHandler(cli.NewRegisterCoinProposalCmd, rest.RegisterCoinProposalRESTHandler)
	RegisterERC20ProposalHandler        = govclient.NewProposalHandler(cli.NewRegisterERC20ProposalCmd, rest.RegisterERC20ProposalRESTHandler)
	ToggleTokenRelayProposalHandler     = govclient.NewProposalHandler(cli.NewToggleTokenRelayProposalCmd, rest.ToggleTokenRelayRESTHandler)
	UpdateTokenPairERC20ProposalHandler = govclient.NewProposalHandler(cli.NewUpdateTokenPairERC20ProposalCmd, rest.UpdateTokenPairERC20ProposalRESTHandler)
)
