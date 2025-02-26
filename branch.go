package main

import (
	"fmt"
	"slices"
)

func (r *Repository) Branch(branchName string) {
	branch := &Branch{
		Name:       branchName,
		Commits:    slices.Clone(r.ActiveBranch.Commits),
		LastCommit: r.Head,
	}
	r.Branches = append(r.Branches, branch)

	fmt.Printf("Branch %s created at commit %s\n", branch.Name, branch.LastCommit.Id)
}
