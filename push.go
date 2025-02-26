package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Push(remoteBranchName string) {
	if r.Origin.Url == "" {
		log.Warn(fmt.Println("No remote repository to push to."))
		return
	}

	remote := r.Origin

	if remote.IsRemoteRepositoryEmpty() {
		remote.DefaultBranch = r.ActiveBranch
		remote.DefaultBranch.Name = remoteBranchName
		return
	}

	remote.Log()

	log.Info(fmt.Sprintf("Pushed successfully on %s!", RenderEl(remote.Url)))
	return
}

func (r *RemoteRepository) IsRemoteRepositoryEmpty() bool {
	return r.DefaultBranch.Name == "" || len(r.Repository.Branches) == 0
}
