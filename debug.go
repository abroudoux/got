package main

import (
	"fmt"
)

func (repository *Repository) printName() {
	fmt.Printf("%s\n", repository.Name)
}

func (repository *Repository) printBranches() {
	for _, branch := range repository.Branches {
		fmt.Printf("%s\n", branch.Name)
	}
}

func (repository *Repository) printHead() {
	if repository.Head == nil {
		fmt.Println("No commits yet")
	} else {
		fmt.Printf("%s (%s)\n", repository.Head.Message, repository.Head.Date)
	}
}

func (repository *Repository) printActiveBranch() {
	if repository.ActiveBranch.Name == "" {
		fmt.Println("No active branch")
		return
	}

	fmt.Printf("%s\n", repository.ActiveBranch.Name)
}

func (repository *Repository) printOrigin() {
	if repository.Origin == nil {
		fmt.Println("No origin")
		return
	}

	fmt.Printf("%s\n", repository.Origin.Repository.Name)
}

func (repository *Repository) Debug() {
	println("----------------")
	println("REPOSITORY DEBUG")
	println("----------------")
	println("COMMITS")
	repository.Log()
	println("BRANCHES")
	repository.printBranches()
	println("ACTIVE BRANCH")
	repository.printActiveBranch()
	println("HEAD")
	repository.printHead()
	println("ORIGIN")
	repository.printOrigin()
}

func (remoteRepository *RemoteRepository) printName() {
	fmt.Printf("%s\n", remoteRepository.Name)
}

func (remoteRepository *RemoteRepository) printCommits() {
	if remoteRepository.Repository.ActiveBranch == nil {
		fmt.Println("No commits yet")
		return
	}

	for _, commit := range remoteRepository.Repository.ActiveBranch.Commits {
		fmt.Printf("%s: %s -- %s \n", commit.Id, commit.Message, commit.Date)
	}
}

func (remoteRepository *RemoteRepository) printBranches() {
	if remoteRepository.Repository.Branches == nil {
		fmt.Println("No branches yet")
		return
	}

	for _, branch := range remoteRepository.Repository.Branches {
		fmt.Printf("%s\n", branch.Name)
	}
}

func (remoteRepository *RemoteRepository) printActiveBranch() {
	if remoteRepository.Repository.ActiveBranch.Name == "" {
		fmt.Println("No active branch")
		return
	}

	fmt.Printf("%s\n", remoteRepository.Repository.ActiveBranch.Name)
}

func (remoteRepository *RemoteRepository) Debug() {
	println("----------------")
	println("REMOTE REPOSITORY DEBUG")
	println("----------------")
	println("BRANCHES")
	remoteRepository.printBranches()
	println("COMMITS")
	remoteRepository.printCommits()
}
