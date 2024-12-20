package main

import (
	git "github.com/abroudoux/got/internal/git"
)

type Repository = git.Repository
type RemoteRepository = git.RemoteRepository

func main() {
	var repository Repository = Repository{}
	repository.Init("my-repo")
	repository.Debug()
	println("--------------------------------")
	repository.Commit("commit-1")
	repository.Debug()
	println("--------------------------------")
	repository.Branch("branch-1")
	repository.Debug()
	println("--------------------------------")
	repository.Checkout("branch-1")
	repository.Debug()
	println("--------------------------------")
	repository.Commit("commit-2")
	repository.Debug()
	println("--------------------------------")
	repository.Checkout("main")
	repository.Debug()
	println("--------------------------------")
	repository.Merge("branch-1")
	repository.Debug()
	println("--------------------------------")
	var remoteRepository RemoteRepository = RemoteRepository{}
	remoteRepository.CreateRemoteRepository("my-remote-repo")
	println("--------------------------------")
	repository.RemoteAdd(remoteRepository)
	repository.Debug()
	println("--------------------------------")
	repository.Push()
	repository.Debug()
	remoteRepository.Debug()
}
