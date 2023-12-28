/*
Copyright © 2023 Alan Hildebrandt <Alanhild715@gmail.com>
*/
package cmd

import (
	"fmt"

	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func startup() startupcalls {
	return startupcalls
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
		i := tea.NewProgram(initialModel())
		if _, err := i.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

		// Add a delay before exiting the program
		time.Sleep(1 * time.Second)
		// Now that the program has exited, we can use the height value from the model
		option1 := lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder(), true, true, true, true).
			Foreground(lipgloss.Color("#6f18f2")).
			Background(lipgloss.Color("#000000")).
			Height(10).
			Width(10).
			Faint(true).
			Italic(true).
			AlignVertical(lipgloss.Center)

		option2 := lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder(), true, true, true, true).
			Foreground(lipgloss.Color("#ea00ff")).
			Background(lipgloss.Color("#000000")).
			Height(m.height).
			Width(10).
			Padding(2).
			Margin(10)

		fmt.Print(option1.Render("joe mama"))
		fmt.Print(option2.Render("cream balls"))
		fmt.Println(m.height, m.width)

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
