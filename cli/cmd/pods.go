package cmd

import (
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
)

var podsCmd = &cobra.Command{
	Use:     "list-pods",
	Aliases: []string{"pods"},
	Short:   "Lists Pods in context",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Flags().Lookup("context").Value.String()
		ns := cmd.Flags().Lookup("namespace").Value.String()

		switch {
		case len(ctx) == 0 && len(ns) == 0:
			color.Red("\nNo context specified!\nUsing current context.\n\n")
			parser.ListPodsUsingPrompt("")
		case len(ns) != 0 && len(ctx) != 0:
			parser.ListPods(ns, ctx)
		case len(ns) != 0 && len(ctx) == 0:
			color.Red("\nNo context specified!\nUsing current context.\n\n")
			parser.ListPods(ns, "")
		case len(ns) == 0 && len(ctx) != 0:
			parser.ListPodsUsingPrompt(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(podsCmd)
	podsCmd.Flags().String("context", "", "(Optional) A context to list in.\nIf missing, a default one will be used.")
	podsCmd.Flags().String("namespace", "", "(Optional) A namespace to list in.\nIf missing, you will be prompted")
}
