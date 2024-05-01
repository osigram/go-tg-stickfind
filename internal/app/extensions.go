package app

import "github.com/NicoNex/echotron/v3"

func (app *App) SendTextReply(text string, chatID int64, replyID int) error {
	_, err := app.SendMessage(text, chatID, &echotron.MessageOptions{
		ReplyParameters: echotron.ReplyParameters{MessageID: replyID},
	})

	return err
}
