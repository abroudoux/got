package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func Init(repositoryName string) *Repository {
	var repository = &Repository{}

	repository.Name = repositoryName
	mainBranch := &Branch{Name: "main", Commits: []*Commit{}, LastCommit: nil}
	repository.Branches = append(repository.Branches, mainBranch)
	repository.ActiveBranch = mainBranch
	repository.Head = repository.ActiveBranch.LastCommit

	log.Info(fmt.Sprintf("Repository %s initialized.", repositoryName))
	return repository
}
