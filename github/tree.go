package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TreeResponse struct {
	Tree []TreeItem `json:"tree"`
}

type TreeItem struct {
	Path string `json:"path"`
	Type string `json:"type"` // "blob" or "tree"
}

func GetRepoTree(owner, repo string) ([]TreeItem, error) {
	// Get default branch SHA
	branchURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	resp, err := http.Get(branchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		DefaultBranch string `json:"default_branch"`
	}
	json.NewDecoder(resp.Body).Decode(&data)

	shaURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1", owner, repo, data.DefaultBranch)
	treeResp, err := http.Get(shaURL)
	if err != nil {
		return nil, err
	}
	defer treeResp.Body.Close()

	var tree TreeResponse
	err = json.NewDecoder(treeResp.Body).Decode(&tree)
	if err != nil {
		return nil, err
	}

	return tree.Tree, nil
}
