package cmd

import (
	"fmt"
	"os"

	"github.com/aqyuki/soreppoi/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "soreppoi",
	Short: "It displays a screen that looks like that.",
	Long:  `You can use it in front of your friends who are PC novices and be a little proud of yourself!`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.PrintSorrepoiStrings(time)
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	},
}

// time is a variable that stores the startup time.
var time int32

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Int32VarP(&time, "time", "t", 5, "Specify the startup time. (in seconds) Default is 5 seconds")
}
