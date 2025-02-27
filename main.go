package main

func main() {
	projectName := "got"

	repository := Init(projectName)
	remoteRepository := CreateRemoteRepository(projectName)
	repository.RemoteAdd(remoteRepository)
	repository.Commit("commit-1")
	repository.Commit("commit-2")
	repository.Push(repository.ActiveBranch.Name)
	remoteRepository.LogRemoteRepository()
	repository.Checkout("new-branch", true)
	repository.Commit("commit-3")
	repository.Push(repository.ActiveBranch.Name)
	remoteRepository.LogRemoteRepository()
	repository.Checkout("main", false)
	repository.Merge("new-branch")
	repository.Push(repository.ActiveBranch.Name)
	remoteRepository.LogRemoteRepository()
	repository.Checkout("new-branch", false)
	repository.Commit("commit-4")
	repository.Push(repository.ActiveBranch.Name)
	remoteRepository.LogRemoteRepository()
}
