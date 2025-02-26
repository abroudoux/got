package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (r *Repository) Checkout(branchName string) {
	for _, branch := range r.Branches {
		if branch.Name == branchName {
			r.ActiveBranch = branch
			r.Head = branch.LastCommit

			log.Info(fmt.Sprintf("Switched to branch %s", branchName))
			return
		}
	}

	log.Warn(fmt.Sprintf("Branch %s not found", branchName))
}
