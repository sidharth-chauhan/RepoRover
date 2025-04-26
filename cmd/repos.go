package cmd

import (
	"fmt"

	"github.com/sidharth-chauhan/reporover/github"
	"github.com/spf13/cobra"
	
)

var sortBy, language string
var minStars int
var reposCmd = &cobra.Command{
	Use:   "repos [username]",
	Short: "List public repos of a GitHub user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		repos, err := github.GetRepos(username, sortBy, language, minStars)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, repo := range repos {
			fmt.Printf(" %s\n    %s\n    %s\n\n", repo.Name, repo.Description, repo.HTMLURL)
		}
	},
}

func init() {
	rootCmd.AddCommand(reposCmd)
	reposCmd.Flags().StringVar(&sortBy, "sort", "", "Sort by: stars|forks|updated")
	reposCmd.Flags().StringVar(&language, "language", "", "Filter by language")
	reposCmd.Flags().IntVar(&minStars, "min-stars", 0, "Filter by minimum stars")
}
