package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func Init(repositoryName string) *LocalRepository {
	var localRepository = &LocalRepository{}

	localRepository.Repository.Name = repositoryName
	mainBranch := &Branch{Name: "main", Commits: []*Commit{}, LastCommit: nil}
	localRepository.Repository.Branches = append(localRepository.Repository.Branches, mainBranch)
	localRepository.ActiveBranch = mainBranch
	localRepository.Head = localRepository.ActiveBranch.LastCommit

	log.Info(fmt.Sprintf("Repository %s initialized.", repositoryName))
	return localRepository
}
