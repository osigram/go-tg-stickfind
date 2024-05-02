package handlers

import (
	"github.com/NicoNex/echotron/v3"
)

func (b *Bot) Update(update *echotron.Update) {
	if update.InlineQuery != nil {
		b.ProcessInlineQuery(update.InlineQuery)
		return
	}

	if update.Message != nil {
		b.Message(update.Message)
		return
	}
}
