#!/usr/bin/env bash

cd "$(dirname "$0")"

# CAUTION: this script will replace every occurrence of the word
# `passport-chain` in the project folder with whatever argument
# you pass. Be very careful.

if [ $# -lt 1 ] ; then
  printf "Usage:\n$ ./renameProject.sh <my-project-name>\n"
  exit
fi

grep -rl passport-chain ../ --exclude-dir={.git,node_modules} | xargs sed -i s/passport-chain/$1/g
