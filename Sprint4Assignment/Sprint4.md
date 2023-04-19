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
- [#84](https://github.com/TylerMetz/Saucier720/issues/84) - Create Recipes Page
- [#98](https://github.com/TylerMetz/Saucier720/issues/98) - Create wireframe for recipes page 
- [#230](https://github.com/TylerMetz/Saucier720/issues/230) - Further HTML Component Testing
- [#232](https://github.com/TylerMetz/Saucier720/issues/232) - Deprecation
- [#234](https://github.com/TylerMetz/Saucier720/issues/234) - Use Logins an STAYS Logged in
- [#235](https://github.com/TylerMetz/Saucier720/issues/235) - Cookie Generation
- [#287](https://github.com/TylerMetz/Saucier720/issues/287) - Stay on same page after reset
- [#289](https://github.com/TylerMetz/Saucier720/issues/289) - Recipe cards
- [#300](https://github.com/TylerMetz/Saucier720/issues/300) - Backend update database
- [#314](https://github.com/TylerMetz/Saucier720/issues/314) - Create sub recipe cards on recipe cards
- [#324](https://github.com/TylerMetz/Saucier720/issues/324) - Delete last empty bullet point on recipe card
- [#341](https://github.com/TylerMetz/Saucier720/issues/341) - Highlight / invert wehn hovering over a button
- [#342](https://github.com/TylerMetz/Saucier720/issues/342) - Add box behind the deals on the deals page
- [#343](https://github.com/TylerMetz/Saucier720/issues/343) - Buttons on Pantry design
- [#345](https://github.com/TylerMetz/Saucier720/issues/345) - Adding an index to the recommended recipe page
- [#346](https://github.com/TylerMetz/Saucier720/issues/346) - Number the steps of the recipes

## Backend Failures 
- [#65](https://github.com/TylerMetz/Saucier720/issues/65) - Have BackendPkg import from GitHub
- [#129](https://github.com/TylerMetz/Saucier720/issues/129) - Develop function to calculate price of sale items
- [#130](https://github.com/TylerMetz/Saucier720/issues/130) - Develop function to Calculate the Cost of a Recipe

## Frontend Failures (Pushed to Sprint 5)
- [#82](https://github.com/TylerMetz/Saucier720/issues/82) - Create list page 
- [#96](https://github.com/TylerMetz/Saucier720/issues/96) - Create wireframe for list page
- [#101](https://github.com/TylerMetz/Saucier720/issues/101) - Learn github autodeployment and workflows
- [#347](https://github.com/TylerMetz/Saucier720/issues/347) - Information about how each pagworks

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
  - Pantry service should get Pantry in [pantry.service.spec.cy.ts](../Saucier720-app/src/app/testing/pantry/pantry.service.spec.cy.ts)
  - Pantry component displays table, post button, and gets pantry in [pantry.component.spec.cy.ts](../Saucier720-app/src/app/testing/pantry/pantry.component.spec.cy.ts)
  - Deals component displays table in [deals.component.spec.cy.ts](../Saucier720-app/src/app/testing/deals/deals.component.spec.cy.ts)

### End to End Testing
- Our end to end tests in cypress are
  - Requests deals from backend database and receives status code `200` if properly received and responsed to in [deals.http.spec.cy.ts](../Saucier720-app/cypress/e2e/HttpRequests/deals.http.spec.cy.ts)
  - Requests user pantry from backend database and recevies status code 200 if properly received[pantry.http.spec.cy.ts](../Saucier720-app/cypress/e2e/HttpRequests/pantry.http.spec.cy.ts)
  - Posts new FoodItem from the frontend to the backend database and receives status code 200 if properly received and posted to database[pantry.http.spec.cy.ts](../Saucier720-app/cypress/e2e/HttpRequests/pantry.http.spec.cy.ts)


## Backend API Documentation 
#### Recommendation.go
- **Purpose:** This file holds a `Recommendation` struct which contains a `R` Recipe, `ItemsInPantry` slice of Fooditems, and `ItemsOnSale` slice of FoodItems. It also contains the functions its basic function `BestRecipes()` which is an algroithm that matches the user pantry to recipes from our dataset. 
  
### Functioning API Flowchart
![Flowchart Image](../Other/Images/Saucier720Api.png)
