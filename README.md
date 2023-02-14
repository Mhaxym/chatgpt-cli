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
You can also use the `-list-browsers` flag to see the supported browsers.

Don't forget to restart your terminal or update it with `source ~/.bashrc`.

## Usage

*Only works with .bashrc based shells.*

```
chatgpt-cli --help

Information about the chatgpt command
Usage: chatgpt-cli [options]
Options:
  -browser string
        Opens ChatGPT in the specified browser
  -help
        Shows the help message
  -list-browsers
        Lists the supported browsers
  -set-default-browser string
        Sets the default browser as a variable in the .bashrc file named CHATGPT_DEFAULT_BROWSER
```