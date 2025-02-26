package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (repository *Repository) Push() {
	if repository.Origin == nil {
		log.Error(fmt.Println("No remote repository to push to."))
		return
	}

	log.Info("Pushed to remote")
}
