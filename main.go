package main

import (
	git "github.com/abroudoux/got/internal/git"
)

type Repository = git.Repository
type RemoteRepository = git.RemoteRepository

func main() {
	var repository = &Repository{}
	repository.Init("my-repo")
	repository.Debug()
	repository.Commit("commit-1")
	repository.Debug()
	repository.Branch("branch-1")
	repository.Debug()
	repository.Checkout("branch-1")
	repository.Debug()
	repository.Commit("commit-2")
	repository.Debug()
	repository.Checkout("main")
	repository.Debug()
	repository.Merge("branch-1")
	repository.Debug()
	var remoteRepository = &RemoteRepository{}
	remoteRepository.CreateRemoteRepository("my-remote-repo")
	remoteRepository.Debug()
	repository.RemoteAdd(remoteRepository)
	repository.Debug()
	repository.Push()
	repository.Debug()
	remoteRepository.Debug()
}
