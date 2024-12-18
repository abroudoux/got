package git

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Repository struct {
	Name string
	Commits []Commit
	Head Commit
	Branches []Branch
}

type Commit struct {
	Id string
	Message string
	Date string
}

type Branch struct {
	Name string
	Commit Commit
}

func (repository *Repository) Init(repositoryName string) {
	repository.Name = repositoryName
	repository.Commits = []Commit{}
	masterBranch := Branch{Name: "master", Commit: Commit{}}
	repository.Branches = append(repository.Branches, masterBranch)
	repository.Head = masterBranch.Commit
}

func (repository *Repository) Commit(message string) {
	id := uuid.New().String()
	date := time.Now().Format("2006-01-02 15:04:05")
	newCommit := Commit{Id: id, Message: message, Date: date}
	repository.Commits = append(repository.Commits, newCommit)
	repository.Head = newCommit
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
		fmt.Printf("%s: %s -- %s \n", commit.Id, commit.Message, commit.Date)
	}
}

func reverseCommitsOrder(commits []Commit) ([]Commit, error) {
	var reversedCommits []Commit
	for i := len(commits) - 1; i >= 0; i-- {
		reversedCommits = append(reversedCommits, commits[i])
	}

	return reversedCommits, nil
}

func (repository *Repository) Branch(branchName string) {
	branch := Branch{Name: branchName, Commit: repository.Head}
	fmt.Printf("Branch %s created at commit %s\n", branch.Name, branch.Commit.Id)
}