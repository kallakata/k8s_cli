package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/internal/interfaces"
	"github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/kallakata/k8s_cli/pretty/pretty_clusters" // Import the pretty_clusters package
	"github.com/spf13/cobra"
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
			clusters, err := parser.ListClusters(project, "-")
			if err != nil {
				log.Printf("Error listing clusters: %v", err)
				os.Exit(1)
			}
			runPrettyClusters(clusters)
		} else {
			clusters, err := parser.ListClusters(project, zone)
			if err != nil {
				log.Printf("Error listing clusters: %v", err)
				os.Exit(1)
			}
			runPrettyClusters(clusters)
		}
	},
}

var NodePoolsFetcherInstance interfaces.NodePoolsFetcher

func init() {
	rootCmd.AddCommand(clustersCmd)
	clustersCmd.Flags().String("project_id", "", "A project to list in")
	clustersCmd.Flags().String("zone", "", "A zone to list in")
}

func runPrettyClusters(clusters []model.Cluster) {
	model := pretty_clusters.NewModel(clusters)
	p := tea.NewProgram(model)
	if err, _ := p.Run(); err != nil {
		log.Printf("Error starting Bubbletea program: %v", err)
		os.Exit(1)
	}
}
