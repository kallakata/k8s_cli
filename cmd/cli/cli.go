package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/kallakata/k8s_cli/auth"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kubeparser",
	Short: "A parser to list nodes and pods",
	Long: `A parser that lets you list pods and nodes in a context and namespace, along with help functionality and
	native local authentication via kubeconfig`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Welcome to Kubernetes parser!")
	},
}

func init() {
    rootCmd.AddCommand(nsCmd)
	rootCmd.AddCommand(podsCmd)
	nsCmd.Flags().String("context", "", "A context to list in")
	podsCmd.Flags().String("context", "", "A context to list in")
	podsCmd.Flags().String("namespace", "", "A namespace to list in")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
    // if err := rootCmd.ParseFlags(os.Args[1:]); err != nil {
    //     fmt.Fprintf(os.Stderr, "Error parsing flags: %s\n", err)
    //     os.Exit(1)
    // }
    
    // if err := rootCmd.Execute(); err != nil {
    //     fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'\n", err)
    //     os.Exit(1)
    // }
}

type Clientset struct {

}

var nsCmd = &cobra.Command{
	Use:     "list-namespaces",
	Aliases: []string{"ns"},
	Short:   "List namespaces",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := args[0]

		if ctx != "" {
			auth.Auth(ctx)
			ListNamespaces(ctx)
		} else {
			fmt.Printf("Please specify a context")
		}
	},
}

var podsCmd = &cobra.Command{
	Use:     "list-pods",
	Aliases: []string{"pods"},
	Short:   "Lists pods",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ns, _ := cmd.Flags().GetString("namespace")
		ctx, _ := cmd.Flags().GetString("context")

		if ctx != "" {
			auth.Auth(args[0])
			if ns != "" {
				ListPods(ns)
			} else {
				fmt.Printf("Please specify a namespace")
			}
		} else {
			fmt.Printf("Please specify a context")
		}
	},
}
