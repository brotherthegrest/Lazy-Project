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
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type model struct {
	left  viewport.Model
	right viewport.Model
	input textinput.Model
}

var littlething = lipgloss.NewStyle().
	Background(lipgloss.Color("#ffffff")).
	MarginTop(2).
	MarginRight(4).
	MarginBottom(2).
	MarginLeft(20)

var bigthing = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff")).
	Margin(2).
	Width(2).
	Padding(2).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("228")).
	BorderBackground(lipgloss.Color("63")).
	BorderTop(true).
	BorderLeft(true)

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.left, _ = m.left.Update(msg)
	m.right, _ = m.right.Update(msg)
	m.input, _ = m.input.Update(msg)

	return m, nil
}
func (m model) View() string {
	return m.left.View() + m.input.View() + m.right.View()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Lazy-Project",
	Short: "terminal based project manager",
	Long:  `Lazy-Project is a terminal based project manager insipred by LazyGit`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// var tprogram *tea.Program
		fmt.Println(bigthing.Render("hello"))
		p := tea.NewProgram(model{})
		if err := p.Start(); err != nil {
			fmt.Printf("could not start program: %v", err)
			os.Exit(1)
		}

		fmt.Println(littlething.Render("hi"))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
