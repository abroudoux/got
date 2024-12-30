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
	reversedCommits, err := ReverseCommitsOrder(repository.ActiveBranch.Commits)
	if err != nil {
		fmt.Println("Error reversing commits order")
		return
	}

	repository.ActiveBranch.Commits = reversedCommits
	repository.ActiveBranch.LastCommit = newCommit
	repository.Head = repository.ActiveBranch.LastCommit

	fmt.Printf("Commit %s created at %s\n", id, date)
}

func (repository *Repository) Log() {
	if len(repository.ActiveBranch.Commits) == 0 {
		fmt.Println("No commits yet")
		return
	}

	for _, commit := range repository.ActiveBranch.Commits {
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

func (remoteRepository *RemoteRepository) CreateRemoteRepository(repositoryName string) {
    remoteRepository.Name = repositoryName
    remoteRepository.Url = fmt.Sprintf("https://gothub.com/%s", remoteRepository.Name)
	remoteRepository.Repository = &Repository{}

    fmt.Printf("Remote repository %s created at %s\n", remoteRepository.Name, remoteRepository.Url)
}
func (repository *Repository) RemoteAdd(remoteRepository *RemoteRepository) {
    repository.Origin = remoteRepository
    fmt.Printf("Remote %s added\n", remoteRepository.Url)
}

func (repository *Repository) Push() {
    if repository.Origin == nil {
        fmt.Println("No remote repository to push to.")
        return
    }

    // if repository.ActiveBranch != nil {
    //     repository.Origin.Repository.Branches = append(repository.Origin.Repository.Branches, repository.ActiveBranch)
	// 	repository.Origin.Repository.ActiveBranch.Commits = repository.ActiveBranch.Commits
    //     fmt.Printf("Pushed branch %s to %s\n", repository.ActiveBranch.Name, repository.Origin.Url)
    // } else {
    //     fmt.Println("No active branch to push.")
    // }

	println("Pushed to remote")
}