package main

func main() {
	projectName := "my-project"
	repository := Init(projectName)
	repository.Commit("commit-1")
	remoteRepository := CreateRemoteRepository(projectName)
	repository.RemoteAdd(remoteRepository.Url)
	repository.Push(repository.ActiveBranch.Name)
}
