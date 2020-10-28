#!/bin/bash

echo "Time to hustle for the kicks ;)"
echo ""

echo "Deleting all existing profiles and tasks in Phantom..."
rm -f ~/Library/ApplicationSupport/Phantom/ProfileManager.json
rm -f ~/Documents/upper_echelon_180/profiles/*
rm -f ~/Documents/upper_echelon_180/tasks/*
echo ""

cd ~/Documents/upper_echelon_180/upperechelon180/read_cc
node index.js
echo ""

cd ~/Documents/upper_echelon_180/upperechelon180/create_profiles_and_tasks/src
go run main.go

echo ""
echo "Good Luck ;)"
