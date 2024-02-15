package language

import (
	"encoding/json"

	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func InitBundle(config *config.Config) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(config.FolderConfig.RootPath + "messages/ru.json")
	bundle.MustLoadMessageFile(config.FolderConfig.RootPath + "messages/tk.json")
	return bundle
}
