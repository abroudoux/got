package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Commit(message string) {
	if len(message) == 0 {
		log.Error("Commit message is empty.")
		return
	}

	duration := time.Duration(1) * time.Second
	time.Sleep(duration)

	date := time.Now().Format("2006-01-02 15:04:05")
	newCommit := &Commit{Message: message, Date: date}
	r.ActiveBranch.Commits = append([]*Commit{newCommit}, r.ActiveBranch.Commits...)
	r.Head = r.ActiveBranch.Commits[0]

	log.Info(fmt.Sprintf("Commit %s created at %s", RenderEl(message), date))
	return
}
