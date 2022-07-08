package common

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
)

func TestLocaleString_Get(t *testing.T) {
	norwegianEntry := "norwegian entry"
	englishEntry := "english entry"
	swedishEntry := "swedish entry"
	localeString := LocaleString{
		"no": null.StringFrom(norwegianEntry),
		"en": null.StringFrom(englishEntry),
		"sv": null.StringFrom(swedishEntry),
	}

	assert.Equal(t, localeString.Get([]string{"no"}), norwegianEntry)
	// Should fall back to english
	assert.Equal(t, localeString.Get([]string{}), englishEntry)
	assert.Equal(t, localeString.Get([]string{"da", "sv"}), swedishEntry)
}
