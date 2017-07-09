# scv

SAKURA Cloud VNC Opener for macOS

## Installation

Download a binary from [here](https://github.com/blp1526/scv/releases).

## Precondition

If you use this cli tool, you have to turn on a SAKURA Cloud server power.

## Usage

* Create a config file at `$HOME/scv.json`.
* Write a config file refering to [`scv.sample.json`](scv.sample.json).
  * This file's server name don't have to match the SAKURA cloud sever name.
* Run below command.

```
# format
$ scv ZONE_NAME SERVER_NAME

# example
$ scv tk1a centos7
```
