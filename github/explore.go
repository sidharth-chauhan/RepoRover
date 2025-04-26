package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RepoItem struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Stargazers  int    `json:"stargazers_count"`
	Owner       struct {
		Login string `json:"login"`
	} `json:"owner"`
}

type SearchResponse struct {
	Items []RepoItem `json:"items"`
}

func ExploreRepos(language string) ([]RepoItem, error) {
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&sort=stars&order=desc&per_page=10", language)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Items, nil
}
