# Backend Testing Framework

## Host Testing
- In order to test if data was being correctly outputed to the `localhost:8080` we setup up `func TestOne()` in the `Host_test.go` file
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src
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
- If the test <span style = "color:red"> <b>failed</b> </span> the following code block should be returned:
    ```
    --- FAIL: TestOne (0.00s)
    Host_test.go:64: Result was incorrect, got: <data read from localhost>, want: <data translated to JSON>.
    FAIL
    FAIL	command-line-arguments	0.017s
    FAIL
    ```
<br>

## Web Scraper Testing
- In order to test if data was being correctly webscraped from the Publix website we setup up `func TestTwo()` in the `Scrape_test.go` file
  <br>
- **How to Test:**
  **1.** In the command line naviagte to the `src` folder:
   ```
   $ cd Saucier720/src
   ```
  **2.** Run the test function:
   ```
   $ go test Scrape_test.go
   ```
   <br>
-  If the test <span style = "color:green"> <b>passed</b> </span> the following code block should be returned:
    ```
    ok  	command-line-arguments	127.572s
    ```
- If the test <span style = "color:red"> <b>failed</b> </span> the following code block should be returned:
    ```
    --- FAIL: TestTwo (129.93s)
    Scrape_test.go:30: Data was not scraped.
    FAIL
    FAIL	command-line-arguments	129.949s
    FAIL
    ```
<br>