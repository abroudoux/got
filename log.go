package main

import "fmt"

func (repository *Repository) Log() {
	if len(repository.ActiveBranch.Commits) == 0 {
		fmt.Println("No commits yet")
		return
	}

	for _, commit := range repository.ActiveBranch.Commits {
		fmt.Printf("%s: %s -- %s \n", commit.Id, commit.Message, commit.Date)
	}
}
