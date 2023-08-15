/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "fmt"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var clustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Aliases: []string{"cs"},
	Short: "List clusters in context",
	Run: func(cmd *cobra.Command, args []string) {
		project := cmd.Flags().Lookup("project").Value.String()
		zone := cmd.Flags().Lookup("zone").Value.String()

		// if project == "" || zone == "" {
		// 	fmt.Printf("Please specify a project and zone")
		// } else {
		// parser.ListClusters(project, zone)
		// }
		parser.ListClusters(project, zone)
	},
}

func init() {
	rootCmd.AddCommand(clustersCmd)
	clustersCmd.Flags().String("project_id", "", "A project to list in")
	clustersCmd.Flags().String("zone", "", "A zone to list in")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clustersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clustersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
