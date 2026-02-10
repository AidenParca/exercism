#!/usr/bin/env bash
echo "$@" | grep -o . | tac | tr -d '\n' ; echo
