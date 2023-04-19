# Sprint 4

## Overall Successes 
- [#235](https://github.com/TylerMetz/Saucier720/issues/235) - Cookie generation 

## Backend Successes
- [#71](https://github.com/TylerMetz/Saucier720/issues/71) - Implement Recipe API or Database
- [#206](https://github.com/TylerMetz/Saucier720/issues/206) - Update database after a POST request
- [#207](https://github.com/TylerMetz/Saucier720/issues/207) - Fixed the minor issues of scraping into the database
- [#228](https://github.com/TylerMetz/Saucier720/issues/228) - Current deals sorting 
- [#241](https://github.com/TylerMetz/Saucier720/issues/241) - Backend authentification service
- [#242](https://github.com/TylerMetz/Saucier720/issues/242) - Backend Http Helper Class
- [#244](https://github.com/TylerMetz/Saucier720/issues/244) - Backend Reorganization 
- [#279](https://github.com/TylerMetz/Saucier720/issues/279) - Fix Scraper Time Function
- [#301](https://github.com/TylerMetz/Saucier720/issues/301) - Unnecessary data in recipes
## Frontend Successes

## Backend Failures 
- [#65](https://github.com/TylerMetz/Saucier720/issues/65) - Have BackendPkg import from GitHub
- [#129](https://github.com/TylerMetz/Saucier720/issues/129) - Develop function to calculate price of sale items
- [#130](https://github.com/TylerMetz/Saucier720/issues/130) - Develop function to Calculate the Cost of a Recipe
## Frontend Failures 

## Backend Unit Testing (For new Sprint 4 functionalities)
### GET Request Testing
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

### GoRoutine Cancel Testing
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

### Cookie Generation Testing
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

### Recipes.Json File Read-in Testing
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

### POST Request Testing
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

### Recommended Recipes Testing
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

### Shutdown Servers Testing
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

### Update Database Tables Testing
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

## Frontend Unit Testing (For new Sprint 4 functionalities)
- We are using to conduct our frontend unit tests

### Component Testing
- Our component tests in cypress are

### End to End Testing
- Our end to end tests in cypress are


## Backend API Documentation 
#### Recommendation.go
- **Purpose:** This file holds a `Reccomendation` struct which contains a `R` Recipe, `ItemsInPantry` slice of Fooditems, and `ItemsOnSale` slice of FoodItems. It also contains the functions its basic function `BestRecipes()` which is an algroithm that matches the user pantry to recipes from our dataset. 
  
### Functioning API Flowchart
![Flowchart Image](../Other/Images/Saucier720Api.png)
