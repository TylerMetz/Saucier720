# Sprint 3

## Overall Successes 
- 

## Backend Successes
- 

## Frontend Successes
- 

## Backend Failures (Pushed to Sprint 4)
- 

## Frontend Failures (Pushed to Sprint 4)
- 

## Backend Unit Testing (For new Sprint 3 functionalities)
#### User Database Testing
- In order to test if user data was being written and read correctly to and from the SQL database we setup up `func TestThree()` in the `Db_test.go` file
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
- If the test <span style = "color:red"> <b>failed</b> </span> one of the following code blocks should be returned:
    ```
    --- FAIL: TestThree (1.324s)
    Db_test.go:57: Pantries do not match.
    FAIL
    FAIL	command-line-arguments	1.324s
    FAIL
    ```
    ```
    --- FAIL: TestThree (1.544s)
    Db_test.go:69: User data does not match.
    FAIL
    FAIL	command-line-arguments	1.544s
    FAIL
    ```
<br>

#### Deals Database Testing
- In order to test if the scraped deals are being written and read correctly to and from the SQL database we setup up `func TestFour()` in the `DealsDb_test.go` file
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
- If the test <span style = "color:red"> <b>failed</b> </span> one of the following code blocks should be returned:
    ```
    --- FAIL: TestFour (153.804s)
    DealsDb_test.go:39: Scrape times don't match.
    FAIL
    FAIL	command-line-arguments	153.804s
    FAIL
    ```
    ```
    --- FAIL: TestFour (135.878s)
    DealsDb_test.go:45: Failed to return item #_ correctly.
    FAIL
    FAIL	command-line-arguments	135.878s
    FAIL
    ```
<br>


## Frontend Unit Testing (For new Sprint 3 functionalities)
- 

## Backend API Documentation 
#### Database.go
- **Purpose:** This file holds a `Database` struct, as well as its basic functions, and utilizes functions from the SQLite and `database\sql` libraries to read and write data from our SQL database.

#### User.go
- **Purpose:** This file holds a `User` struct which contains a `FirstName`, `LastName`, `Email`, `UserName`, `Password`, and `User Pantry`. This data is stored together in the database, then read back to the backend, then routed to the frontend. 
  

## Functioning API Flowchart
![Flowchart Image](../Other/Images/Saucier720Api.png)
