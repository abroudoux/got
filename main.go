package main

func main() {
	projectName := "my-project"
	// newBranchName := "new-branch"

	repository := Init(projectName)
	repository.Commit("commit-1")
	remoteRepository := CreateRemoteRepository(projectName)
	repository.RemoteAdd(remoteRepository)
	repository.Push(repository.ActiveBranch.Name)
}
