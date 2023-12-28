/*
Copyright © 2023 Alan Hildebrandt <Alanhild715@gmail.com>
*/
package cmd

import (
	"fmt"

	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// calling the full screen msg
type startupcalls struct{}

type model struct {
	left       viewport.Model
	right      viewport.Model
	input      textinput.Model
	width      int
	height     int
	fullscreen bool
}

func (m model) Init() tea.Cmd {
	return nil

}

// dosent matter
func initialModel() model {
	return model{}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		m.left, _ = m.left.Update(msg)
		m.right, _ = m.right.Update(msg)
		m.input, _ = m.input.Update(msg)

	}
	return m, nil

}

func (m model) View() string {
	return m.left.View() + m.input.View() + m.right.View()

}

var rootCmd = &cobra.Command{
	Use:   "Lazy-Project",
	Short: "terminal based project manager",
	Long:  `Lazy-Project is a terminal based project manager insipred by LazyGit`,
	Run: func(cmd *cobra.Command, args []string) {
		m := model{}
		p := tea.NewProgram(&m)
		if _, err := p.Run(); err != nil {
			fmt.Printf("could not start program: %v", err)
			os.Exit(1)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Lazy-Project.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
