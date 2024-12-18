package main

import git "github.com/abroudoux/got/internal/git"

type Repository = git.Repository

func main() {
	var repository Repository = Repository{}
	repository.InitRepository("my-repo")
	repository.Commit("Initial commit")
	repository.Commit("Second commit")
	repository.Log()
}
