package cmd

import (
	"fmt"
	"strings"

	"github.com/sidharth-chauhan/reporover/github"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var exploreCmd = &cobra.Command{
	Use:   "explore <language>",
	Short: "Explore top GitHub repositories by programming language",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		language := strings.ToLower(args[0])
		repos, err := github.ExploreRepos(language)
		if err != nil {
			color.Red("❌ Failed to fetch trending repos: %v", err)
			return
		}

		color.Cyan("\n🔥 Top Repositories for language: %s\n", language)
		for _, repo := range repos {
			color.Yellow("⭐ %s (%d stars)", repo.FullName, repo.Stargazers)
			fmt.Printf("🔗 %s\n", repo.HTMLURL)
			color.Green("📦 %s\n\n", repo.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(exploreCmd)
}
