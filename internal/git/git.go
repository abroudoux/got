package git

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (repository *Repository) Init(repositoryName string) {
	repository.Name = repositoryName
	mainBranch := &Branch{Name: "main", Commits: []*Commit{}, LastCommit: nil}
	repository.Branches = append(repository.Branches, mainBranch)
	repository.ActiveBranch = mainBranch
	repository.Head = repository.ActiveBranch.LastCommit
}

func (repository *Repository) Commit(message string) {
	id := uuid.New().String()
	date := time.Now().Format("2006-01-02 15:04:05")
	newCommit := &Commit{Id: id, Message: message, Date: date}
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
		Name:       branchName,
		Commits:    append([]*Commit{}, repository.ActiveBranch.Commits...),
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
			existingCommits := make(map[*Commit]bool)
			for _, commit := range repository.ActiveBranch.Commits {
				existingCommits[commit] = true
			}

			for _, commit := range branch.Commits {
				if !existingCommits[commit] {
					repository.ActiveBranch.Commits = append(repository.ActiveBranch.Commits, commit)
					existingCommits[commit] = true
				}
			}

			repository.ActiveBranch.LastCommit = branch.LastCommit
			repository.Head = branch.LastCommit
			fmt.Printf("Branch %s merged\n", branchName)
			return
		}
	}

	fmt.Printf("Branch %s not found\n", branchName)
}

func reverseCommitsOrder(commits []*Commit) ([]*Commit, error) {
    var reversedCommits []*Commit
    for i := len(commits) - 1; i >= 0; i-- {
        reversedCommits = append(reversedCommits, commits[i])
    }
    return reversedCommits, nil
}