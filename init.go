package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func Init(repositoryName string) *LocalRepository {
	if repositoryName == "" {
		log.Error("Repository name is empty.")
		return nil
	}

	mainBranch := &Branch{Name: "main", Commits: []*Commit{}}
	var localRepository = &LocalRepository{
		Name: repositoryName,
		Repository: &Repository{
			Branches: []*Branch{mainBranch},
		},
		ActiveBranch: mainBranch,
		Head:         nil,
	}

	log.Info(fmt.Sprintf("Repository %s initialized.", RenderEl(repositoryName)))
	return localRepository
}

func (r *LocalRepository) LogName() {
	log.Info("Repository name:")
	log.Info(fmt.Sprintf("  %s", RenderEl(r.Name)))
	return
}
