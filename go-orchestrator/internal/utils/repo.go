package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
)

// CloneRepo clones a GitHub repo locally into /tmp and returns the path.
func CloneRepo(repoURL string) (string, error) {
	dir, err := os.MkdirTemp("", "repo-*")
	if err != nil {
		return "", err
	}

	fmt.Println("Cloning repo:", repoURL)
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	if err != nil {
		os.RemoveAll(dir)
		return "", err
	}

	return dir, nil
}

// SummarizeRepo scans the repo and returns simple metadata.
func SummarizeRepo(repoPath string) map[string]interface{} {
	summary := map[string]interface{}{
		"total_files":   0,
		"languages":     map[string]int{},
		"total_lines":   0,
		"largest_files": []string{},
	}

	fileCount := 0
	lines := 0
	languages := map[string]int{}

	filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		fileCount++

		ext := strings.TrimPrefix(filepath.Ext(info.Name()), ".")
		if ext != "" {
			languages[ext]++
		}

		return nil
	})

	summary["total_files"] = fileCount
	summary["languages"] = languages
	summary["total_lines"] = lines
	return summary
}
