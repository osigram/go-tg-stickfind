package handlers

import "log/slog"

func (b *Bot) SetOCRKey(userID int64, params ...string) string {
	l := b.app.Logger.With(slog.String("op", "internal.handlers.SetOCRToken"))

	if len(params) != 1 {
		return "Invalid number of parameters."
	}
	key := params[0]

	if key == "" {
		return "Token is empty."
	}

	if err := b.app.Storage.SetUserKey(userID, key); err != nil {
		l.Error("Error to set OCR token", slog.String("error", err.Error()))
		return "Error to set OCR token."
	}

	return "OCR token has been set."
}
