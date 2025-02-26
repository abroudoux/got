package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Checkout(branchName string, create bool) {
	if create {
		for _, branch := range r.Repository.Branches {
			if branch.Name == branchName {
				log.Warn(fmt.Sprintf("Branch %s already exists", RenderEl(branchName)))
				return
			}
		}

		r.Branch(branchName)
	} else {
		for _, branch := range r.Repository.Branches {
			if branch.Name == branchName {
				r.ActiveBranch = branch
				r.Head = branch.Commits[0]

				log.Info(fmt.Sprintf("Switched to branch %s", RenderEl(branchName)))
				return
			}
		}

		log.Warn(fmt.Sprintf("Branch %s not found", RenderEl(branchName)))
		return
	}

	return
}
