package cmd

import (
	"github.com/devopsgig/arithproducer/src/arithproducer"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the arithproducer",
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

// Starts the server
func run(cmd *cobra.Command, args []string) {
	arithproducer.Run()
}
