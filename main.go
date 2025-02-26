package main

func main() {
	projectName := "my-project"
	// newBranchName := "new-branch"

	repository := Init(projectName)
	repository.Commit("commit-1")
	repository.Log()
	remoteRepository := CreateRemoteRepository(projectName)
	repository.RemoteAdd(remoteRepository)
	remoteRepository.LogCommits()
	remoteRepository.LogBranches()
}
