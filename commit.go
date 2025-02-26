package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (repository *Repository) Commit(message string) {
	id := uuid.New().String()
	date := time.Now().Format("2006-01-02 15:04:05")
	newCommit := &Commit{Id: id, Message: message, Date: date}
	repository.ActiveBranch.Commits = append(repository.ActiveBranch.Commits, newCommit)

	reversedCommits, err := ReverseCommitsOrder(repository.ActiveBranch.Commits)
	if err != nil {
		fmt.Println("Error reversing commits order")
		return
	}

	repository.ActiveBranch.Commits = reversedCommits
	repository.ActiveBranch.LastCommit = newCommit
	repository.Head = repository.ActiveBranch.LastCommit

	fmt.Printf("Commit %s created at %s\n", id, date)
}
