package main

import (
	"fmt"
	"slices"

	"github.com/charmbracelet/log"
)

func (r *Repository) Branch(branchName string) {
	branch := &Branch{
		Name:       branchName,
		Commits:    slices.Clone(r.ActiveBranch.Commits),
		LastCommit: r.Head,
	}
	r.Branches = append(r.Branches, branch)

	log.Info(fmt.Sprintf("Branch %s created at commit %s", branch.Name, branch.LastCommit.Id))
}

func (r *Repository) LogBranches() {
	if len(r.Branches) == 0 {
		log.Info("No branches created yet.")
		return
	}

	log.Info("Branches:")
	for _, branch := range r.Branches {
		log.Info(fmt.Sprintf("Branch %s created at commit %s", branch.Name, branch.LastCommit.Id))
	}
}
