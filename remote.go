package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func CreateRemoteRepository(remoteRepositoryName string) *RemoteRepository {
	if remoteRepositoryName == "" {
		log.Error("Remote repository name is empty.")
		return nil
	}

	var remoteRepository = &RemoteRepository{
		Name:          remoteRepositoryName,
		Url:           fmt.Sprintf("https://gothub.com/%s", remoteRepositoryName),
		DefaultBranch: &Branch{},
		Repository:    &Repository{},
	}

	log.Info(fmt.Sprintf("Remote repository %s created at %s", remoteRepository.Name, remoteRepository.Url))
	return remoteRepository
}

func (r *LocalRepository) RemoteAdd(remoteRepository *RemoteRepository) {
	r.Origin = remoteRepository
	log.Info(fmt.Sprintf("Remote %s added", remoteRepository.Url))
	return
}
