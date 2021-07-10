#!/usr/bin/env bash

replaceTokens() {
  local str=""
  if [ -p /dev/stdin ]; then
    str=$(cat)
  else
    str=$1
  fi
  str=$(printf "%s" "$str" | sed "s+<repoRoot>+$PWD+g")
  str=$(printf "%s" "$str" | sed "s+<home>+$HOME+g")
  printf "%s" "$str"
}

stat $1 &> /dev/null && cat $1 | replaceTokens && exit 0
