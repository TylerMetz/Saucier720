# MealDealz

## Frontend Engineers
1. Tyler Metz
2. Edward Menello
## Backend Engineers
1. Sam Forstot
2. Riley Cleavenger

## Project Description
This web app scrapes weekly deals from groccery store websites, and gives users the most cost-effecient recipes based on filters such as serving size, dietary restrictions, and what ingredients you already have. Users can manage their "pantry" by adding items/ingredients they already own. Prices is calculated based on the needed ingredients, taking into account what the user already has. 

# How to Use


# How to Run
1. Initialize the submodules:

    ```
    $ git submodule update --init --recursive
    ```
    <br>


2. Set Go Environment variables
   
   - Set **GO111MODULE** to off:
     - On Windows:
        ```
        $ SET GO111MODULE=off
        ```
     - On MacOS:
        ```
        $ export GO111MODULE="off"
        ```
   - Enure **GOROOT** and **GOPATH** are the correct locations:
     - **GOPATH** must be the folder of the Saucier720 local repository
     - **GOROOT** must be the GO install location
       - On Windows:
         - ```C:\Program Files\go``` or ```C:\Program Files (x86)\go```
       - On bash (MacOS):
         - ``` /usr/local/Cellar/go/1.19.5/libexec ```
    - **Note:** On Windows you have to set all three variables within "Edit the system environmental controls" > "Environmental Variables" 
<br>


3. Install all external packages to the **GOROOT**
    ```
    $ cd src
    ```
    ```
    $ go get -u github.com/gorilla/mux
    ```
    ```
    $ go get -u github.com/rs/cors
    ```
    ```
    $ go get -u github.com/tebeka/selenium
    ```
    ```
    $ go get -u golang.org/x/exp/slices
    ```
    <br>

4. Install Internal Package
   ```
   $ cd src
   ```
   ```
   $ go install BackendPkg
   ```
   <br>

5. [Install Node.js](https://nodejs.org/en/download/)
   <br>

6. Run in Saucier720 folder:
   ``` 
   $ npm install -g @angular/cli
   ```
   ```
   $ cd Saucier720-app
   ```
   ```
   $ npm install
   ```
   <br>

7. Run the program:
   ```
   $ cd Saucier720
   ```
   ```
   $ ./build.sh
   ```


