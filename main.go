package main

import (
	"flag"
	"fmt"
	"os"
)

var ALLOWED_BROWSERS = map[string]bool{
	"firefox":       true,
	"brave":         true,
	"chromium":      true,
	"brave-browser": true,
	"safari":        true,
}
var DEFAULT_BROWSER = os.Getenv("CHATGPT_DEFAULT_BROWSER")
var FLAGS *Flags = &Flags{}

func main() {

	flag.BoolVar(&FLAGS.Help, "help", false, "Shows the help message")

	// Flags to set the default browser
	flag.Var(&FLAGS.SetDefaultBrowser, "set-default-browser", "Sets the default browser as a variable in the .bashrc file named CHATGPT_DEFAULT_BROWSER")
	flag.StringVar(&FLAGS.ConfigFile, "config-file", "bashrc", "If your shell environment is not bash, please specify the config file path (e.g. zshrc)\nThis file will be used to set the CHATGPT_DEFAULT_BROWSER variable.\nOnly use this flag with the --set-default-browser flag.")

	flag.BoolVar(&FLAGS.ListBrowsers, "list-browsers", false, "Lists the supported browsers")
	flag.Parse()

	if err := FLAGS.Validate(); err != nil {
		fmt.Println(ConsoleStyler(err.Error(), &StylerConfig{Color: RED, Bold: true}))
		return
	}

	if flag.NFlag() == 0 {
		openChatGPT()
		return
	}

	if FLAGS.Help {
		help()
		return
	}

	if FLAGS.ListBrowsers {
		supportedBrowsers()
		return
	}

	if FLAGS.SetDefaultBrowser.Active {
		setDefaultBrowser()
		return
	}
}
