#!/bin/bash

HOOK_DIR=$(git rev-parse --show-toplevel)/.git/hooks
find "$HOOK_DIR" -type l -exec rm {} \;