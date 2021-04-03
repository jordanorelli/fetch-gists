package main

import (
	"fmt"
	"os/exec"
	"time"
)

type gist struct {
	ID      string    `json:"id"`
	Created time.Time `json:"created_at"`
}

func (g gist) clone() error {
	path := fmt.Sprintf("git@gist.github.com:%s.git", g.ID)
	cmd := exec.Command("git", "clone", path)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone gist %s: %w", g.ID, err)
	}
	return nil
}
