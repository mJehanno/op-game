#!/bin/bash
echo "sed -i -r 's,\"version\": \"([0-9]*\.){2}[0-9]*\",\"version\": \"$1\",g' frontend/package.json"
sed -i -r 's,"version": "([0-9]*\.){2}[0-9]*","version": "'$1'",g' frontend/package.json