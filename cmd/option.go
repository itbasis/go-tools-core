package cmd

import (
	itbasisCoreOption "github.com/itbasis/go-tools-core/v1/option"
	"github.com/spf13/cobra"
)

type Option = itbasisCoreOption.Option[cobra.Command]
