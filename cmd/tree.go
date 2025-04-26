package cmd

import (
	"fmt"
	"strings"

	"github.com/sidharth-chauhan/reporover/github"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree [owner/repo]",
	Short: "View the file tree of a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parts := strings.Split(args[0], "/")
		if len(parts) != 2 {
			fmt.Println("Use format: owner/repo")
			return
		}
		owner, repo := parts[0], parts[1]

		tree, err := github.GetRepoTree(owner, repo)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, item := range tree {
			if item.Type == "tree" {
				fmt.Printf(" %s/\n", item.Path)
			} else {
				fmt.Printf(" %s\n", item.Path)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}
