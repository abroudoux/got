package main

func main() {
	projectName := "got"
	newBranchName := "new-branch"

	repository := Init(projectName)
	repository.Commit("commit-1")
	repository.Checkout(newBranchName, true)
	repository.Commit("commit-2")
	repository.Checkout("main", false)
	repository.Log()
	repository.Merge(newBranchName)
	repository.Log()

	// repository.RemoteAdd(remoteRepository)
	// repository.Push(repository.ActiveBranch.Name)
}
