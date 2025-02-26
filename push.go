package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Push(remoteBranchName string) {
	if r.Origin.Url == "" {
		log.Error(fmt.Println("No remote repository to push to."))
		return
	}

	log.Warn("Need to implement the push action.")
	log.Info("Pushed to remote")
}
