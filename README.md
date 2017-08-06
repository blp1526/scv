# scv

VNC path creator for [SAKURA Cloud](http://cloud.sakura.ad.jp)

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

This command uses SAKURA Cloud API [GET/server/:serverid/vnc/proxy](http://developer.sakura.ad.jp/cloud/api/1.1/server/#get_server_serverid_vnc_proxy).

## Options

|Name|Description|
|-|-|
|-h, --help|print help message and exit|
|-v, --verbose|print debug log|
|-V, --version|print version and exit|
