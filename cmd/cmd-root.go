package cmd

import (
	"context"

	itbasisCoreOption "github.com/itbasis/go-tools-core/option"

	"github.com/spf13/cobra"
)

func InitDefaultCmdRoot(ctx context.Context, shortName string, opts ...Option) (*cobra.Command, error) {
	var cmd = &cobra.Command{Short: shortName}

	if err := itbasisCoreOption.ApplyOptions(
		ctx,
		cmd,
		opts, map[itbasisCoreOption.Key]itbasisCoreOption.LazyOptionFunc[cobra.Command]{
			_optionVersionKey:   WithDefaultVersion,
			_optionOutKey:       WithDefaultOut,
			_optionErrKey:       WithDefaultErr,
			_optionDebugFlagKey: WithDefaultFlagDebug,
		},
	); err != nil {
		return nil, err //nolint:wrapcheck // TODO
	}

	return cmd, nil
}
