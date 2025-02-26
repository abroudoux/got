package main

func Clone(remoteRepository *RemoteRepository) *LocalRepository {
	return &LocalRepository{
		Name:         remoteRepository.Name,
		Head:         remoteRepository.DefaultBranch.Commits[0],
		Origin:       remoteRepository,
		ActiveBranch: remoteRepository.DefaultBranch,
		Repository:   remoteRepository.Repository,
	}
}
