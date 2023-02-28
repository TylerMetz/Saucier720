# Testing Framework

## Running Backend Tests
1. In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src
   ```
2. Run the test function:
   ```
   $ go test
   ```

## Host Testing
- In order to test if data was being correctly outputed to the `localhost:8080` we setup up `func TestOne()` in the `Backend_test.go` file
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    PASS
    ok  	_/<filepath>/Saucier720/src	0.016s
    ```
- If the test <span style = "color:red"> <b>failed</b> </span> the following code block should be returned:
  ```
  --- FAIL: TestOne (0.00s)
    Backend_test.go:63: Result was incorrect, got: <insert desired response>, want: the test to pass.
    FAIL
    exit status 1
    FAIL	_/<filepath>/Saucier720/src	0.020s
    ```