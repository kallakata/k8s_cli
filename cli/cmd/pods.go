
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/fatih/color"
	"os"
)

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:     "list-pods",
	Aliases: []string{"pods"},
	Short:   "Lists Pods in context",
	Run: func(cmd *cobra.Command, args []string) {
		// ns := cmd.Flags().Lookup("namespace").Value.String()
		ctx := cmd.Flags().Lookup("context").Value.String()
		if len(ctx) == 0 {
			color.Red("\nNo context specified!\n\n")
            cmd.Help()
            os.Exit(0)
        } else {
		parser.ListPodsUsingPrompt(ctx)
		}
		// parser.ListPods(ns, ctx)
		// if ctx != "" && ns != "" {
		// 	parser.ListPods(ns, ctx)
		// 	} else if ns == "" {
		// 		parser.ListPods("", ctx)
		// 	} else {
		// 		fmt.Printf("Please specify a context")
		// 	}
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
}
