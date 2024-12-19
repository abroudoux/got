package git

import (
	"fmt"
)

func (repository *Repository) printBranches() {
	for _, branch := range repository.Branches {
		fmt.Printf("%s\n", branch.Name)
	}
}

func (repository *Repository) printHead() {
	if repository.Head.Id == "" {
		fmt.Println("No commits yet")
		return
	}

	fmt.Printf("%s\n", repository.Head.Message)
}

func (repository *Repository) printActiveBranch() {
	if repository.ActiveBranch.Name == "" {
		fmt.Println("No active branch")
		return
	}

	fmt.Printf("%s\n", repository.ActiveBranch.Name)
}

func (repository *Repository) Debug() {
	println("COMMITS")
	repository.Log()
	println("BRANCHES")
	repository.printBranches()
	println("ACTIVE BRANCH")
	repository.printActiveBranch()
	println("HEAD")
	repository.printHead()
}