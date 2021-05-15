package core

import (
	"fmt"
	"os"
	"path"

	godotenv "github.com/joho/godotenv"
	pflag "github.com/spf13/pflag"
	envordef "gitlab.com/rbrt-weiler/go-module-envordef"
)

func LoadEnv() {
	// if envFileName exists in the current directory, load it
	localEnvFile := fmt.Sprintf("./%s", EnvFileName)
	if _, localEnvErr := os.Stat(localEnvFile); localEnvErr == nil {
		if loadErr := godotenv.Load(localEnvFile); loadErr != nil {
			Cons.Fprintf(os.Stderr, "Could not load env file <%s>: %s", localEnvFile, loadErr)
		}
	}

	// if envFileName exists in the user's home directory, load it
	if homeDir, homeErr := os.UserHomeDir(); homeErr == nil {
		homeEnvFile := fmt.Sprintf("%s/%s", homeDir, EnvFileName)
		if _, homeEnvErr := os.Stat(homeEnvFile); homeEnvErr == nil {
			if loadErr := godotenv.Load(homeEnvFile); loadErr != nil {
				Cons.Fprintf(os.Stderr, "Could not load env file <%s>: %s", homeEnvFile, loadErr)
			}
		}
	}
}

func SetupFlags() {
	LoadEnv()
	pflag.StringVarP(&Config.Providers, "providers", "P", envordef.StringVal("COINSPY_PROVIDERS", "Cryptowatch/Kraken"), "Exchange rate providers to use")
	pflag.StringVarP(&Config.Coins, "coins", "C", envordef.StringVal("COINSPY_COINS", ""), "Coins to fetch rates for")
	pflag.StringVarP(&Config.Fiats, "fiats", "F", envordef.StringVal("COINSPY_FIATS", ""), "Fiats to fetch rates for")
	pflag.StringVar(&Config.Pushover.Token, "pushover-token", envordef.StringVal("COINSPY_PUSHOVER_TOKEN", ""), "Token for Pushover API access")
	pflag.StringVar(&Config.Pushover.User, "pushover-user", envordef.StringVal("COINSPY_PUSHOVER_USER", ""), "User for Pushover API access")
	pflag.BoolVar(&Config.Disable.Pushover, "disable-pushover", envordef.BoolVal("COINSPY_DISABLE_PUSHOVER", false), "Disable Pushover notifications")
	pflag.BoolVarP(&Config.Quiet, "quiet", "q", envordef.BoolVal("COINSPY_QUIET", false), "Do not print to stdout")
	pflag.BoolVar(&Config.CompactOutput, "output-compact", envordef.BoolVal("COINSPY_OUTPUT_COMPACT", false), "Use compact output format")
	pflag.BoolVar(&Config.VeryCompactOutput, "output-very-compact", envordef.BoolVal("COINSPY_OUTPUT_VERY_COMPACT", false), "Use very compact output format")
	pflag.Usage = func() {
		Cons.Fprintf(os.Stderr, "%s\n", ToolID)
		Cons.Fprintf(os.Stderr, "%s\n", ToolURL)
		Cons.Fprintf(os.Stderr, "\n")
		Cons.Fprintf(os.Stderr, "A tool to fetch exchange rates for crypto coins.\n")
		Cons.Fprintf(os.Stderr, "\n")
		Cons.Fprintf(os.Stderr, "Usage: %s [options]\n", path.Base(os.Args[0]))
		Cons.Fprintf(os.Stderr, "\n")
		Cons.Fprintf(os.Stderr, "Available options:\n")
		pflag.PrintDefaults()
		Cons.Fprintf(os.Stderr, "\n")
		Cons.Fprintf(os.Stderr, "Valid providers are:\n")
		Cons.Fprintf(os.Stderr, "  - CoinGate\n")
		Cons.Fprintf(os.Stderr, "  - Coingecko\n")
		Cons.Fprintf(os.Stderr, "  - Crpytowatch\n")
		Cons.Fprintf(os.Stderr, "\n")
		Cons.Fprintf(os.Stderr, "For coins and fiats, any well-known symbol (for example BTC for Bitcoin, EUR for Euro) can be used.\n")
		Cons.Fprintf(os.Stderr, "\n")
		Cons.Fprintf(os.Stderr, "Multiple providers, coins and fiats can be provided by using comma-seperated lists.\n")
		os.Exit(ErrUsage)
	}
	pflag.Parse()
}

func CheckArguments() {
	if Config.Coins == "" {
		Cons.Fprintf(os.Stderr, "Error: No coins provided.\n")
		os.Exit(ErrGeneric)
	}
	if Config.Fiats == "" {
		Cons.Fprintf(os.Stderr, "Error: No fiats provided.\n")
		os.Exit(ErrGeneric)
	}
	if !Config.Disable.Pushover {
		poTokenLen := len(Config.Pushover.Token)
		poUserLen := len(Config.Pushover.User)
		if poTokenLen > 0 || poUserLen > 0 {
			Config.Pushover.Enabled = true
			if poTokenLen == 0 || poUserLen == 0 {
				Cons.Fprintf(os.Stderr, "Error: Pushover requires token and user.\n")
				os.Exit(ErrGeneric)
			}
		}
	}
}
