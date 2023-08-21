package cmd

import (
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
	"os"
)

var namespacesCmd = &cobra.Command{
	Use:     "list-namespaces",
	Aliases: []string{"ns"},
	Short:   "List namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Flags().Lookup("context").Value.String()

		if len(ctx) == 0 {
			color.Red("\nNo context specified!\n\n")
			cmd.Help()
			os.Exit(0)
		} else {
			parser.ListNamespaces(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(namespacesCmd)
	namespacesCmd.Flags().String("context", "", "A context to list in")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namespacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// namespacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
