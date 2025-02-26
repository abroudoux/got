package main

import (
	"github.com/charmbracelet/log"
)

func (r *Repository) Push(remoteBranchName string) {
	// if r.Origin == &RemoteRepository{} {} {
	// 	log.Error(fmt.Println("No remote repository to push to."))
	// 	return
	// }

	log.Info("Pushed to remote")
}
