package main

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func (r *LocalRepository) Status() {
	log.Info(fmt.Sprintf("On branch %s", RenderEl(r.ActiveBranch.Name)))
}
