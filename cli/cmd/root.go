package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kubeparser",
	Short: "A parser to list nodes and pods",
	Long: `A parser that lets you list pods and nodes in a context and namespace, along with help functionality and
	native local authentication via kubeconfig`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
