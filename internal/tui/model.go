package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gnarl/pomodoro/internal/data"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title       string
	description string
	id          uint32
}

func (i item) Title() string       { return i.title }
func (i item) FilterValue() string { return i.title }
func (i item) Description() string { return i.description }

type model struct {
	list list.Model
	keys *listKeyMap
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// TODO add switch and break check for fitlering
		if msg.String() == "q" {
			return m, tea.Quit
		}
		if key.Matches(msg, m.keys.toggleTitleBar) {
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil
		}
		if key.Matches(msg, m.keys.toggleStatusBar) {
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func NewModel() model {
	favoriteTimers := data.ReadFavorites()
	items := make([]list.Item, len(favoriteTimers))
	for i, fav := range favoriteTimers {
		items[i] = item{
			title:       fav.Name,
			description: buildDescription(fav),
			id:          fav.Id,
		}
	}

	listKeys := newListKeyMap()
	favoritesList := list.NewModel(items, list.NewDefaultDelegate(), 0, 0)
	favoritesList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
		}
	}

	m := model{
		list: favoritesList,
		keys: newListKeyMap(),
	}
	m.list.Title = "Favorite Timers"

	return m
}

func buildDescription(favorite data.FavoriteTimer) string {
	t := favorite.Timer
	return fmt.Sprintf("Task: %s Duration: %d Message: %s", t.Task, t.Duration, t.Message)
}

type listKeyMap struct {
	toggleTitleBar  key.Binding
	toggleStatusBar key.Binding
}

func newListKeyMap() *listKeyMap {
	lkm := &listKeyMap{
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
	}
	return lkm
}
