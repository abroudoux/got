package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (repository *Repository) Merge(branchName string) {
	if repository.ActiveBranch.Name == branchName {
		log.Error(fmt.Sprintf("Cannot merge branch %s into itself", branchName))
		return
	}

	for _, branch := range repository.Branches {
		if branch.Name == branchName {
			existingCommits := make(map[*Commit]bool)
			for _, commit := range repository.ActiveBranch.Commits {
				existingCommits[commit] = true
			}

			newCommits := []*Commit{}
			for _, commit := range branch.Commits {
				if !existingCommits[commit] {
					newCommits = append(newCommits, commit)
					existingCommits[commit] = true
				}
			}

			repository.ActiveBranch.Commits = append(newCommits, repository.ActiveBranch.Commits...)
			repository.ActiveBranch.LastCommit = branch.LastCommit
			repository.Head = branch.LastCommit

			log.Info(fmt.Sprintf("Branch %s merged", branchName))
			return
		}
	}

	log.Warn(fmt.Sprintf("Branch %s not found", branchName))
}
