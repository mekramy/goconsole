package goconsole_test

import (
	"testing"

	"github.com/mekramy/goconsole"
)

func TestColors(t *testing.T) {
	tokens := []string{
		"@B{Bold} ",
		"@U{Underline} ",
		"@S{Strike} ",
		"@I{Italic} ",
		"\n",

		"@rb{ Red Background } ",
		"@gb{ Green Background } ",
		"@yb{ Yellow Background } ",
		"@bb{ Blue Background } ",
		"@pb{ Purple Background } ",
		"@cb{ Cyan Background } ",
		"@wb{ White Background } ",
		"\n",

		"@r{Red Foreground} ",
		"@g{Green Foreground} ",
		"@y{Yellow Foreground} ",
		"@b{Blue Foreground} ",
		"@p{Purple Foreground} ",
		"@c{Cyan Foreground} ",
		"@w{White Foreground} ",
		"\n",
	}

	for _, token := range tokens {
		goconsole.PrintF(token)
	}

	goconsole.Message().
		Green("Migrate contents").
		Italic().
		Strike().
		Print("")

	goconsole.Message().
		Blue("Message").
		Italic().
		Strike().
		Indent().
		Print("Welcome")
	goconsole.Message().
		Red("Error").
		Italic().
		Underline().
		Indent().
		Tags("One", "Two").
		Print("some data is invalid")

	t.Fatal("OOPS")
}
