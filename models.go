package main

import "fmt"

type StringFlag struct {
	Value  string
	Active bool
}

func (f *StringFlag) Set(value string) error {
	f.Value = value
	f.Active = true
	return nil
}

func (f *StringFlag) String() string {
	return f.Value
}

type Flags struct {
	ListBrowsers      bool
	SetDefaultBrowser StringFlag
	ConfigFile        string
	Help              bool
}

func (f *Flags) Validate() error {
	if f.SetDefaultBrowser.Active {
		if !ALLOWED_BROWSERS[f.SetDefaultBrowser.Value] {
			return fmt.Errorf("Browser %s not supported", f.SetDefaultBrowser.Value)
		}
	}
	return nil
}
