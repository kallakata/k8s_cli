package cmd

import (
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
    "github.com/kallakata/k8s_cli/pretty/pretty_clusters" // Import the pretty_clusters package
    "github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/internal/interfaces"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

var clustersCmd = &cobra.Command{
	Use:     "list-clusters",
	Aliases: []string{"clusters"},
	Short:   "List clusters in project and zone",
    Run: func(cmd *cobra.Command, args []string) {
        project := cmd.Flags().Lookup("project_id").Value.String()
        zone := cmd.Flags().Lookup("zone").Value.String()

        if len(project) == 0 {
            color.Red("\nNo project specified!\n\n")
            cmd.Help()
            os.Exit(0)
        } else if len(zone) == 0 {
            clusters, err := parser.ListClusters(project, "-", NodePoolsFetcherInstance)
            if err != nil {
                log.Printf("Error listing clusters: %v", err)
                os.Exit(1)
            }
            runPrettyClusters(clusters, NodePoolsFetcherInstance)
        } else {
            clusters, err := parser.ListClusters(project, zone, NodePoolsFetcherInstance)
            if err != nil {
                log.Printf("Error listing clusters: %v", err)
                os.Exit(1)
            }
            runPrettyClusters(clusters, NodePoolsFetcherInstance)
        }
    },
}

var NodePoolsFetcherInstance interfaces.NodePoolsFetcher

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

func runPrettyClusters(clusters []model.Cluster, npf interfaces.NodePoolsFetcher) {
    model := pretty_clusters.NewModel(clusters, npf)
    p := tea.NewProgram(model)
    if err, _ := p.Run(); err != nil {
        log.Printf("Error starting Bubbletea program: %v", err)
        os.Exit(1)
    }
}

