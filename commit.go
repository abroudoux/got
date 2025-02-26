package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Commit(message string) {
	duration := time.Duration(1) * time.Second
	time.Sleep(duration)

	date := time.Now().Format("2006-01-02 15:04:05")
	newCommit := &Commit{Message: message, Date: date}
	r.ActiveBranch.Commits = append(r.ActiveBranch.Commits, newCommit)

	reversedCommits, err := ReverseCommitsOrder(r.ActiveBranch.Commits)
	if err != nil {
		log.Error(fmt.Sprintf("Error reversing commits order: %v", err))
		return
	}

	r.ActiveBranch.Commits = reversedCommits
	r.Head = r.ActiveBranch.Commits[0]

	log.Info(fmt.Sprintf("Commit %s created at %s", RenderEl(message), RenderEl(date)))
	return
}

func (r *LocalRepository) LogCommits() {
	if len(r.ActiveBranch.Commits) == 0 {
		log.Info("No commits yet.")
		return
	}

	log.Info("Commits:")
	for _, commit := range r.ActiveBranch.Commits {
		log.Info(fmt.Sprintf("	%s: %s", RenderEl(commit.Date), RenderEl(commit.Message)))
	}

	return
}
