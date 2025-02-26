package main

import "fmt"

func (repository *Repository) Push() {
	if repository.Origin == nil {
		fmt.Println("No remote repository to push to.")
		return
	}

	println("Pushed to remote")
}
