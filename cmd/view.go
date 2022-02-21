package cmd

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	registry "refsiverdur.org/node-pensioner-manager/v2/registry"
)

type model struct {
	textInput textinput.Model
	guesses   map[string]int
	hasResult bool
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "@vleesbrood/unbg"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		hasResult: false,
		guesses:   map[string]int{},
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.hasResult = false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			age := registry.GetPackageAge(m.textInput.Value())
			m.guesses[m.textInput.Value()] = age
			return m, cmd
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func sortedMapKeys(items map[string]int) []string {
	var keys []string
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (m model) View() string {
	var result string

	if len(m.guesses) > 0 {
		result += "Previous guesses:\n\n"
	}

	for _, key := range sortedMapKeys(m.guesses) {
		result += fmt.Sprintf("%s: %d\n", key, m.guesses[key])
	}

	result += "\n"

	result += fmt.Sprintf(
		"Give me a package name!\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
	return result
}
