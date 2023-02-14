package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var ALLOWED_BROWSERS = map[string]bool{
	"firefox":       true,
	"brave":         true,
	"chromium":      true,
	"brave-browser": true,
	"safari":        true,
}
var DEFAULT_BROWSER = os.Getenv("CHATGPT_DEFAULT_BROWSER")

var FLAGS struct {
	Browser           string
	ListBrowsers      bool
	SetDefaultBrowser string
	ConfigFile        string
	Help              bool
}

func main() {

	flag.StringVar(&FLAGS.Browser, "browser", "", "Opens ChatGPT in the specified browser")
	flag.BoolVar(&FLAGS.ListBrowsers, "list-browsers", false, "Lists the supported browsers")
	flag.StringVar(&FLAGS.SetDefaultBrowser, "set-default-browser", "", "Sets the default browser as a variable in the .bashrc file named CHATGPT_DEFAULT_BROWSER")
	flag.StringVar(&FLAGS.ConfigFile, "config-file", "bashrc", "If your shell environment is not bash, please specify the config file path (e.g. zshrc)\nThis file will be used to set the CHATGPT_DEFAULT_BROWSER variable.\nOnly use this flag with the --set-default-browser flag.")
	flag.BoolVar(&FLAGS.Help, "help", false, "Shows the help message")

	flag.Parse()

	if FLAGS.Help {
		help()
		return
	}

	if FLAGS.ListBrowsers {
		supportedBrowsers()
		return
	}

	if FLAGS.SetDefaultBrowser != "" {
		setDefaultBrowser(&FLAGS.SetDefaultBrowser)
		return
	}

	if FLAGS.Browser != "" {
		overrideBrowser()
	}

	openChatGPT()
}

func openChatGPT() {
	if DEFAULT_BROWSER == "" {
		fmt.Println("Default browser not set")
		fmt.Println("Please set the default browser using the --set-default-browser flag")
		return
	}
	fmt.Println("Starting ChatGPT in ", DEFAULT_BROWSER)
	url := "https://chat.openai.com/chat"

	var cmd *exec.Cmd = nil
	_, err := exec.LookPath(DEFAULT_BROWSER)
	if err != nil {
		cmd = exec.Command("open", "-a", DEFAULT_BROWSER, url)
	} else {
		cmd = exec.Command(DEFAULT_BROWSER, url)
	}
	cmd.Start()
}

func help() {
	fmt.Println("Information about the chatgpt-cli command")
	fmt.Println("Usage: chatgpt-cli [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func supportedBrowsers() {
	fmt.Println("Supported browsers:")
	for browser := range ALLOWED_BROWSERS {
		fmt.Println("\t", browser)
	}
}

func overrideBrowser() {
	if !ALLOWED_BROWSERS[FLAGS.Browser] {
		fmt.Printf("Browser %s not supported", FLAGS.Browser)
		return
	}
	DEFAULT_BROWSER = FLAGS.Browser
}

func setDefaultBrowser(defaultBrowser *string) {
	if *defaultBrowser != "" {
		if !ALLOWED_BROWSERS[*defaultBrowser] {
			fmt.Printf("Browser %s not supported \n", *defaultBrowser)
			return
		}
		err := os.Setenv("CHATGPT_DEFAULT_BROWSER", *defaultBrowser)
		if err != nil {
			panic(err)
		}

		configFile := os.Getenv("HOME") + "/." + FLAGS.ConfigFile
		f, err := os.OpenFile(configFile, os.O_RDWR, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		var lines []string
		var found bool

		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "export CHATGPT_DEFAULT_BROWSER=") {
				line = fmt.Sprintf("export CHATGPT_DEFAULT_BROWSER=%s", os.Getenv("CHATGPT_DEFAULT_BROWSER"))
				found = true
			}
			lines = append(lines, line)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

		if !found {
			lines = append(lines, fmt.Sprintf("export CHATGPT_DEFAULT_BROWSER=%s", os.Getenv("CHATGPT_DEFAULT_BROWSER")))
		}

		if err := f.Truncate(0); err != nil {
			panic(err)
		}

		if _, err := f.Seek(0, 0); err != nil {
			panic(err)
		}

		w := bufio.NewWriter(f)
		for _, line := range lines {
			fmt.Fprintln(w, line)
		}
		if err := w.Flush(); err != nil {
			panic(err)
		}

		fmt.Printf("Default browser set to %s \n", *defaultBrowser)
		fmt.Println("Please restart your terminal to apply the changes")
	} else {
		fmt.Println("Default browser not set")
	}
}
