package cmd

import (
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
)

// podsCmd represents the pods command
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
				parser.ListPods(ns, "")
			case len(ns) == 0 && len(ctx) != 0:
				parser.ListPodsUsingPrompt(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(podsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	podsCmd.Flags().String("context", "", "A context to list in")
	podsCmd.Flags().String("namespace", "", "(Optional) A namespace to list in.\nIf missing, you will be prompted")
}
