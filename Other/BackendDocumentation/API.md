# Backend API Framework

## Router.go
- **Purpose:** This file uses gorillamux to successfully route our backend data to a port to where it can be access by the frontend. For now we are using `localhost:8080` to hold the data that is converted into JSON data, then the frontend accesses it through the use of CORS, which enables data access from an external domain.
<br>
- **Variables**
  **`Name string`** stores the name of the router and is used to indentify what the router's function is in the main file
  **`ItemsToBeEncoded []interface{}`** is a slice passed into the router that stores all of the necessary data to be routed to `localhost:8080`, it can contain any type of data
  <br>
- **Functions**
  **`func (t *Router) Rout()`** creates a new router to send data to `localhost:8080`, enables communication to `localhost:4200` through the use of the CORS Golang package, then infinitely hosts the respective data in `localhost:8080` while waiting for a call to retrieve it on the frontend
  **`func (t *Router) sendResponse(response http.ResponseWriter, request *http.Request)`** is referenced in the `Rout()` function and converts the `ItemsToBeEncoded` slice into JSON using the encoding/json package, then sets it to be written to the correct port
