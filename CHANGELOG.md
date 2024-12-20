# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

1. The table that is written to within QuestDB can now be specified with `--questdb-table`.
1. A timeout for QuestDB connections can now be specified with `--questdb-timeout`.
1. Support for storing exchange rates in a SQLite3 database with `--sqlite3-*`.
1. Support for storing exchange rates in a DuckDB database with `--duckdb-*`.
1. Support for including the name of the host that sent the Pushover notification with `--pushover-include-host`.

### Changed

1. QuestDB connections by default time out after 10 seconds with an appropriate error message.
1. Default table name has been changed from _exchange_rates_ to _crypto_rates_ for all database backends. **This may break your setup!**

## [1.2.4] - 2024-04-19

### Added

1. Support for the Bitpanda API.

## [1.2.3] - 2024-04-11

### Removed

1. Support for Cryptowatch, as the API was shut down in 2023.

## [1.2.2] - 2021-07-06

### Added

1. Option to include links to charts in Pushover notifications with `--pushover-include-links`.

## [1.2.1] - 2021-07-06

### Changed

1. Set URL title for Pushover messages.

## [1.2.0] - 2021-07-04

### Added

1. Support for calculating and displaying the total portfolio value.

## [1.1.0] - 2021-05-26

### Added

1. Support for storing results in a QuestDB instance.
1. New provider LiveCoinWatch.

### Fixed

1. Error handling for Pushover now works.
1. Pushover messages are checked for length and split up into multiple messages if necessary.

## [1.0.0] - 2021-05-17

### Added

1. Configuration via environment file.
1. New provider Coingecko.
1. Support for multiple providers at once.
1. Options for shorter output formats.
1. New provider CoinGate.
1. New option `--disable-pushover`.
1. New option `--list-providers`.

## [0.2.0] - 2021-05-08

### Added

1. Basic Pushover functionality.
1. Quiet mode.

## [0.1.0] - 2021-04-18 (code) and 2021-04-19/20 (documentation)

Initial public release.

### Added

1. Fetch exchange rates for an arbitary number of coins and fiats from Cryptowatch.
1. Support for multiple markets as supported by Cryptowatch.

[Unreleased]: https://gitlab.com/rbrt-weiler/coinspy/-/compare/1.2.4...master
[1.2.4]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.2.4
[1.2.3]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.2.3
[1.2.2]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.2.2
[1.2.1]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.2.1
[1.2.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.2.0
[1.1.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.1.0
[1.0.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.0.0
[0.2.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/0.2.0
[0.1.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/0.1.0
