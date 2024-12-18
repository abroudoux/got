package main

import (
	git "github.com/abroudoux/got/internal/git"
)

type Repository = git.Repository

func main() {
	var repository Repository = Repository{}
	repository.Init("my-repo")
	repository.PrintBranches()
	repository.PrintHead()
	println("--------------------------------")
	repository.Commit("commit-1")
	repository.PrintBranches()
	repository.PrintHead()
	println("--------------------------------")
	repository.Branch("branch-1")
	repository.PrintBranches()
	repository.PrintHead()
	println("--------------------------------")
	repository.Checkout("branch-1")
	repository.Commit("commit-2")
	repository.PrintBranches()
	repository.PrintHead()
	println("--------------------------------")
	repository.Checkout("main")
	repository.PrintBranches()
	repository.PrintHead()
}
