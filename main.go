package main

func main() {
	projectName := "got"

	repository := Init(projectName)
	repository.Commit("commit-1")
	repository.Commit("commit-2")
	repository.Checkout("new-branch", true)
	repository.Commit("commit-3")
	repository.Log()
	repository.Checkout("main", false)
	repository.Log()
	repository.Merge("new-branch")
	repository.Log()
	// remoteRepository := CreateRemoteRepository(projectName)
	// repository.RemoteAdd(remoteRepository)
	// repository.Push(repository.ActiveBranch.Name)
}
