package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestCLIFlagsToEnvVars(t *testing.T) {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "test",
			EnvVars: []string{"NODE_TEST_VAR"},
		},
		&cli.IntFlag{
			Name: "no env var",
		},
	}
	res := cliFlagsToEnvVars(flags)
	require.Contains(t, res, "NODE_TEST_VAR")
}

func TestValidateEnvVars(t *testing.T) {
	provided := []string{"BATCHER_CONFIG=true", "BATCHER_FAKE=false", "LD_PRELOAD=/lib/fake.so"}
	defined := map[string]struct{}{
		"BATCHER_CONFIG": {},
		"BATCHER_OTHER":  {},
	}
	invalids := validateEnvVars("BATCHER", provided, defined)
	require.ElementsMatch(t, invalids, []string{"BATCHER_FAKE=false"})
}
