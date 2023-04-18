# Backend Testing Framework

## GET Request Testing
- In order to test if data was being correctly outputed to the `localhost:8080` we setup `func TestGetHandler()` in the `Host_test.go` file
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test Host_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok  	command-line-arguments	0.021s
    ```
<br>

## Web Scraper Testing
- In order to test if data was being correctly webscraped from the Publix website we setup `func TestTwo()` in the `Scrape_test.go` file
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test Scrape_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok  	command-line-arguments	127.572s
    ```
<br>

## User Database Testing
- In order to test if user data was being written and read correctly to and from the SQL database we setup `func TestThree()` in the `Db_test.go` file
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test Db_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok  	command-line-arguments	0.424s
    ```
<br>

## Deals Database Testing
- In order to test if the scraped deals are being written and read correctly to and from the SQL database we setup `func TestFour()` in the `DealsDb_test.go` file
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test DealsDb_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok  	command-line-arguments	143.874s
    ```
<br>

## GoRoutine Cancel Testing
- In order to test if we can cancel ongoing go routines while running other functions using context and cancel, we setup `func TestCancelGoroutine()` in the `CancelGo_test.go` file
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test CancelGo_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok      command-line-arguments  2.243s
    ```
<br>

## Cookie Generation Testing
- To test that cookies could properly be generated on an http request and then be used to validate user data we setup `func TestSetCookie()` in the `GenCookie_test.go` file
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test GenCookie_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok      command-line-arguments  0.442s
    ```
<br>

## Recipes.Json File Read-in Testing
- In order to test that we can properly write .json files to the database and read them back, we setup `func TestGetRecipes()` in the `JsonRead_test.go` file. We used this function in our program to import recipes that are in a json file. 
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test JsonRead_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok      command-line-arguments  0.444s
    ```
<br>

## POST Request Testing
- In order to test that we can receive http POST requests from the frontend and read the data back in, we setup `func TestListenPostReq()` in the `ListenPost_test.go` file.
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test ListenPost_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok      command-line-arguments  0.498s
    ```
<br>

## Recommended Recipes Testing
- In order to test that our algorithm to return recipes based on a user's pantry functioned correctly, we setup `func TestBestRecipes()` in the `RecRecipe_test.go` file.
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test RecRecipe_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok      command-line-arguments  0.422s
    ```
<br>

## Shutdown Servers Testing
- In order to test that we stopped routing previous user data (shutdown servers) after a new user logs in, we setup `func TestShutdownServers()` in the `ShutdownServer_test.go` file.
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test ShutdownServer_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok      command-line-arguments  0.327s
    ```
<br>

## Update Database Tables Testing
- In order to ensure that users' pantries were being properly altered in the database after receiving data from the frontend to update/append/delete items, we setup `func TestUpdatePantry()` in the `UpdateDb_test.go` file.
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src/TestingFiles
   ```
  **2.** Run the test function:
   ```
   $ go test UpdateDb_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok      command-line-arguments  0.470s
    ```
<br>