package core

import (
	"os"
	"path"

	pflag "github.com/spf13/pflag"
)

func SetupFlags() {
	//pflag.StringVarP(&config.Provider, "provider", "P", "Cryptowatch", "Exchange rate provider to use")
	pflag.StringVarP(&Config.Markets, "markets", "M", "Kraken", "Markets to use with multi-market providers (comma-seperated)")
	pflag.StringVarP(&Config.Coins, "coins", "C", "", "Coins to fetch rates for")
	pflag.StringVarP(&Config.Fiats, "fiats", "F", "", "Fiats to fetch rates for")
	pflag.StringVar(&Config.Pushover.Token, "pushover-token", "", "Token for Pushover API access")
	pflag.StringVar(&Config.Pushover.User, "pushover-user", "", "User for Pushover API access")
	pflag.BoolVarP(&Config.Quiet, "quiet", "q", false, "Do not print to stdout")
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
