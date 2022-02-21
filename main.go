package main

import (
	"fmt"
	"os"

	cmd "refsiverdur.org/node-pensioner-manager/v2/cmd"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(cmd.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
