package main

import (
	"github.com/google/uuid"
)

func main() {
	var repository Repository = Repository{}
	repository.InitRepository("my-repo")
	repository.Commit("Initial commit")
	repository.Commit("Second commit")
	repository.Log()
}

type Repository struct {
	Name string
	Commits []Commit
	Head Commit
}

type Commit struct {
	Id string
	Message string
}

func (repository *Repository) InitRepository(repositoryName string) {
	repository.Name = repositoryName
}

func (repository *Repository) Commit(message string) {
	id := uuid.New().String()
	commit := Commit{Id: id, Message: message}
	repository.Commits = append(repository.Commits, commit)
	repository.Head = commit
}

func (repository *Repository) Log() {
	if len(repository.Commits) == 0 {
		println("No commits yet")
		return
	}

	for _, commit := range repository.Commits {
		println(commit.Message)
	}
}