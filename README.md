# scv

VNC path creator for SAKURA Cloud

## Installation

Download a binary from [here](https://github.com/blp1526/scv/releases).

## Precondition

If you use this cli tool, you have to turn on a SAKURA Cloud server power.

## Usage

* Create a config file at `$HOME/.scv.json`.
* Write a config file refering to [`.scv.sample.json`](.scv.sample.json).
  * This file's server name don't have to match the SAKURA cloud sever name.
* Run below command.

```
# Format
$ scv ZONE_NAME SERVER_NAME

# Example
$ scv is1a ubuntu
```

## Options

|Name|Description|
|-|-|
|-h, --help|print help message and exit|
|-v, --verbose|print debug log|
|-V, --version|print version and exit|
