package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cassis",
	Short: "CAsual SSI System",
	Long:  "CAsual SSI System",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err) // TODO remove ( for debug )
		//fmt.Fprintln(os.Stderr, err)
		//os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
