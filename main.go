package main

func main() {
	projectName := "got"
	repository := Init(projectName)
	repository.Commit("commit-1")
	remoteRepository := CreateRemoteRepository(projectName)
	repository.RemoteAdd(remoteRepository)
	repository.Push(repository.ActiveBranch.Name)
}
