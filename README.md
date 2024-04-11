# Coinspy

[![CodeFactor](https://www.codefactor.io/repository/github/rbrt-weiler/coinspy/badge/master)](https://www.codefactor.io/repository/github/rbrt-weiler/coinspy/overview/master)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/rbrt-weiler/coinspy)](https://goreportcard.com/report/gitlab.com/rbrt-weiler/coinspy)
[![build:master](https://img.shields.io/gitlab/pipeline/rbrt-weiler/coinspy/master?label=build%3Amaster)](https://gitlab.com/rbrt-weiler/coinspy/tree/master)
[![pkg.go.dev](https://pkg.go.dev/badge/gitlab.com/rbrt-weiler/coinspy.svg)](https://pkg.go.dev/gitlab.com/rbrt-weiler/coinspy)

Coinspy aids crypto-currency investors by fetching exchange rates and pushing them to devices via Pushover. For fetching exchange rates, multiple providers and markets are supported. Fetched data can also be written to a QuestDB instance, aiding in building a personal, historic repository of exchange rates.

## Usage

`coinspy -h`:

```text
Usage: coinspy [options]

Available options:
  -C, --coins string                  Coins to fetch rates for
      --disable-pushover              Disable Pushover notifications
      --disable-questdb               Disable QuestDB storage
  -F, --fiats string                  Fiats to fetch rates for
      --list-providers                List possible providers
      --livecoinwatch-apikey string   API key for accessing the LiveCoinWatch API
      --output-compact                Use compact output format
      --output-very-compact           Use very compact output format
      --portfolio-value-bottom        Show total portfolio value at bottom of output
      --portfolio-value-top           Show total portfolio value at top of output
  -P, --providers string              Exchange rate providers to use (default "CoinGate")
      --pushover-include-links        Include links to charts in Pushover notifications
      --pushover-token string         Token for Pushover API access
      --pushover-user string          User for Pushover API access
      --questdb-host string           Host running QuestDB
      --questdb-port uint16           Port QuestDB Influx is listening on (default 9009)
  -q, --quiet                         Do not print to stdout

For coins and fiats, any well-known symbol (for example BTC for Bitcoin, EUR for Euro) can be used.
The amount of owned coins can be set by adding =VALUE to the coin, e.g. BTC=1.234.
Multiple providers, coins and fiats can be provided by using comma-separated lists.
```

### Examples

Fetch rates only:

```shell
$ coinspy -C LTC,BTC,ETH -F USD,EUR,AUD
1 BTC = 106689.540000 AUD (on CoinGate as of 2024-04-11T10:55:58Z)
1 BTC = 65409.310000 EUR (on CoinGate as of 2024-04-11T10:55:58Z)
1 BTC = 70332.590000 USD (on CoinGate as of 2024-04-11T10:55:58Z)
1 ETH = 5390.140000 AUD (on CoinGate as of 2024-04-11T10:55:58Z)
1 ETH = 3304.590000 EUR (on CoinGate as of 2024-04-11T10:55:58Z)
1 ETH = 3553.330000 USD (on CoinGate as of 2024-04-11T10:55:58Z)
1 LTC = 146.150000 AUD (on CoinGate as of 2024-04-11T10:55:58Z)
1 LTC = 89.600000 EUR (on CoinGate as of 2024-04-11T10:55:58Z)
1 LTC = 96.340000 USD (on CoinGate as of 2024-04-11T10:55:58Z)
```

Fetch rates and calculate total portfolio value:

```shell
$ coinspy -C LTC=0.654,BTC=1.2345,ETH=3.21 -F EUR,USD,AUD --portfolio-value-top
Total portfolio value: 91414.13 EUR
Total portfolio value: 98294.78 USD
Total portfolio value: 149106.17 AUD
1.234500 BTC = 131708.237130 AUD (on CoinGate as of 2024-04-11T10:56:25Z)
1.234500 BTC = 80747.793195 EUR (on CoinGate as of 2024-04-11T10:56:25Z)
1.234500 BTC = 86825.582355 USD (on CoinGate as of 2024-04-11T10:56:25Z)
3.210000 ETH = 17302.349400 AUD (on CoinGate as of 2024-04-11T10:56:25Z)
3.210000 ETH = 10607.733900 EUR (on CoinGate as of 2024-04-11T10:56:25Z)
3.210000 ETH = 11406.189300 USD (on CoinGate as of 2024-04-11T10:56:25Z)
0.654000 LTC = 95.582100 AUD (on CoinGate as of 2024-04-11T10:56:25Z)
0.654000 LTC = 58.598400 EUR (on CoinGate as of 2024-04-11T10:56:25Z)
0.654000 LTC = 63.006360 USD (on CoinGate as of 2024-04-11T10:56:25Z)
```

## Supported Providers and Markets

The major providers supported by Coinspy are:

* CoinGate
* Coingecko

All providers listed above use free API endpoints, so no credentials are required to fetch exchange rates. In addition, the following providers, which require authentication, are supported:

* LiveCoinWatch (requires _--livecoinwatch-*_)

Please call `coinspy --list-providers` to obtain a full list of all supported provider/market combinations.

## Configuration

All functionality implemented by Coinspy can be configured by passing CLI arguments. However, an environment file called `.coinspyenv` can also be used to permanently configure Coinspy.

On startup, Coinspy will search the environment file in the current directory or the home directory of the user running Coinspy. If found, Coinspy will load and parse the file, which is expected to be a simple list of _KEY=VALUE_ pairs, separated by newlines. The following table shows the mapping of environment variables to CLI arguments.

| Variable | CLI Argument |
| --- | --- |
| COINSPY_PROVIDERS | -P, --providers |
| COINSPY_COINS | -C, --coins |
| COINSPY_FIATS | -F, --fiats |
| COINSPY_LIVECOINWATCH_APIKEY | --livecoinwatch-apikey |
| COINSPY_QUESTDB_HOST | --questdb-host |
| COINSPY_QUESTDB_PORT | --questdb-port |
| COINSPY_DISABLE_QUESTDB | --disable-questdb |
| COINSPY_PUSHOVER_USER | --pushover-user |
| COINSPY_PUSHOVER_TOKEN | --pushover-token |
| COINSPY_PUSHOVER_INCLUDE_LINKS | --pushover-include-links |
| COINSPY_DISABLE_PUSHOVER | --disable-pushover |
| COINSPY_QUIET | -q, --quiet |
| COINSPY_OUTPUT_COMPACT | --output-compact |
| COINSPY_OUTPUT_VERY_COMPACT | --output-very-compact |
| COINSPY_PORTFOLIO_VALUE_TOP | --portfolio-value-top |
| COINSPY_PORTFOLIO_VALUE_BOTTOM | --portfolio-value-bottom |

Environment variables can also be passed by actually creating environment variables, for example by calling Coinspy like `COINSPY_PROVIDERS="Coingecko" COINSPY_COINS="BTC" COINSPY_FIATS="EUR" coinspy`.

## Storing Exchange Rates in a QuestDB Instance

Every time exchange rates are fetched by Coinspy they can also be stored in a [QuestDB](https://questdb.io/) instance. In order to do so, provide the host running QuestDB via _--questdb-host_ and, if necessary, the port exposed by QuestDB via _--questdb-port_. Please note that only the InfluxDB line protocol is supported for writing to QuestDB as of now.

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

### Go Toolchain

Use `go run .` to run the tool directly or `go build -o coinspy .` to compile a binary. Dependencies are managed via Go modules and thus automatically handled at compile time.

### GNU make

This project includes a Makefile for GNU make. Run `make` to get help; essentially `make build` will produce a binary for your current system in the directory _out_, but the Makefile provides way more functionality.

### Docker-based Builds

Docker can be used to compile binaries by running `docker run --rm -v $PWD:/go/src -w /go/src golang:1.16 go build -o coinspy .`. By passing the `GOOS` and `GOARCH` environment variables (via `-e`) this also enables cross compiling using Docker, for example `docker run --rm -v $PWD:/go/src -w /go/src -e CGO_ENABLED=0 -e GOOS=darwin -e GOARCH=arm64 golang:1.16 go build -o coinspy .` would compile a static binary for ARM-based macOS.

## Source

The original project is [hosted at GitLab](https://gitlab.com/rbrt-weiler/coinspy), with a [copy over at GitHub](https://github.com/rbrt-weiler/coinspy) for the folks over there.
