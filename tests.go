package main

func TestMerge() {
	repo := Init("test-merge")
	repo.Commit("commit-1")
	repo.Log()
	repo.Checkout("new-branch", true)
	repo.Commit("commit-2")
	repo.Log()
	repo.Checkout("main", false)
	repo.Log()
	repo.Merge("new-branch")
	repo.Log()
}

func TestPushRemote() {
	repo := Init("test-push-remote")
	repo.Commit("commit-1")
	remoteRepo := CreateRemoteRepository("test-push-remote")
	repo.RemoteAdd(remoteRepo)
	repo.Push(repo.ActiveBranch.Name)
	remoteRepo.LogRemoteRepository()
}
