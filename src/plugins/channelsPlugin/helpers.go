package channelsPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	registerCommand := handlers.NewCommand(registerCmd, registerCommandResponse)
	tmpIgnoreCommand := handlers.NewCommand(tmpIgnoreCmd, registerCommandResponse)

	registerCommand.Triggers = t
	tmpIgnoreCommand.Triggers = t

	d.AddHandler(registerCommand)
	d.AddHandler(tmpIgnoreCommand)
}
