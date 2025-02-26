package main

func main() {
	projectName := "my-project"
	newBranchName := "new-branch"

	repository := Init(projectName)
	repository.Commit("commit-1")
	repository.Commit("commit-2")
	repository.Branch(newBranchName)
	repository.Checkout(newBranchName)
	repository.Commit("commit-3")
	repository.Log()
	repository.Checkout("main")
	repository.Log()
	repository.Merge(newBranchName)
	repository.Log()
}
