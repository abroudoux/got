package main

import (
	git "github.com/abroudoux/got/internal/git"
)

type Repository = git.Repository

func main() {
	var repository Repository = Repository{}
	repository.Init("my-repo")
	repository.Commit("Initial commit")
	repository.Commit("Second commit")
	repository.Commit("Third commit")
	repository.PrintHead()
	repository.Log()
}
