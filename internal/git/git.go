package git

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func Init(repositoryName string) *Repository {
	mainBranch := &Branch{Name: "main", Commits: []*Commit{}}
	newRepository := &Repository{
		Name:        repositoryName,
		Branches:    []*Branch{mainBranch},
		ActiveBranch: mainBranch,
		Head:        nil,
	}

	return newRepository
}

func (repository *Repository) LogBranches() {
	if len(repository.Branches) == 0 {
		fmt.Println("No branches yet")
		return
	}

	for _, branch := range repository.Branches {
		fmt.Println(branch.Name)
	}
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
	repository.Head = repository.ActiveBranch.Commits[len(repository.ActiveBranch.Commits) - 1]

	fmt.Printf("Commit: %s: %s -- %s \n", id, message, date)
}

func (repository *Repository) LogCommits() {
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
	}
	repository.Branches = append(repository.Branches, branch)

	fmt.Printf("Branch %s created at commit %s\n", branch.Name, repository.Head.Id)
}

func (repository *Repository) Checkout(branchName string) {
	for _, branch := range repository.Branches {
		if branch.Name == branchName {
			repository.ActiveBranch = branch
			repository.Head = branch.Commits[len(branch.Commits) - 1]
			fmt.Printf("Switched to branch %s\n", branchName)
			return
		}
	}

	fmt.Printf("Branch %s not found\n", branchName)
}

// func (repository *Repository) Merge(branchName string) {
// 	if repository.ActiveBranch.Name == branchName {
// 		fmt.Printf("Cannot merge branch %s into itself\n", branchName)
// 		return
// 	}

// 	for _, branch := range repository.Branches {
// 		if branch.Name == branchName {
// 			existingCommits := make(map[*Commit]bool)
// 			for _, commit := range repository.ActiveBranch.Commits {
// 				existingCommits[commit] = true
// 			}

// 			for _, commit := range branch.Commits {
// 				if !existingCommits[commit] {
// 					repository.ActiveBranch.Commits = append(repository.ActiveBranch.Commits, commit)
// 					existingCommits[commit] = true
// 				}
// 			}

// 			repository.ActiveBranch.LastCommit = branch.LastCommit
// 			repository.Head = branch.LastCommit
// 			fmt.Printf("Branch %s merged\n", branchName)
// 			return
// 		}
// 	}

// 	fmt.Printf("Branch %s not found\n", branchName)
// }

// func (remoteRepository *RemoteRepository) CreateRemoteRepository(repositoryName string) {
//     remoteRepository.Name = repositoryName
//     remoteRepository.Url = fmt.Sprintf("https://gothub.com/%s", remoteRepository.Name)
// 	remoteRepository.Repository = &Repository{}

//     fmt.Printf("Remote repository %s created at %s\n", remoteRepository.Name, remoteRepository.Url)
// }
// func (repository *Repository) RemoteAdd(remoteRepository *RemoteRepository) {
//     repository.Origin = remoteRepository
//     fmt.Printf("Remote %s added\n", remoteRepository.Url)
// }

// func (repository *Repository) Push() {
//     if repository.Origin == nil {
//         fmt.Println("No remote repository to push to.")
//         return
//     }

//     // if repository.ActiveBranch != nil {
//     //     repository.Origin.Repository.Branches = append(repository.Origin.Repository.Branches, repository.ActiveBranch)
// 	// 	repository.Origin.Repository.ActiveBranch.Commits = repository.ActiveBranch.Commits
//     //     fmt.Printf("Pushed branch %s to %s\n", repository.ActiveBranch.Name, repository.Origin.Url)
//     // } else {
//     //     fmt.Println("No active branch to push.")
//     // }

// 	println("Pushed to remote")
// }