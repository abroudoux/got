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
		r.ActiveBranch.Name = remoteBranchName
		remote.DefaultBranch = r.ActiveBranch
		remote.Repository.Branches = []*Branch{r.ActiveBranch}
	} else {
		found := false

		for _, branch := range remote.Repository.Branches {
			if branch.Name == remoteBranchName {
				found = true

				existingCommits := make(map[*Commit]bool)
				for _, commit := range branch.Commits {
					existingCommits[commit] = true
				}

				newCommits := []*Commit{}
				for _, commit := range r.ActiveBranch.Commits {
					if !existingCommits[commit] {
						newCommits = append([]*Commit{commit}, newCommits...)
					}
				}

				branch.Commits = append(newCommits, branch.Commits...)
				break
			}
		}

		if !found {
			newBranch := &Branch{
				Name:    remoteBranchName,
				Commits: r.ActiveBranch.Commits,
			}
			remote.Repository.Branches = append(remote.Repository.Branches, newBranch)
		}
	}

	log.Info(fmt.Sprintf("Pushed successfully on %s!", RenderEl(remote.Url)))
	return
}

func (r *RemoteRepository) IsRemoteRepositoryEmpty() bool {
	return r.DefaultBranch.Name == "" || len(r.Repository.Branches) == 0
}
