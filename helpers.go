package main

import "fmt"

func ReverseCommitsOrder(commits []*Commit) ([]*Commit, error) {
	var reversedCommits []*Commit
	for i := len(commits) - 1; i >= 0; i-- {
		reversedCommits = append(reversedCommits, commits[i])
	}

	return reversedCommits, nil
}

func RenderEl(element string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", "38;2;214;112;214", element)
}
