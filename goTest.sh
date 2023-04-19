#!/bin/bash
# Everytime a new go test file is created/modified just change underneath
cd src/TestingFiles
go test Host_test.go            2> /dev/null
go test Scrape_test.go          2> /dev/null
go test Db_test.go              2> /dev/null
go test CancelGo_test.go        2> /dev/null
go test DealsDb_test.go         2> /dev/null
go test GenCookie_test.go       2> /dev/null
go test Host_test.go            2> /dev/null
go test JsonRead_test.go        2> /dev/null
go test ListenPost_test.go      2> /dev/null
go test RecRecipe_test.go       2> /dev/null
go test ShutdownServer_test.go  2> /dev/null
go test UpdateDb_test.go        2> /dev/null
