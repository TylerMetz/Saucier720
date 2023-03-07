# How to Initialize Submodules
1. Initialize the submodules:

```
$ git submodule init
```
<br>

2. Update the submodules:
```
$ git submodule update
```
<br>

3. Set GO111MODULE to off:
- On Windows:
     ```
     $ SET GO111MODULE=off
     ```
- On MacOS:
    ```
    $ export GO111MODULE="off"
    ```
<br>

4. Enure **GOROOT** and **GOPATH** are the correct locations:
   - **GOPATH** must be the folder of the Saucier720 local repository
   - **GOROOT** must be the GO install location
     - On Windows:
       - ```C:\Program Files\go``` or ```C:\Program Files (x86)\go```
     - On MacOS:
       - ``` /usr/local/Cellar/go/1.19.5/libexec ```
<br>
5. Install all external packages to the **GOROOT**:
    ```
    $ go get -u github.com/gorilla/mux
    ```
    ```
    $ go get -u github.com/rs/cors
    ```
  