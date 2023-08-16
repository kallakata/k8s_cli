
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kallakata/k8s_cli/parser"
	// "github.com/kallakata/k8s_cli/prompt"
	// tea "github.com/charmbracelet/bubbletea"
	// "log"
)

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:     "list-pods",
	Aliases: []string{"pods"},
	Short:   "Lists pods",
	Run: func(cmd *cobra.Command, args []string) {
		ns := cmd.Flags().Lookup("namespace").Value.String()
		ctx := cmd.Flags().Lookup("context").Value.String()

		// var nsx prompt.Namespace
		// p := tea.NewProgram(prompt.InitialModel())
		// 	if _, err := p.Run(); err != nil {
		// 	log.Fatal(err)
		// }
		parser.ListPods(ns, ctx)
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
	podsCmd.Flags().String("namespace", "", "A namespace to list in")
}
