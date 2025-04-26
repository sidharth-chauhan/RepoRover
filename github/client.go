	package github

	import (
	    "encoding/json"
	    "fmt"
	    "net/http"
	    "sort"
	    "time"
	)

	// Repo represents a GitHub repository
	type Repo struct {
	    Name            string    `json:"name"`
	    Description     string    `json:"description"`
	    HTMLURL         string    `json:"html_url"`
	    StargazersCount int       `json:"stargazers_count"`
	    ForksCount      int       `json:"forks_count"`
	    UpdatedAt       time.Time `json:"updated_at"`
	    Language        string    `json:"language"`
	}

	// fetchReposFromGitHub makes the API call to GitHub to fetch repositories
	func fetchReposFromGitHub(username string) ([]Repo, error) {
	    // Create HTTP client with timeout
	    client := &http.Client{
	        Timeout: 10 * time.Second,
	    }

	    // Build the GitHub API URL
	    url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	    
	    // Make the HTTP request
	    resp, err := client.Get(url)
	    if err != nil {
	        return nil, fmt.Errorf("error making request to GitHub API: %v", err)
	    }
	    defer resp.Body.Close()

	    // Check for non-200 status code
	    if resp.StatusCode != http.StatusOK {
	        return nil, fmt.Errorf("GitHub API returned status code %d", resp.StatusCode)
	    }

	    // Decode the response
	    var repos []Repo
	    err = json.NewDecoder(resp.Body).Decode(&repos)
	    if err != nil {
	        return nil, fmt.Errorf("error decoding GitHub API response: %v", err)
	    }

	    return repos, nil
	}

	// GetRepos fetches repositories for a given GitHub username with filtering and sorting options
	func GetRepos(username string, sortBy string, language string, minStars int) ([]Repo, error) {
	    // Fetch repositories from GitHub
	    repos, err := fetchReposFromGitHub(username)
	    if err != nil {
	        return nil, err
	    }

	    // Filter repositories by language and minimum stars
	    var filtered []Repo
	    for _, repo := range repos {
	        if language != "" && repo.Language != language {
	            continue
	        }
	        if minStars > 0 && repo.StargazersCount < minStars {
	            continue
	        }
	        filtered = append(filtered, repo)
	    }

	    // Sort repositories based on the specified criterion
	    switch sortBy {
	    case "stars":
	        sort.Slice(filtered, func(i, j int) bool {
	            return filtered[i].StargazersCount > filtered[j].StargazersCount
	        })
	    case "forks":
	        sort.Slice(filtered, func(i, j int) bool {
	            return filtered[i].ForksCount > filtered[j].ForksCount
	        })
	    case "updated":
	        sort.Slice(filtered, func(i, j int) bool {
	            return filtered[i].UpdatedAt.After(filtered[j].UpdatedAt)
	        })
	    }

	    return filtered, nil
	}

