package main

import (
	git "github.com/abroudoux/got/internal/git"
)

type Repository = git.Repository
type RemoteRepository = git.RemoteRepository

func main() {
	repo := git.Init("my-repo")
	repo.LogBranches()
	repo.Commit("commit-1")
	repo.LogCommits()
	repo.Branch("branch-1")
	repo.Checkout("branch-1")
	repo.Commit("commit-2")
	repo.LogCommits()
	repo.Checkout("main")
	repo.LogCommits()
	repo.Merge("branch-1")
	repo.LogCommits()
}
