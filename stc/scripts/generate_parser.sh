#!/usr/bin/env bash

antlr4_dir="/usr/local/lib"
antlr4_version="4.7.1"

java -jar "${antlr4_dir}/antlr-${antlr4_version}-complete.jar" \
  -Dlanguage=Go -o "$PWD/pkg/parser" "$PWD/smalltalk.g4"
