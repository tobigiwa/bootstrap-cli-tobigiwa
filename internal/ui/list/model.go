package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
)

var (
	vh int
	vw int
)

type sessionState int

const (
	navigationState sessionState = iota
	templateState
	installState
)

type Bubble struct {
	navigationList       list.Model
	frontendTemplateList list.Model
	backendTemplateList  list.Model
	frameworkType        string
	framework            string
	installOutput        []byte
	installError         error
	isInstalling         bool
	spinner              spinner.Model
	state                sessionState
}

func (b Bubble) Init() tea.Cmd {
	return b.spinner.Tick
}

func (b Bubble) GetFrameworkType() string {
	return b.frameworkType
}

func (b Bubble) GetFramework() string {
	return b.framework
}

func (b Bubble) GetSelecteNavigationItem() Item {
	item, ok := b.navigationList.SelectedItem().(Item)
	if ok {
		return item
	}
	return Item{}
}

func (b *Bubble) SetSize(width, height int) {
	b.navigationList.SetSize(
		width,
		height,
	)
	b.frontendTemplateList.SetSize(
		width,
		height,
	)
	b.backendTemplateList.SetSize(
		width,
		height,
	)
}

func New() Bubble {
	const defaultWidth = 20
	var navigationList list.Model
	var frontendTemplateList list.Model
	var backendTemplateList list.Model
	var items []list.Item

	// navigation
	items = []list.Item{
		Item{title: _constants.FRONTEND_FRAMEWORKS, desc: _constants.FRONTEND_FRAMEWORKS_DESC},
		Item{title: _constants.BACKEND_FRAMEWORKS, desc: _constants.BACKEND_FRAMEWORKS_DESC},
		Item{title: _constants.KUBERNETES_FRAMEWORKS, desc: _constants.KUBERNETES_FRAMEWORKS_DESC},
		Item{title: _constants.DOCKER_FRAMEWORKS, desc: _constants.DOCKER_FRAMEWORKS_DESC},
	}
	listDelegate := list.NewDefaultDelegate()
	listDelegate.Styles.SelectedTitle = delegateStyle
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy()

	navigationList = list.New(items, listDelegate, defaultWidth, listHeight)
	navigationList.Title = _constants.NAVIGATION_TEMPLATE_LIST_TITLE
	navigationList.Styles.Title = titleStyle
	navigationList.SetShowStatusBar(false)
	navigationList.SetFilteringEnabled(true)
	navigationList.Styles.PaginationStyle = paginationStyle
	navigationList.Styles.HelpStyle = helpStyle

	// frontend
	items = []list.Item{
		Item{title: _constants.FRONTEND_VUE, desc: _constants.FRONTEND_VUE_DESC, command: _constants.FRONTEND_VUE_CMD, commandArgs: _constants.FRONTEND_VUE_CMD_ARG},
		Item{title: _constants.FRONTEND_VUE_TS, desc: _constants.FRONTEND_VUE_TS_DESC, command: _constants.FRONTEND_VUE_TS_CMD, commandArgs: _constants.FRONTEND_VUE_TS_CMD_ARG},
		Item{title: _constants.FRONTEND_REACT, desc: _constants.FRONTEND_REACT_DESC, command: _constants.FRONTEND_REACT_CMD, commandArgs: _constants.FRONTEND_REACT_CMD_ARG},
		Item{title: _constants.FRONTEND_REACT_TS, desc: _constants.FRONTEND_REACT_TS_DESC, command: _constants.FRONTEND_REACT_TS_CMD, commandArgs: _constants.FRONTEND_REACT_TS_CMD_ARG},
		Item{title: _constants.FRONTEND_NEXT, desc: _constants.FRONTEND_NEXT_DESC, command: _constants.FRONTEND_NEXT_CMD, commandArgs: _constants.FRONTEND_NEXT_CMD_ARG},
		Item{title: _constants.FRONTEND_NEXT_TS, desc: _constants.FRONTEND_NEXT_TS_DESC, command: _constants.FRONTEND_NEXT_TS_CMD, commandArgs: _constants.FRONTEND_NEXT_TS_CMD_ARG},
		Item{title: _constants.FRONTEND_VANILLA, desc: _constants.FRONTEND_VANILLA_DESC, command: _constants.FRONTEND_VANILLA_CMD, commandArgs: _constants.FRONTEND_VANILLA_CMD_ARG},
		Item{title: _constants.FRONTEND_VANILLA_TS, desc: _constants.FRONTEND_VANILLA_TS_DESC, command: _constants.FRONTEND_VANILLA_TS_CMD, commandArgs: _constants.FRONTEND_VANILLA_TS_CMD_ARG},
		Item{title: _constants.FRONTEND_GATSBY, desc: _constants.FRONTEND_GATSBY_DESC, command: _constants.FRONTEND_GATSBY_CMD, commandArgs: _constants.FRONTEND_GATSBY_CMD_ARG},
		Item{title: _constants.FRONTEND_GATSBY_TS, desc: _constants.FRONTEND_GATSBY_TS_DESC, command: _constants.FRONTEND_GATSBY_TS_CMD, commandArgs: _constants.FRONTEND_GATSBY_TS_CMD_ARG},
	}
	frontendTemplateList = list.New(items, listDelegate, defaultWidth, listHeight)
	frontendTemplateList.Title = _constants.FRONTEND_TEMPLATE_LIST_TITLE

	frontendTemplateList.Styles.Title = titleStyle
	frontendTemplateList.SetShowStatusBar(false)
	frontendTemplateList.SetFilteringEnabled(true)
	frontendTemplateList.Styles.PaginationStyle = paginationStyle
	frontendTemplateList.Styles.HelpStyle = helpStyle

	// backend
	items = []list.Item{
		Item{
			title:       _constants.BACKEND_EXPRESS,
			desc:        _constants.BACKEND_EXPRESS_DESC,
			command:     _constants.BACKEND_EXPRESS_CMD,
			commandArgs: _constants.BACKEND_EXPRESS_CMD_ARG,
		},
	}
	backendTemplateList = list.New(items, listDelegate, defaultWidth, listHeight)
	backendTemplateList.Title = _constants.BACKEND_TEMPLATE_LIST_TITLE

	backendTemplateList.Styles.Title = titleStyle
	backendTemplateList.SetShowStatusBar(false)
	backendTemplateList.SetFilteringEnabled(true)
	backendTemplateList.Styles.PaginationStyle = paginationStyle
	backendTemplateList.Styles.HelpStyle = helpStyle

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return Bubble{
		navigationList:       navigationList,
		frontendTemplateList: frontendTemplateList,
		backendTemplateList:  backendTemplateList,
		spinner:              s,
		isInstalling:         false,
	}
}
