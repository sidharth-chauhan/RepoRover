package cmd

import (
	"fmt"

	"github.com/sidharth-chauhan/reporover/github"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile [username]",
	Short: "Get GitHub user profile information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		user, err := github.GetUser(username)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("👤 %s (%s)\n", user.Name, user.Login)
		fmt.Printf("📜 Bio: %s\n", user.Bio)
		fmt.Printf("🏢 Company: %s\n", user.Company)
		fmt.Printf("📍 Location: %s\n", user.Location)
		fmt.Printf("📦 Public Repos: %d\n", user.PublicRepos)
		fmt.Printf("👥 Followers: %d | Following: %d\n", user.Followers, user.Following)
		fmt.Printf("🔗 Profile: %s\n", user.HTMLURL)
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
