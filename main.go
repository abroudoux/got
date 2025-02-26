package main

func main() {
	projectName := "my-project"
	newBranchName := "new-branch"

	repository := Init(projectName)
	repository.LogBranches()
	repository.Commit("commit-1")
	repository.Branch(newBranchName)
	repository.LogBranches()
	repository.Checkout(newBranchName)
	repository.Commit("commit-2")
	repository.LogCommits()
	repository.Checkout("main")
	repository.LogCommits()
	repository.Merge(newBranchName)
	repository.LogCommits()

	// repository.RemoteAdd(remoteRepository)
	// repository.Push(repository.ActiveBranch.Name)
}
