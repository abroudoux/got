package main

func main() {
	projectName := "my-project"
	// newBranchName := "new-branch"

	repository := Init(projectName)
	repository.LogBranches()
	repository.Commit("commit-1")
	repository.LogCommits()

	// repository.RemoteAdd(remoteRepository)
	// repository.Push(repository.ActiveBranch.Name)
}
