package cmd

import (
	"os"
	"log"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

var clustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Aliases: []string{"clusters"},
	Short: "List clusters in project and zone",
	Run: func(cmd *cobra.Command, args []string) {
		project := cmd.Flags().Lookup("project_id").Value.String()
		zone := cmd.Flags().Lookup("zone").Value.String()
        if len(project) == 0 {
            color.Red("\nNo project specified!\n\n")
            cmd.Help()
            os.Exit(0)
        } else if len(zone) == 0 {
            _, err := parser.ListClusters(project, "-")
            if err != nil {
                log.Printf("Error listing clusters: %v", err)
            }
        } else {
            _, err := parser.ListClusters(project, zone)
            if err != nil {
                log.Printf("Error listing clusters: %v", err)
            }
        }
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
