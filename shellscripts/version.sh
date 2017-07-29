#!/bin/bash
git describe --tags | awk -F '-' '{ gsub("v", "", $1); print $1 }'
