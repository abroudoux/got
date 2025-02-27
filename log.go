package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Log() {
	if len(r.ActiveBranch.Commits) == 0 {
		log.Warn("No commits yet.")
		return
	}

	log.Info("Commits:")
	for _, commit := range r.ActiveBranch.Commits {
		log.Info(fmt.Sprintf("	%s: %s", commit.Date, RenderEl(commit.Message)))
	}

	return
}

func (r *RemoteRepository) Log() {
	if len(r.DefaultBranch.Commits) == 0 {
		log.Warn("No commits yet.")
		return
	}

	log.Info("Commits:")
	for _, commit := range r.DefaultBranch.Commits {
		log.Info(fmt.Sprintf("	%s: %s", commit.Date, RenderEl(commit.Message)))
	}

	return
}

func (r *RemoteRepository) LogRemoteRepository() {
	log.Info(fmt.Sprintf("Remote repository %s:", RenderEl(r.Name)))
	log.Info("Default branch:")
	log.Info(fmt.Sprintf("	%s", RenderEl(r.DefaultBranch.Name)))

	log.Info("Commits:")
	for _, branch := range r.Repository.Branches {
		log.Info(fmt.Sprintf("	%s:", RenderEl(branch.Name)))
		for _, commit := range branch.Commits {
			log.Info(fmt.Sprintf("		%s: %s", commit.Date, RenderEl(commit.Message)))
		}
	}
}
