package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func CreateRemoteRepository(remoteRepositoryName string) *RemoteRepository {
	var remoteRepository = &RemoteRepository{}

	remoteRepository.Name = remoteRepositoryName
	remoteRepository.Url = fmt.Sprintf("https://gothub.com/%s", remoteRepository.Name)
	remoteRepository.Repository = &Repository{}

	log.Info(fmt.Sprintf("Remote repository %s created at %s", remoteRepository.Name, remoteRepository.Url))
	return remoteRepository
}

func (r *Repository) RemoteAdd(remoteRepository *RemoteRepository) {
	r.Origin = remoteRepository.Url
	log.Info(fmt.Sprintf("Remote %s added", remoteRepository.Url))
}

func (r *RemoteRepository) LogBranches() {
	if r.Repository == nil || len(r.Repository.Branches) == 0 {
		log.Info(fmt.Sprintf("No branches in remote repository %s", r.Name))
		return
	}

	log.Info(fmt.Sprintf("Branches remote repository %s: ", r.Name))
	for _, branch := range r.Repository.Branches {
		log.Info(fmt.Sprintf("  %s", branch.Name))
	}
}

func (r *RemoteRepository) LogCommits() {
	if r.Repository == nil || r.Repository.ActiveBranch == nil || len(r.Repository.ActiveBranch.Commits) == 0 {
		log.Info(fmt.Sprintf("No commits in remote repository %s", r.Name))
		return
	}

	log.Info(fmt.Sprintf("Commits remote repository %s: ", r.Name))
	for _, commit := range r.Repository.ActiveBranch.Commits {
		log.Info(fmt.Sprintf("  %s", commit.Message))
	}
}
