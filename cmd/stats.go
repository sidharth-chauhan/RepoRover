package cmd

import (
	"fmt"
	"strings"

	"github.com/sidharth-chauhan/reporover/github"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats <owner>/<repo>",
	Short: "Get GitHub repository statistics",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parts := strings.Split(args[0], "/")
		if len(parts) != 2 {
			color.Red("❌ Invalid format. Use <owner>/<repo>")
			return
		}

		owner, repo := parts[0], parts[1]
		stats, err := github.FetchRepoStats(owner, repo)
		if err != nil {
			color.Red("❌ Error fetching stats: %v", err)
			return
		}

		color.Cyan("\n📊 Stats for: %s", stats.FullName)
		fmt.Println("-------------------------------------")
		color.Yellow("📦 Description: %s", stats.Description)
		color.Green("⭐ Stars: %d", stats.StargazersCount)
		color.Green("🍴 Forks: %d", stats.ForksCount)
		color.Green("👀 Watchers: %d", stats.WatchersCount)
		color.Red("🐛 Open Issues: %d", stats.OpenIssuesCount)
		color.Magenta("🌲 Default Branch: %s", stats.DefaultBranch)
		color.Cyan("📅 Last Updated: %s", stats.UpdatedAt.Format("Jan 2, 2006 3:04PM"))
		color.Blue("🔗 URL: %s\n", stats.HTMLURL)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
