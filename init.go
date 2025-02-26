package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func Init(repositoryName string) *LocalRepository {
	var repository = &Repository{}
	var localRepository = &LocalRepository{}

	localRepository.Name = repositoryName
	mainBranch := &Branch{Name: "main", Commits: []*Commit{}}
	localRepository.Repository = repository
	localRepository.Repository.Branches = append(localRepository.Repository.Branches, mainBranch)
	localRepository.ActiveBranch = mainBranch
	localRepository.Head = nil

	log.Info(fmt.Sprintf("Repository %s initialized.", RenderEl(repositoryName)))
	return localRepository
}

func (r *LocalRepository) LogName() {
	log.Info("Repository name:")
	log.Info(fmt.Sprintf("  %s", RenderEl(r.Name)))
}
