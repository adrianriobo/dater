package cmd

import (
	"fmt"

	"github.com/adrianriobo/dater/pkg/dater/xunit"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	xunitStatusCmdName string = "status"

	fileUrl  string = "file-url"
	filePath string = "file-path"
)

func init() {
	xunitCmd.AddCommand(xunitStatusCmd)
	flagSet := pflag.NewFlagSet(xunitStatusCmdName, pflag.ExitOnError)
	flagSet.StringP(fileUrl, "u", "", "url to download the xunit file")
	flagSet.StringP(filePath, "p", "", "file path to load the xunit file")
	xunitStatusCmd.Flags().AddFlagSet(flagSet)
}

var xunitStatusCmd = &cobra.Command{
	Use:   xunitStatusCmdName,
	Short: "global status for test results",
	Long:  "check global status for test results based on an xunit file: SUCCESS or FAIL",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		status, err := xunit.GlobalStatusRemote(
			viper.GetString(fileUrl),
			viper.GetString(filePath))
		if err == nil {
			fmt.Println(status)
		}
		return err
	},
}
