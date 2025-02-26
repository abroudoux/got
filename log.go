package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (repository *Repository) Log() {
	if len(repository.ActiveBranch.Commits) == 0 {
		log.Warn("No commits yet")
		return
	}

	for _, commit := range repository.ActiveBranch.Commits {
		log.Debug(fmt.Sprintf("%s: %s -- %s \n", commit.Id, commit.Message, commit.Date))
	}
}
