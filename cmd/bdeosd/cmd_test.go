package main_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cmdcfg "github.com/tharsis/bdeos/cmd/config"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/tharsis/bdeos/app"
	bdeosd "github.com/tharsis/bdeos/cmd/bdeosd"
)

func TestInitCmd(t *testing.T) {

	rootCmd, _ := bdeosd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"bde-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "bdeos_9000-1"),
	})

	err := svrcmd.Execute(rootCmd, app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestStartCmd(t *testing.T) {
	setupConfig()
	cmdcfg.RegisterDenoms()
	rootCmd, _ := bdeosd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"start",
	})

	err := svrcmd.Execute(rootCmd, app.DefaultNodeHome)
	require.NoError(t, err)
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	if err := cmdcfg.EnableObservability(); err != nil {
		panic(err)
	}
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}