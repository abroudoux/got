package git

import (
	"fmt"

	"github.com/google/uuid"
)

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
		fmt.Println("No commits yet")
		return
	}

	reversedCommits, err := reverseCommitsOrder(repository.Commits)
	if err != nil {
		fmt.Println("Error reversing commits order")
		return
	}

	for _, commit := range reversedCommits {
		fmt.Printf("%s: %s \n", commit.Id, commit.Message)
	}

	for _, commit := range reversedCommits {
		fmt.Printf("%s: %s \n", commit.Id, commit.Message)
	}
}

func reverseCommitsOrder(commits []Commit) ([]Commit, error) {
	var reversedCommits []Commit
	for i := len(commits) - 1; i >= 0; i-- {
		reversedCommits = append(reversedCommits, commits[i])
	}

	return reversedCommits, nil
}