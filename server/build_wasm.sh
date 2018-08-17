#!/bin/sh

GOARCH=wasm GOOS=js go build -o mkkeypair.wasm ../keypair/mkkeypair
