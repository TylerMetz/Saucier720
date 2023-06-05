# MealDealz

## Frontend Engineers
1. Tyler Metz
2. Edward Menello
## Backend Engineers
1. Sam Forstot
2. Riley Cleavenger

## Project Description
This web app scrapes weekly deals from groccery store websites, and gives users the most cost-effecient recipes based on filters such as what ingredients you already have. Users can manage their `pantry` by adding ingredients they already own.

# How to Use
1. User first must signup if they do not have an existing account.(for testing use username: ri, password: ri)
   1. To signup click the signup link on the login page and fill all associated forms with proper inputs
2. User logs in
3. Can then navigate to pantry/deals/recipes pages
4. `Pantry Page Usage`
   1. The pantry table on the right displays users Ingredients in their pantry
   2. The form on the left lets users add new ingredients into their pantry
5. `Deals Page Usage`
   1. Deals table shows current Ingredients on sale at publix and the deals associated with them
6. `Recipes Page Usage`
   1. Recipe card displays the highest scored recipe with the users current pantry and the current deals at Publix
   2. Can navigate to 29 other recipes with buttons underneath the card

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


