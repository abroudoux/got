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
