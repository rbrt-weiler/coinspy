# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

(Currently nothing to note.)

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

[Unreleased]: https://gitlab.com/rbrt-weiler/coinspy/-/compare/1.1.0...master
[1.2.2]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.2.2
[1.2.1]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.2.1
[1.2.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.2.0
[1.1.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.1.0
[1.0.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/1.0.0
[0.2.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/0.2.0
[0.1.0]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/0.1.0
