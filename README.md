# ChatGPT-Command
Small command to open ChatGPT in your browser. It's useful for people who want to avoid some clicks.

## Installation
To install the command, you need to have [Go installed](https://go.dev/doc/install). Then, run the following command:

```
go install github.com/Mhaxym/chatgpt-cli@latest
```
After that, you need to set the default browser:
```
chatgpt-cli -set-default-browser firefox
```
If you are using a shell other than bash, you need to specify the config file path:
```
chatgpt-cli -set-default-browser firefox -config-file zshrc
```

You can also use the `-list-browsers` flag to see the supported browsers.

Don't forget to restart your terminal or update it with `source ~/.bashrc`.

## Usage

To open ChatGPT in your browser, just run the following command:
```
chatgpt-cli
```

You can also specify the browser you want to use:
```
chatgpt-cli -browser firefox
```

## Help

```
chatgpt-cli --help

Information about the chatgpt-cli command
Usage: chatgpt-cli [options]
Options:
  -browser string
        Opens ChatGPT in the specified browser
  -config-file string
        If your shell environment is not bash, please specify the config file path (e.g. zshrc)
        This file will be used to set the CHATGPT_DEFAULT_BROWSER variable.
        Only use this flag with the --set-default-browser flag. (default "bashrc")
  -help
        Shows the help message
  -list-browsers
        Lists the supported browsers
  -set-default-browser string
        Sets the default browser as a variable in the .bashrc file named CHATGPT_DEFAULT_BROWSER
```