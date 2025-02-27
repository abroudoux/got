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
	} else {
		for _, branch := range remote.Repository.Branches {
			if branch.Name == remoteBranchName {
				// todo: I'll need to compare both
				branch.Commits = append(branch.Commits, r.ActiveBranch.Commits...)
				return
			}

			remote.Repository.Branches = append(remote.Repository.Branches, &Branch{
				Name:    remoteBranchName,
				Commits: r.ActiveBranch.Commits,
			})
		}
	}

	log.Info(fmt.Sprintf("Pushed successfully on %s!", RenderEl(remote.Url)))
	return
}

func (r *RemoteRepository) IsRemoteRepositoryEmpty() bool {
	return r.DefaultBranch.Name == "" || len(r.Repository.Branches) == 0
}
