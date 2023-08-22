/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "github.com/fatih/color"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
	// "log"
	// "os"
)

// nodepoolsCmd represents the nodepools command
var nodepoolsCmd = &cobra.Command{
	Use:   "list-nodepools",
	Aliases: []string{"pools"},
	Short: "Lists nodepools in a cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		project := cmd.Flags().Lookup("project").Value.String()
		zone := cmd.Flags().Lookup("zone").Value.String()
		cluster := cmd.Flags().Lookup("cluster").Value.String()

		parser.ListNodepools(project, zone, cluster)
		// if len(project) == 0 {
		// 	color.Red("\nNo project specified!\n\n")
		// 	cmd.Help()
		// 	os.Exit(0)
		// } else if len(cluster) == 0 {
		// 	_, err := parser.ListPoolsUsingPrompt(project, zone)
		// 	if err != nil {
		// 		log.Printf("Error listing nodepools: %v", err)
		// 	}
		// } else {
		// 	_, err := parser.ListNodepools(project, zone, cluster)
		// 	if err != nil {
		// 		log.Printf("Error listing nodepools: %v", err)
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(nodepoolsCmd)
	nodepoolsCmd.Flags().String("project", "", "A project to list in")
	nodepoolsCmd.Flags().String("zone", "", "A zone to list in")
	nodepoolsCmd.Flags().String("cluster", "", "(Optional) A cluster to list in.\nIf missing, you will be prompted")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodepoolsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodepoolsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
