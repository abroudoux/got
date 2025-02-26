package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Merge(branchName string) {
	if r.ActiveBranch.Name == branchName {
		log.Error(fmt.Sprintf("Cannot merge branch %s into itself", RenderEl(branchName)))
		return
	}

	for _, branch := range r.Repository.Branches {
		if branch.Name == branchName {
			existingCommits := make(map[*Commit]bool)
			for _, commit := range r.ActiveBranch.Commits {
				existingCommits[commit] = true
			}

			newCommits := []*Commit{}
			for _, commit := range branch.Commits {
				if !existingCommits[commit] {
					newCommits = append(newCommits, commit)
					existingCommits[commit] = true
				}
			}

			r.ActiveBranch.Commits = append(newCommits, r.ActiveBranch.Commits...)
			r.Head = branch.Commits[0]

			log.Info(fmt.Sprintf("Branch %s merged", RenderEl(branchName)))
			return
		}
	}

	log.Warn(fmt.Sprintf("Branch %s not found", RenderEl(branchName)))
	return
}
