package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

func openChatGPT() {
	if DEFAULT_BROWSER == "" {
		stylerConfig := &StylerConfig{Color: RED, Bold: true}
		fmt.Println(ConsoleStyler("Default browser not set", stylerConfig))
		fmt.Println(ConsoleStyler("Please set the default browser using the -set-default-browser flag", stylerConfig))
		return
	}
	fmt.Println(ConsoleStyler("Starting ChatGPT in "+strings.Title(DEFAULT_BROWSER), &StylerConfig{Color: GREEN, Bold: true}))
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

func setDefaultBrowser() {
	if FLAGS.SetDefaultBrowser.Value != "" {
		err := os.Setenv("CHATGPT_DEFAULT_BROWSER", FLAGS.SetDefaultBrowser.Value)
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

		stylerConfig := &StylerConfig{Color: GREEN, Bold: true, Italic: false}
		fmt.Println(ConsoleStyler(fmt.Sprintf("Default browser set to %s \n", FLAGS.SetDefaultBrowser.Value), stylerConfig))
		fmt.Println(ConsoleStyler("Please restart your terminal to apply the changes", stylerConfig))
	} else {
		fmt.Println(ConsoleStyler("Default browser not set", &StylerConfig{Color: GREEN, Bold: true}))
	}
}
