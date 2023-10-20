package cmd

import (
	"fmt"
	"github.com/kallakata/k8s_cli/parser"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var azsCmd = &cobra.Command{
	Use:     "list-azs-cluster",
	Aliases: []string{"azs-cluster"},
	Short:   "Lists clusters in Azure",
	Run: func(cmd *cobra.Command, args []string) {
		subscription := cmd.Flags().Lookup("subscription").Value.String()
		if len(subscription) == 0 {
			fmt.Printf("Subscription ID missing!\n\n")
			cmd.Help()
			os.Exit(0)
		} else {
			err, _ := parser.ListAzsClusters(subscription)
			if err != nil {
				log.Printf("Error listing clusters: %v", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(azsCmd)
	azsCmd.Flags().String("subscription", "", "A subscription to list in")
}
