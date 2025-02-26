package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Checkout(branchName string) {
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
