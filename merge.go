package main

import "fmt"

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
