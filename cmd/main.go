package main

import (
	git "github.com/abroudoux/got/internal/git"
)

type Repository = git.Repository
type RemoteRepository = git.RemoteRepository

func main() {
	// var repository = &Repository{}
	// repository.Init("my-repo")
	// repository.Commit("commit-1")
	// repository.Commit("commit-2")
	// repository.Log()
	// var remoteRepository = &RemoteRepository{}
	// remoteRepository.CreateRemoteRepository("my-remote-repo")
	// remoteRepository.Debug()
	// repository.RemoteAdd(remoteRepository)
	// remoteRepository.Debug()
	// repository.Push()
	// remoteRepository.Debug()

	repo := git.Init("my-repo")
	println(repo.Name)
	println(repo.ActiveBranch.Name)
	
	for _, branch := range repo.Branches {
		println(branch.Name)
		for _, commit := range branch.Commits {
			println(commit.Id)
		}
	}
}
