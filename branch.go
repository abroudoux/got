package main

import (
	"fmt"
	"slices"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Branch(branchName string) {
	branch := &Branch{
		Name:    branchName,
		Commits: slices.Clone(r.ActiveBranch.Commits),
	}
	r.Repository.Branches = append(r.Repository.Branches, branch)

	log.Info(fmt.Sprintf("Branch %s created at commit %s", RenderEl(branch.Name), RenderEl(branch.Commits[0].Message)))
	return
}

func (r *LocalRepository) LogBranches() {
	if len(r.Repository.Branches) == 0 {
		log.Info("No branches created yet.")
		return
	}

	log.Info("Branches:")
	for _, branch := range r.Repository.Branches {
		log.Info(fmt.Sprintf("	%s", RenderEl(branch.Name)))
	}
	return
}

func (r *RemoteRepository) LogBranches() {
	if r.Repository == nil || len(r.Repository.Branches) == 0 {
		log.Info(fmt.Sprintf("No branches in remote repository %s", r.Name))
		return
	}

	log.Info(fmt.Sprintf("Branches remote repository %s: ", r.Name))
	for _, branch := range r.Repository.Branches {
		log.Info(fmt.Sprintf("  %s", branch.Name))
	}

	return
}
