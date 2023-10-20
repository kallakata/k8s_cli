package cmd

import (
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
)

var namespacesCmd = &cobra.Command{
	Use:     "list-namespaces",
	Aliases: []string{"ns"},
	Short:   "List namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Flags().Lookup("context").Value.String()

		if len(ctx) == 0 {
			color.Red("\nNo context specified!\nUsing current context.\n\n")
			parser.ListNamespaces(ctx)
		} else {
			parser.ListNamespaces(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(namespacesCmd)
	namespacesCmd.Flags().String("context", "", "A context to list in")
}
