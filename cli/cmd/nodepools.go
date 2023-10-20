package cmd

import (
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
)

var nodepoolsCmd = &cobra.Command{
	Use:     "list-nodepools",
	Aliases: []string{"pools"},
	Short:   "Lists nodepools in a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		project := cmd.Flags().Lookup("project").Value.String()
		zone := cmd.Flags().Lookup("zone").Value.String()
		cluster := cmd.Flags().Lookup("cluster").Value.String()

		if cluster != "" {
			parser.ListNodepools(project, zone, cluster)
		} else {
			parser.ListPoolsUsingPrompt(project, zone)
		}
	},
}

func init() {
	rootCmd.AddCommand(nodepoolsCmd)
	nodepoolsCmd.Flags().String("project", "", "A project to list in")
	nodepoolsCmd.Flags().String("zone", "", "A zone to list in")
	nodepoolsCmd.Flags().String("cluster", "", "(Optional) A cluster to list in.\nIf missing, you will be prompted")
}
