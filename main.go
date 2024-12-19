package main

import (
	git "github.com/abroudoux/got/internal/git"
)

type Repository = git.Repository

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
}
