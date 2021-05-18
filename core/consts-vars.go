package core

import (
	consolehelper "gitlab.com/rbrt-weiler/go-module-consolehelper"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	// ToolName contains the name of the application.
	ToolName string = "Coinspy"
	// ToolVersion contains the version number of the application.
	ToolVersion string = "1.0.0"
	// ToolID contains a user-agent-style representation of ToolName/ToolVersion.
	ToolID string = ToolName + "/" + ToolVersion
	// ToolURL contains the URL where the application can be found.
	ToolURL string = "https://gitlab.com/rbrt-weiler/coinspy"

	// EnvFileName contains the expected name for the env file.
	EnvFileName string = ".coinspyenv"
	// LineBreak defines the format of linebreaks in the Pushover message.
	LineBreak string = "\r\n"
	// PushoverMesssageLength defines the maximum length of a Pushover message.
	PushoverMessageLength int = 1024

	// ErrSuccess means no error.
	ErrSuccess int = 0
	// ErrGeneric represents an generic error.
	ErrGeneric int = 1
	// ErrUsage is thrown when the usage message is shown.
	ErrUsage int = 2
)

var (
	// Config stores the parsed CLI arguments.
	Config types.AppConfig
	// Cons is a helper variable used for printing output.
	Cons consolehelper.ConsoleHelper
)
