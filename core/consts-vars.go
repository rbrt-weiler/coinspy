package core

import (
	consolehelper "gitlab.com/rbrt-weiler/go-module-consolehelper"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	ToolName    string = "coinspy"
	ToolVersion string = "0.1.0"
	ToolID      string = ToolName + "/" + ToolVersion
	ToolURL     string = "https://gitlab.com/rbrt-weiler/coinspy"

	ErrSuccess int = 0
	ErrGeneric int = 1
	ErrUsage   int = 2
)

var (
	Config types.AppConfig
	Cons   consolehelper.ConsoleHelper
)