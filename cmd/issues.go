package cmd

import (
	"fmt"
	"strings"

	"github.com/sidharth-chauhan/reporover/github"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var label string

var issuesCmd = &cobra.Command{
	Use:   "issues [owner/repo]",
	Short: "View repository issues",
	Long:  `Display open issues for a GitHub repository.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parts := strings.Split(args[0], "/")
		if len(parts) != 2 {
			fmt.Println("Please provide the repository in the format: owner/repo")
			return
		}

		issues, err := github.GetIssues(parts[0], parts[1], label)
		if err != nil {
			fmt.Printf("Error fetching issues: %v\n", err)
			return
		}

		if len(issues) == 0 {
			fmt.Println("No issues found.")
			return
		}

		for _, issue := range issues {
			color.Green("🔖 %s", issue.Title)
			fmt.Printf("   by @%s\n   %s\n\n", issue.User.Login, issue.HTMLURL)
		}
	},
}

func init() {
	issuesCmd.Flags().StringVarP(&label, "label", "l", "good first issue", "Label to filter issues")
	rootCmd.AddCommand(issuesCmd)
}
