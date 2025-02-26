package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

func (repository *Repository) Commit(message string) {
	duration := time.Duration(1) * time.Second
	time.Sleep(duration)

	id := uuid.New().String()
	date := time.Now().Format("2006-01-02 15:04:05")
	newCommit := &Commit{Id: id, Message: message, Date: date}
	repository.ActiveBranch.Commits = append(repository.ActiveBranch.Commits, newCommit)

	reversedCommits, err := ReverseCommitsOrder(repository.ActiveBranch.Commits)
	if err != nil {
		log.Error(fmt.Sprintf("Error reversing commits order: %v", err))
		return
	}

	repository.ActiveBranch.Commits = reversedCommits
	repository.ActiveBranch.LastCommit = newCommit
	repository.Head = repository.ActiveBranch.LastCommit

	log.Info(fmt.Sprintf("Commit %s created at %s", id, date))
}
