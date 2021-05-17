package core

import (
	consolehelper "gitlab.com/rbrt-weiler/go-module-consolehelper"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	// ToolName contains the name of the application.
	ToolName string = "coinspy"
	// ToolVersion contains the version number of the application.
	ToolVersion string = "0.2.0"
	// ToolID contains a user-agent-style representation of ToolName/ToolVersion.
	ToolID string = ToolName + "/" + ToolVersion
	// ToolURL contains the URL where the application can be found.
	ToolURL string = "https://gitlab.com/rbrt-weiler/coinspy"

	// EnvFileName contains the expected name for the env file.
	EnvFileName string = ".coinspyenv"

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
