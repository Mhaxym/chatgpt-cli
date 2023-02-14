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
	"google-chrome": true,
	"brave-browser": true,
	"safari":        true,
}
var DEFAULT_BROWSER = os.Getenv("CHATGPT_DEFAULT_BROWSER")

func main() {

	browserPtr := flag.String("browser", "", "Opens ChatGPT in the specified browser")
	browserListPtr := flag.Bool("list-browsers", false, "Lists the supported browsers")
	setDefaultBrowserPtr := flag.String("set-default-browser", "", "Sets the default browser as a variable in the .bashrc file named CHATGPT_DEFAULT_BROWSER")
	helpPtr := flag.Bool("help", false, "Shows the help message")
	flag.Parse()

	if *helpPtr {
		help()
		return
	}

	if *browserListPtr {
		supportedBrowsers()
		return
	}

	if *setDefaultBrowserPtr != "" {
		setDefaultBrowser(setDefaultBrowserPtr)
		return
	}

	if *browserPtr != "" {
		if !ALLOWED_BROWSERS[*browserPtr] {
			fmt.Printf("Browser %s not supported", *browserPtr)
			return
		}
		DEFAULT_BROWSER = *browserPtr
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
	cmd := exec.Command(DEFAULT_BROWSER, url)
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

		bashrc := os.Getenv("HOME") + "/.bashrc"
		f, err := os.OpenFile(bashrc, os.O_RDWR, 0600)
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
