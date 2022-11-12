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
	allowFooterCommand := sHandlers.SudoOnlyCommand(allowFooterCmd, allowFooterResponse)
	disallowFooterCommand := sHandlers.SudoOnlyCommand(disallowFooterCmd, disallowFooterResponse)
	allowCaptionCommand := sHandlers.SudoOnlyCommand(allowCaptionCmd, allowCaptionResponse)
	disallowCaptionCommand := sHandlers.SudoOnlyCommand(disallowCaptionCmd, disallowCaptionResponse)
	allowButtonsCommand := sHandlers.SudoOnlyCommand(allowButtonsCmd, allowButtonsResponse)
	disallowButtonsCommand := sHandlers.SudoOnlyCommand(disallowButtonsCmd, disallowButtonsResponse)

	registerCommand.Triggers = t
	tmpIgnoreCommand.Triggers = t
	addUserCommand.Triggers = t
	setFooterCommand.Triggers = t
	setButtonsCommand.Triggers = t
	addButtonsCommand.Triggers = t
	removeButtonsCommand.Triggers = t
	allowFooterCommand.Triggers = t
	disallowFooterCommand.Triggers = t
	allowCaptionCommand.Triggers = t
	disallowCaptionCommand.Triggers = t
	allowButtonsCommand.Triggers = t
	disallowButtonsCommand.Triggers = t

	d.AddHandler(registerCommand)
	d.AddHandler(tmpIgnoreCommand)
	d.AddHandler(addUserCommand)
	d.AddHandler(setFooterCommand)
	d.AddHandler(setButtonsCommand)
	d.AddHandler(addButtonsCommand)
	d.AddHandler(removeButtonsCommand)
	d.AddHandler(allowFooterCommand)
	d.AddHandler(disallowFooterCommand)
	d.AddHandler(allowCaptionCommand)
	d.AddHandler(disallowCaptionCommand)
	d.AddHandler(allowButtonsCommand)
	d.AddHandler(disallowButtonsCommand)
}
