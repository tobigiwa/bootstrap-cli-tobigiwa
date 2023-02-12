package list

import (
	"fmt"

	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
)

func (b Bubble) View() string {
	var view string
	if b.state == navigationState {
		view = b.navigationList.View()
	} else if b.state == templateState {
		if b.frameworkType == _constants.FRONTEND_FRAMEWORKS {
			view = b.frontendTemplateList.View()
		} else if b.frameworkType == _constants.BACKEND_FRAMEWORKS {
			view = b.backendTemplateList.View()
		}
	} else if b.state == installState {
		if b.installError != nil {
			view = "Error: " + b.installError.Error() + "\n"
		} else if b.isInstalling == true {
			view = fmt.Sprintf("%s Installing ... ", b.spinner.View())
		} else {
			view = fmt.Sprintf("%s", b.installOutput)
		}
	}
	return bubbleStyle.Render(view)
}
