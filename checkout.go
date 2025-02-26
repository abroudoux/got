package main

import "fmt"

func (r *Repository) Checkout(branchName string) {
	for _, branch := range r.Branches {
		if branch.Name == branchName {
			r.ActiveBranch = branch
			r.Head = branch.LastCommit
			fmt.Printf("Switched to branch %s\n", branchName)
			return
		}
	}

	fmt.Printf("Branch %s not found\n", branchName)
}
