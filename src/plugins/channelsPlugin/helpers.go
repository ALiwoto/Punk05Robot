package channelsPlugin

import (
	sHandlers "github.com/AnimeKaizoku/Punk05Robot/src/core/specialHandlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	registerCommand := sHandlers.SudoOnlyCommand(registerCmd, registerCommandResponse)
	tmpIgnoreCommand := sHandlers.SudoOnlyCommand(tmpIgnoreCmd, tmpIgnoreResponse)
	addUserCommand := sHandlers.SudoOnlyCommand(addUserCmd, addUserResponse)
	setFooterCommand := sHandlers.SudoOnlyCommand(setFooterCmd, setFooterResponse)
	setButtonsCommand := sHandlers.SudoOnlyCommand(setButtonsCmd, setButtonsResponse)
	addButtonsCommand := sHandlers.SudoOnlyCommand(addButtonsCmd, addButtonsResponse)
	removeButtonsCommand := sHandlers.SudoOnlyCommand(removeButtonsCmd, removeButtonsResponse)

	registerCommand.Triggers = t
	tmpIgnoreCommand.Triggers = t
	addUserCommand.Triggers = t
	setFooterCommand.Triggers = t
	setButtonsCommand.Triggers = t
	addButtonsCommand.Triggers = t
	removeButtonsCommand.Triggers = t

	d.AddHandler(registerCommand)
	d.AddHandler(tmpIgnoreCommand)
	d.AddHandler(addUserCommand)
	d.AddHandler(setFooterCommand)
	d.AddHandler(setButtonsCommand)
	d.AddHandler(addButtonsCommand)
	d.AddHandler(removeButtonsCommand)
}
