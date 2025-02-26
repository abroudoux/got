package main

func main() {
	var repository = &Repository{}
	repository.Init("my-repo")
	repository.Commit("commit-1")
	repository.Commit("commit-2")
	repository.Log()
}
