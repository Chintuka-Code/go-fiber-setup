package localize

import (
	"github.com/fatih/color"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"

	"struct-validation/src/middleware"
)

var (
	DefaultLanguage = language.Danish
	AcceptLanguages = []language.Tag{language.Danish, language.English}
)

func NewLocalization() func(*fiber.Ctx) error {
	return fiberi18n.New(&fiberi18n.Config{
		RootPath:         "./locales",
		AcceptLanguages:  AcceptLanguages,
		DefaultLanguage:  DefaultLanguage,
		FormatBundleFile: "json",
	})
}

func GetTranslation(c *fiber.Ctx, key string) string {
	message, err := fiberi18n.Localize(c, key)
	lang := c.Get("Accept-Language")

	if lang == "" {
		lang = DefaultLanguage.String()
	}

	middleware.LOG.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	if err != nil {
		middleware.LOG.Warn(color.New(color.FgRed, color.Bold).Sprint("[LanguageError] ") +
			"Key '" + key + "' not found in the '" + lang + ".json' file")

		return key
	}

	return message
}
