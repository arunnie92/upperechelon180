#!/bin/bash

echo "Deleting all existing profiles in Phantom..."
rm -f ~/Library/ApplicationSupport/Phantom/ProfileManager.json
rm -f ~/Documents/upper_echelon_180/profiles/*
echo ""

cd ~/Documents/upper_echelon_180/upperechelon180/read_cc
node index.js
echo ""

cd ~/Documents/upper_echelon_180/upperechelon180/create_profiles/src
go run main.go
