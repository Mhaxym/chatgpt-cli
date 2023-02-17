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
chatgpt-cli --help
```