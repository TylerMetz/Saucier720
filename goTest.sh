#!/bin/bash
# Everytime a new go test file is created/modified just change underneath
cd src/TestingFiles
go test Host_test.go
go test Scrape_test.go
go test Db_test.go
go test DealsDb_test.go
