#!/bin/bash
# Everytime a new go test file is created/modified just change underneath
cd src/TestingFiles
go test Host_test.go
go test Scrape_test.go
go test Db_test.go
go test CancelGo_test.go
go test DealsDb_test.go
go test GenCookie_test.go
go test Host_test.go
go test JsonRead_test.go
go test ListenPost_test.go
go test RecRecipe_test.go
go test ShutdownServer_test.go
go test UpdateDb_test.go
