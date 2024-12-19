package git

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Repository struct {
	Name string
	Head Commit
	Branches []*Branch
	ActiveBranch *Branch
}

type Commit struct {
	Id string
	Message string
	Date string
}

type Branch struct {
	Name string
	Commits []Commit
	LastCommit Commit
}

func (repository *Repository) Init(repositoryName string) {
	repository.Name = repositoryName
	mainBranch := &Branch{Name: "main", Commits: []Commit{}, LastCommit: Commit{}}
	repository.Branches = append(repository.Branches, mainBranch)
	repository.ActiveBranch = mainBranch
	repository.Head = repository.ActiveBranch.LastCommit
}

func (repository *Repository) Commit(message string) {
	id := uuid.New().String()
	date := time.Now().Format("2006-01-02 15:04:05")
	newCommit := Commit{Id: id, Message: message, Date: date}
	repository.ActiveBranch.Commits = append(repository.ActiveBranch.Commits, newCommit)
	repository.ActiveBranch.LastCommit = newCommit
	repository.Head = repository.ActiveBranch.LastCommit

	fmt.Printf("Commit %s created at %s\n", id, date)
}

func (repository *Repository) Log() {
	if len(repository.ActiveBranch.Commits) == 0 {
		fmt.Println("No commits yet")
		return
	}

	reversedCommits, err := reverseCommitsOrder(repository.ActiveBranch.Commits)
	if err != nil {
		fmt.Println("Error reversing commits order")
		return
	}

	for _, commit := range reversedCommits {
		fmt.Printf("%s: %s -- %s \n", commit.Id, commit.Message, commit.Date)
	}
}

func (repository *Repository) Branch(branchName string) {
	branch := &Branch{
		Name:     branchName,
		Commits:  append([]Commit{}, repository.ActiveBranch.Commits...),
		LastCommit: repository.Head,
	}
	repository.Branches = append(repository.Branches, branch)

	fmt.Printf("Branch %s created at commit %s\n", branch.Name, branch.LastCommit.Id)
}

func (repository *Repository) Checkout(branchName string) {
	for _, branch := range repository.Branches {
		if branch.Name == branchName {
			repository.ActiveBranch = branch
			repository.Head = branch.LastCommit
			fmt.Printf("Switched to branch %s\n", branchName)
			return
		}
	}

	fmt.Printf("Branch %s not found\n", branchName)
}

func (repository *Repository) Merge(branchName string) {
	if repository.ActiveBranch.Name == branchName {
		fmt.Printf("Cannot merge branch %s into itself\n", branchName)
		return
	}

	for _, branch := range repository.Branches {
		if branch.Name == branchName {
			repository.ActiveBranch.Commits = append(repository.ActiveBranch.Commits, branch.Commits...)
			repository.ActiveBranch.LastCommit = branch.LastCommit
			repository.Head = branch.LastCommit
			fmt.Printf("Branch %s merged\n", branchName)
			return
		}
	}

	fmt.Printf("Branch %s not found\n", branchName)
}

func reverseCommitsOrder(commits []Commit) ([]Commit, error) {
	var reversedCommits []Commit
	for i := len(commits) - 1; i >= 0; i-- {
		reversedCommits = append(reversedCommits, commits[i])
	}

	return reversedCommits, nil
}

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