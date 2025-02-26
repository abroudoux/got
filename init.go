package main

func (repository *Repository) Init(repositoryName string) {
	repository.Name = repositoryName
	mainBranch := &Branch{Name: "main", Commits: []*Commit{}, LastCommit: nil}
	repository.Branches = append(repository.Branches, mainBranch)
	repository.ActiveBranch = mainBranch
	repository.Head = repository.ActiveBranch.LastCommit
}
