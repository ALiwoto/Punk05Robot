package channelsPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	registerCommand := handlers.NewCommand(registerCmd, registerCommandResponse)
	tmpIgnoreCommand := handlers.NewCommand(tmpIgnoreCmd, tmpIgnoreResponse)
	addUserCommand := handlers.NewCommand(addUserCmd, addUserResponse)
	setFooterCommand := handlers.NewCommand(setFooterCmd, setFooterResponse)
	setButtonsCommand := handlers.NewCommand(setButtonsCmd, setButtonsResponse)
	addButtonsCommand := handlers.NewCommand(addButtonsCmd, addButtonsResponse)
	removeButtonsCommand := handlers.NewCommand(removeButtonsCmd, removeButtonsResponse)

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
