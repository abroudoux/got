package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func CreateRemoteRepository(repositoryName string) *RemoteRepository {
	var remoteRepository = &RemoteRepository{}

	remoteRepository.Name = repositoryName
	remoteRepository.Url = fmt.Sprintf("https://gothub.com/%s", remoteRepository.Name)
	remoteRepository.Repository = &Repository{}

	log.Info(fmt.Printf("Remote repository %s created at %s\n", remoteRepository.Name, remoteRepository.Url))
	return remoteRepository
}

func (repository *Repository) RemoteAdd(remoteRepository *RemoteRepository) {
	repository.Origin = remoteRepository
	remoteRepository.Repository = repository
	log.Info(fmt.Sprintf("Remote %s added", remoteRepository.Url))
}
