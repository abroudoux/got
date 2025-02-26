package main

import "github.com/charmbracelet/log"

func (r *RemoteRepository) Clone() *LocalRepository {
	var localRepository = &LocalRepository{
		Name:       r.Name,
		Repository: r.Repository,
		Origin:     r,
	}

	if r.DefaultBranch == nil {
		log.Warn("You're trying to clone an empty repository.")
	} else {
		localRepository.Head = r.DefaultBranch.Commits[0]
		localRepository.ActiveBranch = r.DefaultBranch
	}

	return localRepository
}
