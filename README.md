# CoinSpy

CoinSpy aids crypto currency investors by fetching market prices and pushing them to devices via Pushover.

**Currently in development and not feature-complete.**

## Usage

`coinspy -h`:

```text
Usage: coinspy [options]

Available options:
  -C, --coins string            Coins to fetch rates for
  -F, --fiats string            Fiats to fetch rates for
  -M, --markets string          Markets to use with multi-market providers (comma-seperated) (default "Kraken")
      --pushover-token string   Token for Pushover API access
      --pushover-user string    User for Pushover API access
  -q, --quiet                   Do not print to stdout
```

### Example

```shell
$ coinspy -M Kraken,Binance -C BTC,ETH -F EUR,USD
Error: Market not found (Binance ETH/USD; 9.900000 allowance remaining)
Error: Market not found (Binance BTC/USD; 9.890000 allowance remaining)
1 BTC = 45231.910000 EUR (on Binance as of 2021-04-20T08:40:54+02:00)
1 BTC = 45196.700000 EUR (on Kraken as of 2021-04-20T08:40:54+02:00)
1 BTC = 54330.100000 USD (on Kraken as of 2021-04-20T08:40:54+02:00)
1 ETH = 1737.210000 EUR (on Binance as of 2021-04-20T08:40:54+02:00)
1 ETH = 1734.000000 EUR (on Kraken as of 2021-04-20T08:40:54+02:00)
1 ETH = 2087.210000 USD (on Kraken as of 2021-04-20T08:40:54+02:00)
```

## Running / Compiling

Use `go run .` to run the tool directly or `go build -o coinspy .` to compile a binary. Dependencies are managed via Go modules and thus automatically handled at compile time.

Alternatively, Docker can be used to compile binaries by running `docker run --rm -v $PWD:/go/src -w /go/src golang:1.16 go build -o coinspy .`. By passing the `GOOS` and `GOARCH` environment variables (via `-e`) this also enables cross compiling using Docker, for example `docker run --rm -v $PWD:/go/src -w /go/src -e CGO_ENABLED=0 -e GOOS=darwin -e GOARCH=arm64 golang:1.16 go build -o coinspy .` would compile a static binary for ARM-based macOS.

Tested with [go1.16](https://golang.org/doc/go1.16).

## Source

The original project is [hosted at GitLab](https://gitlab.com/rbrt-weiler/coinspy), with a [copy over at GitHub](https://github.com/rbrt-weiler/coinspy) for the folks over there.
