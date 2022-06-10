package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	xunitCmdName string = "xunit"
)

func init() {
	rootCmd.AddCommand(xunitCmd)
	flagSet := pflag.NewFlagSet(xunitCmdName, pflag.ExitOnError)
	xunitCmd.Flags().AddFlagSet(flagSet)
}

var xunitCmd = &cobra.Command{
	Use:   xunitCmdName,
	Short: "xunit operations",
	Long:  "xunit operations",
}
