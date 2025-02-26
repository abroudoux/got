package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

func (r *Repository) Commit(message string) {
	duration := time.Duration(1) * time.Second
	time.Sleep(duration)

	id := uuid.New().String()
	date := time.Now().Format("2006-01-02 15:04:05")
	newCommit := &Commit{Id: id, Message: message, Date: date}
	r.ActiveBranch.Commits = append(r.ActiveBranch.Commits, newCommit)

	reversedCommits, err := ReverseCommitsOrder(r.ActiveBranch.Commits)
	if err != nil {
		log.Error(fmt.Sprintf("Error reversing commits order: %v", err))
		return
	}

	r.ActiveBranch.Commits = reversedCommits
	r.ActiveBranch.LastCommit = newCommit
	r.Head = r.ActiveBranch.LastCommit

	log.Info(fmt.Sprintf("Commit %s created at %s", id, date))
}

func (r *Repository) LogCommits() {
	if len(r.ActiveBranch.Commits) == 0 {
		log.Info("No commits yet.")
		return
	}

	log.Info("Commits:")
	for _, commit := range r.ActiveBranch.Commits {
		log.Info(fmt.Sprintf("Commit %s: %s", commit.Id, commit.Message))
	}
}
