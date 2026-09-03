package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/term"
)

type BoxConfig struct {
	Content []string
	Width   int
	Height  int
	Header  string
	Footer  string
	Border  lipgloss.Border
}

func borderText(line, text string) string {
	runes := []rune(line)
	pos := (len(runes) - len([]rune(text))) / 2
	return string(runes[:pos]) + text + string(runes[pos+len([]rune(text)):])
}

func Box(config BoxConfig) string {
	style := lipgloss.NewStyle().Width(config.Width).Height(config.Height).Align(lipgloss.Center).AlignVertical(lipgloss.Center).Border(config.Border)
	box := style.Render(config.Content...)
	lines := strings.Split(box, "\n")

	if config.Header != "" {
		lines[0] = borderText(lines[0], config.Header)
	}

	if config.Footer != "" {
		lines[len(lines)-1] = borderText(lines[len(lines)-1], config.Footer)
	}

	return strings.Join(lines, "\n")
}

func TerminalSize() (width, height int) {
	width, height, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		return 0, 0
	}

	return width, height
}

func ScreenBox(header, footer string) string {
	w, h := TerminalSize()
	box := BoxConfig{
		Content: []string{},
		Width:   w,
		Height:  h - 1,
		Header:  header,
		Footer:  footer,
		Border:  lipgloss.RoundedBorder(),
	}
	return Box(box) //HEY FUTURE ME REMOVE THIS -1 H WHEN U ARE DONE WITH DEBUGGING
}

func main() {
	box := BoxConfig{
		Content: []string{"LOL"},
		Width:   32,
		Height:  8,
		Header:  "PRANA",
		Footer:  "ESC TO EXIT",
		Border:  lipgloss.RoundedBorder(),
	}

	fmt.Println(Box(box))
}
