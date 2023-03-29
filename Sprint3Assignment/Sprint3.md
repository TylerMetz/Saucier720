# Sprint 3

## Overall Successes 
- [#138](https://github.com/TylerMetz/Saucier720/issues/138) - Make an extensive How-to-run file
- [#137](https://github.com/TylerMetz/Saucier720/issues/137) - Create a script to run all necessary start files
- [#143](https://github.com/TylerMetz/Saucier720/issues/143) - Clean up Branches

## Backend Successes
- [#112](https://github.com/TylerMetz/Saucier720/issues/112) - Fully Implement Web Scraper
- [#189](https://github.com/TylerMetz/Saucier720/issues/189) - Develop Unit Tests for Sprint 3 functions
- [#207](https://github.com/TylerMetz/Saucier720/issues/207) - First Item in Database scraped incorrectly
- [#192](https://github.com/TylerMetz/Saucier720/issues/188) - Figure out how to rout to multiple localhost ports at once
- [#176](https://github.com/TylerMetz/Saucier720/issues/176) - Bash Test Script
- [#131](https://github.com/TylerMetz/Saucier720/issues/131) - Implement Database
- [#128](https://github.com/TylerMetz/Saucier720/issues/128) - Create functions to parse through Publix data
- [#170](https://github.com/TylerMetz/Saucier720/issues/170) - Backend Submodules
- [#119](https://github.com/TylerMetz/Saucier720/issues/119) - Figure out how to include other Packages in our repo
- [#132](https://github.com/TylerMetz/Saucier720/issues/132) - Create a User Data Struct

## Frontend Successes
- [#83](https://github.com/TylerMetz/Saucier720/issues/83) - Create Deals Page
- [#93](https://github.com/TylerMetz/Saucier720/issues/93) - Create login page
- [#95](https://github.com/TylerMetz/Saucier720/issues/95) - Meet with Backend to see what they are planning to store so we can form a wireframe 
- [#97](https://github.com/TylerMetz/Saucier720/issues/97) - Create wireframe for Deals Page
- [#104](https://github.com/TylerMetz/Saucier720/issues/104) - Create dashboard Page
- [#156](https://github.com/TylerMetz/Saucier720/issues/156) - Further pantry functionality
- [#173](https://github.com/TylerMetz/Saucier720/issues/173) - Use bootstrap for styling
- [#192](https://github.com/TylerMetz/Saucier720/issues/192) - Deals Page Update
- [#194](https://github.com/TylerMetz/Saucier720/issues/194) - Deals Page Bootstrap Update
- [#195](https://github.com/TylerMetz/Saucier720/issues/195) - Pantry Page Bootstrap Update
- [#196](https://github.com/TylerMetz/Saucier720/issues/196) - Login Page Button
- [#203](https://github.com/TylerMetz/Saucier720/issues/203) - New HTTP Requests
- [#204](https://github.com/TylerMetz/Saucier720/issues/204) - Post New Pantry Item
- [#208](https://github.com/TylerMetz/Saucier720/issues/208) - Post to Database from Frontend
- [#221](https://github.com/TylerMetz/Saucier720/issues/221) - Implement end to end testing

## Backend Failures 

- [#117](https://github.com/TylerMetz/Saucier720/issues/117) - Push Gorillamux data to non-local server port 
  - discovered this isn't needed
- [#130](https://github.com/TylerMetz/Saucier720/issues/130) - Develop Function to Calculate the Cost of a Recipe
  - Pushed to Sprint 4
- [#71](https://github.com/TylerMetz/Saucier720/issues/71) - Implement Recipe API or Database
  - Pushed to Sprint 4
- [#129](https://github.com/TylerMetz/Saucier720/issues/129) - Develop function to calculate price of sale items
  - Pushed to Sprint 4
- [#65](https://github.com/TylerMetz/Saucier720/issues/65) - Have BackendPkg import from GitHub
  - Pushed to Sprint 4, not a priority
- [#206](https://github.com/TylerMetz/Saucier720/issues/206) - Update Database after POST request
  - Working on it, but it isn't fully functioning yet so pushed to Sprint 4

## Frontend Failures (Pushed to Sprint 4)
- [#82](https://github.com/TylerMetz/Saucier720/issues/221) - Create list page
- [#84](https://github.com/TylerMetz/Saucier720/issues/221) - Create recipes page
- [#96](https://github.com/TylerMetz/Saucier720/issues/221) - Create wireframe for list page
- [#101](https://github.com/TylerMetz/Saucier720/issues/221) - Create wireframe for recipes page


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
- We are using to conduct our frontend unit tests

### Component Testing
- Our component tests in cypress are
  - Pantry service should get Pantry in [pantry.service.spec.cy.ts](../Saucier720-app/src/app/testing/pantry/pantry.service.spec.cy.ts)
  - Pantry component displays table, post button, and gets pantry in [pantry.component.spec.cy.ts](../Saucier720-app/src/app/testing/pantry/pantry.component.spec.cy.ts)
  - Deals component displays table in [deals.component.spec.cy.ts](../Saucier720-app/src/app/testing/deals/deals.component.spec.cy.ts)
### End to End Testing
- Our end to end tests in cypress are
  - Requests deals from backend database and receives status code `200` if properly received and responsed to in [deals.http.spec.cy.ts](../Saucier720-app/cypress/e2e/HttpRequests/deals.http.spec.cy.ts)
  - Requests user pantry from backend database and recevies status code 200 if properly received[pantry.http.spec.cy.ts](../Saucier720-app/cypress/e2e/HttpRequests/pantry.http.spec.cy.ts)
  - Posts new FoodItem from the frontend to the backend database and receives status code 200 if properly received and posted to database[pantry.http.spec.cy.ts](../Saucier720-app/cypress/e2e/HttpRequests/pantry.http.spec.cy.ts)

## Backend API Documentation 
#### Database.go
- **Purpose:** This file holds a `Database` struct, as well as its basic functions, and utilizes functions from the SQLite and `database\sql` libraries to read and write data from our SQL database.

#### User.go
- **Purpose:** This file holds a `User` struct which contains a `FirstName`, `LastName`, `Email`, `UserName`, `Password`, and `User Pantry`. This data is stored together in the database, then read back to the backend, then routed to the frontend. 
  

## Functioning API Flowchart
![Flowchart Image](../Other/Images/Saucier720Api.png)
