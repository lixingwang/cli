package commands

import (
	"fmt"
	"os"

	"github.com/project-flogo/cli/api"
	"github.com/project-flogo/cli/common"
	"github.com/spf13/cobra"
)

var localContrib bool
var installCmd = &cobra.Command{
	Use:   "install [flags] <contribution>",
	Short: "install a flogo contribution",
	Long:  "Installs a flogo contribution",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.InstallPackage(common.CurrentProject(), args, localContrib)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	installCmd.Flags().BoolVarP(&localContrib, "localContrib", "l", false, "Specify local Contrib")
	rootCmd.AddCommand(installCmd)

}
