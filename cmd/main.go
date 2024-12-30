package main

import (
	git "github.com/abroudoux/got/internal/git"
)

type Repository = git.Repository
type RemoteRepository = git.RemoteRepository

func main() {
	repo := git.Init("my-repo")
	repo.Commit("commit-1")
	remoteRepo := git.InitRemoteRepository("my-remote-repo")
	repo.RemoteAdd(remoteRepo)
	println(repo.Origin.Url)
}
