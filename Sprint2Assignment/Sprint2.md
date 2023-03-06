# Sprint 2

## Overall Successes 
- [#87](https://github.com/TylerMetz/Saucier720/issues/87) - Frontend-backend communication
- [#120](https://github.com/TylerMetz/Saucier720/issues/120) - Draft api documentation 

## Backend Successes
- [#94](https://github.com/TylerMetz/Saucier720/issues/94) - Found new way to webscrape 
- [#70](https://github.com/TylerMetz/Saucier720/issues/70) - Completed webscraping
- [#67](https://github.com/TylerMetz/Saucier720/issues/67) - Added a quantity variable 
- [#116](https://github.com/TylerMetz/Saucier720/issues/116) - Converted golang data to JSON
- [#111](https://github.com/TylerMetz/Saucier720/issues/111) - Set up Gorillamux 
- [#113](https://github.com/TylerMetz/Saucier720/issues/113) - Cleaned up webscraping code
- [#108](https://github.com/TylerMetz/Saucier720/issues/108) - Got frontend files into a backend branch 
- [#136](https://github.com/TylerMetz/Saucier720/issues/136) - Unit testing 
- [#118](https://github.com/TylerMetz/Saucier720/issues/118) - Sprint 2 documentation 

## Frontend Successes
- [#90](https://github.com/TylerMetz/Saucier720/issues/90) - setup Unit Testing
- [#105](https://github.com/TylerMetz/Saucier720/issues/105) - create HTML Client
- [#106](https://github.com/TylerMetz/Saucier720/issues/106) - create HTML Client tests

## Backend Failures (Pushed to Sprint 3)
- [#65](https://github.com/TylerMetz/Saucier720/issues/65) - Have backendpkg import from github
- [#117](https://github.com/TylerMetz/Saucier720/issues/117) - Push Gorillamux to non-local server port
- [#71](https://github.com/TylerMetz/Saucier720/issues/71) - Implement recipe API / database 

## Frontend Failures (Pushed to Sprint 3)
- [#93](https://github.com/TylerMetz/Saucier720/issues/93) - User login page
- [#82](https://github.com/TylerMetz/Saucier720/issues/82) - Create list page
- [#83](https://github.com/TylerMetz/Saucier720/issues/83) - Create deals page
- [#84](https://github.com/TylerMetz/Saucier720/issues/84) - Create recipes page



## Backend Unit Testing (For new Sprint 2 functionalities)
- `Host_test.go` - Test to ensure that gorillamux is outputting correct data to server port that can be read back in, tests the router functions. 

- `Scrape_test.go` - Test to ensure that data was scraped from public.com , tests the scraper functions 

## Frontend Unit Testing (For new Sprint 2 functionalities)
- currently only have proper tests for `pantry` and `pantryService` in `../Saucier720-app/src/app/testing/pantry`
- in `PantryService` it simulates receiving a `GET` Http request in its function of `getPantry` and then receives the mock data from the backend
- in `PantryComponent` it simulates loading the component. When the component loads it makes a Http Reqeust(the same as in PantryService) so the testing for PantryComponent tests the exact same as the previous PantryService test

## Backend API Documentation 
### Router.go
- **Purpose:** This file uses gorillamux to successfully route our backend data to a port to where it can be access by the frontend. For now we are using `localhost:8080` to hold the data that is converted into JSON data, then the frontend accesses it through the use of CORS, which enables data access from an external domain.

- **Variables**
  **`Name string`** stores the name of the router and is used to indentify what the router's function is in the main file
  **`ItemsToBeEncoded []interface{}`** is a slice passed into the router that stores all of the necessary data to be routed to `localhost:8080`, it can contain any type of data

- **Functions**
  **`func (t *Router) Rout()`** creates a new router to send data to `localhost:8080`, enables communication to `localhost:4200` through the use of the CORS Golang package, then infinitely hosts the respective data in `localhost:8080` while waiting for a call to retrieve it on the frontend
  **`func (t *Router) sendResponse(response http.ResponseWriter, request *http.Request)`** is referenced in the `Rout()` function and converts the `ItemsToBeEncoded` slice into JSON using the encoding/json package, then sets it to be written to the correct port

## Functioning API Flowchart
![Flowchart Image](../Other/Images/Saucier720Api.png)
