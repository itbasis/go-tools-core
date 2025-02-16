package cmd

import (
	"log/slog"

	itbasisCoreLog "github.com/itbasis/go-tools-core/v1/log"
	itbasisCoreOs "github.com/itbasis/go-tools-core/v1/os"
	"github.com/spf13/cobra"
)

func LogCommand(cmd *cobra.Command, args []string) {
	slog.Debug(
		"execute command",
		itbasisCoreLog.SlogAttrCommand(
			cmd.Name(),
			itbasisCoreOs.Pwd(),
			args,
		),
	)
}
