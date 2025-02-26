package main

import "fmt"

func (remoteRepository *RemoteRepository) CreateRemoteRepository(repositoryName string) {
	remoteRepository.Name = repositoryName
	remoteRepository.Url = fmt.Sprintf("https://gothub.com/%s", remoteRepository.Name)
	remoteRepository.Repository = &Repository{}

	fmt.Printf("Remote repository %s created at %s\n", remoteRepository.Name, remoteRepository.Url)
}

func (repository *Repository) RemoteAdd(remoteRepository *RemoteRepository) {
	repository.Origin = remoteRepository
	fmt.Printf("Remote %s added\n", remoteRepository.Url)
}
