# Coinspy

Coinspy aids crypto-currency investors by fetching exchange rates and pushing them to devices via Pushover. For fetching exchange rates, multiple providers and markets are supported.

## Usage

`coinspy -h`:

```text
Usage: coinspy [options]

Available options:
  -C, --coins string            Coins to fetch rates for
      --disable-pushover        Disable Pushover notifications
  -F, --fiats string            Fiats to fetch rates for
      --list-providers          List possible providers
      --output-compact          Use compact output format
      --output-very-compact     Use very compact output format
  -P, --providers string        Exchange rate providers to use (default "Cryptowatch/Kraken")
      --pushover-token string   Token for Pushover API access
      --pushover-user string    User for Pushover API access
  -q, --quiet                   Do not print to stdout

For coins and fiats, any well-known symbol (for example BTC for Bitcoin, EUR for Euro) can be used.

Multiple providers, coins and fiats can be provided by using comma-separated lists.
```

### Example

```shell
$ coinspy -P Coingecko,Cryptowatch/Binance -C LTC,BTC,ETH -F USD,EUR,AUD
Error: Market not found (LTC/AUD on Cryptowatch/Binance; 9.716000 allowance remaining)
Error: Market not found (ETH/USD on Cryptowatch/Binance; 9.726000 allowance remaining)
Error: Market not found (LTC/USD on Cryptowatch/Binance; 9.731000 allowance remaining)
Error: Market not found (BTC/USD on Cryptowatch/Binance; 9.701000 allowance remaining)
1 BTC = 58621.000000 AUD (on Coingecko as of 2021-05-17T11:07:33+02:00)
1 BTC = 58720.670000 AUD (on Cryptowatch/Binance as of 2021-05-17T11:07:33+02:00)
1 BTC = 37449.000000 EUR (on Coingecko as of 2021-05-17T11:07:33+02:00)
1 BTC = 37251.630000 EUR (on Cryptowatch/Binance as of 2021-05-17T11:07:33+02:00)
1 BTC = 45492.000000 USD (on Coingecko as of 2021-05-17T11:07:33+02:00)
1 ETH = 4583.470000 AUD (on Coingecko as of 2021-05-17T11:07:33+02:00)
1 ETH = 4571.430000 AUD (on Cryptowatch/Binance as of 2021-05-17T11:07:33+02:00)
1 ETH = 2928.060000 EUR (on Coingecko as of 2021-05-17T11:07:33+02:00)
1 ETH = 2899.530000 EUR (on Cryptowatch/Binance as of 2021-05-17T11:07:33+02:00)
1 ETH = 3556.950000 USD (on Coingecko as of 2021-05-17T11:07:33+02:00)
1 LTC = 372.200000 AUD (on Coingecko as of 2021-05-17T11:07:33+02:00)
1 LTC = 237.770000 EUR (on Coingecko as of 2021-05-17T11:07:33+02:00)
1 LTC = 236.890000 EUR (on Cryptowatch/Binance as of 2021-05-17T11:07:33+02:00)
1 LTC = 288.840000 USD (on Coingecko as of 2021-05-17T11:07:33+02:00)
```

## Supported Providers and Markets

The major providers supported by Coinspy are:

* CoinGate
* Coingecko
* Cryptowatch with all available markets

All providers use free API endpoints, so no credentials are required to fetch exchange rates. Please call `coinspy --list-providers` to obtain a full list of all supported provider/market combinations.

## Configuration

All functionality implemented by Coinspy can be configured by passing CLI arguments. However, an environment file called `.coinspyenv` can also be used to permanently configure Coinspy.

On startup, Coinspy will search the environment file in the current directory or the home directory of the user running Coinspy. If found, Coinspy will load and parse the file, which is expected to be a simple list of _KEY=VALUE_ pairs, separated by newlines. The following table shows the mapping of environment variables to CLI arguments.

| Variable | CLI Argument |
| --- | --- |
| COINSPY_PROVIDERS | -P, --providers |
| COINSPY_COINS | -C, --coins |
| COINSPY_FIATS | -F, --fiats |
| COINSPY_PUSHOVER_USER | --pushover-user |
| COINSPY_PUSHOVER_TOKEN | --pushover-token |
| COINSPY_DISABLE_PUSHOVER | --disable-pushover |
| COINSPY_QUIET | -q, --quiet |
| COINSPY_OUTPUT_COMPACT | --output-compact |
| COINSPY_OUTPUT_VERY_COMPACT | --output-very-compact |

Environment variables can also be passed by actually creating environment variables, for example by calling Coinspy like `COINSPY_PROVIDERS="Coingecko" COINSPY_COINS="BTC" COINSPY_FIATS="EUR" coinspy`.

## Pushover Configuration

In order to actually push exchange rates to a device using Pushover, you will need all of the following:

* A Pushover account.
* Your Pushover user key.
* An application configured in Pushover.
* The API token for that application.

### Pushover Account and User Key

Head over to [Pushover's sign up page](https://pushover.net/signup) and sign up for an account or login via [Pushover's login page](https://pushover.net/login) if you already have an account. Once you are logged in, Pushover will display your user key as part of their dashboard. This is the key required for Coinspy's _--pushover-user_ argument.

### Pushover Application and API Token

In the Pushover dashboard, click on _[Create an Application/API Token](https://pushover.net/apps/build)_. On the following page, fill in the required Application Information and click on _Create Application_. If successful, the API token will be displayed on the next page. This token is required for Coinspy's _--pushover-token_ argument.

## Download

Precompiled binaries for Windows, Linux and macOS are available as _deploy-tagged:archive_ artifacts from the [GitLab CI pipelines for tagged releases](https://gitlab.com/rbrt-weiler/coinspy/-/pipelines?scope=tags).

## Running / Compiling

Coinspy has been developed and tested with [go1.16](https://golang.org/doc/go1.16).

### GNU make

This project includes a Makefile for GNU make. Run `make` to get help; essentially `make build` will produce a binary for your current system in the directory _out_, but the Makefile provides way more functionality.

### Go Toolchain

Use `go run .` to run the tool directly or `go build -o coinspy .` to compile a binary. Dependencies are managed via Go modules and thus automatically handled at compile time.

### Docker-based Builds

Docker can be used to compile binaries by running `docker run --rm -v $PWD:/go/src -w /go/src golang:1.16 go build -o coinspy .`. By passing the `GOOS` and `GOARCH` environment variables (via `-e`) this also enables cross compiling using Docker, for example `docker run --rm -v $PWD:/go/src -w /go/src -e CGO_ENABLED=0 -e GOOS=darwin -e GOARCH=arm64 golang:1.16 go build -o coinspy .` would compile a static binary for ARM-based macOS.

## Source

The original project is [hosted at GitLab](https://gitlab.com/rbrt-weiler/coinspy), with a [copy over at GitHub](https://github.com/rbrt-weiler/coinspy) for the folks over there.
