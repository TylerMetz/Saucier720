## How to Run MealDealz

1. Clone the repository:
   ``` git clone https://github.com/TylerMetz/Saucier720.git```

<br>

2. Set GOPATH as location of main.go

3. Use ``` go get -u <insert package url> ``` to download the necessary packages

4. Launch Selenium:
    ```
    echo $PATH
    echo 'export PATH=$PATH:/src/SeleniumDrivers/chromedriver_mac64' >> ~/.bash_profile
    source ~/.bash_profile
    ```
    If needed install the correct driver from https://chromedriver.chromium.org/downloads and replace "/src/SeleniumDrivers/chromedriver_mac64" with the file path to the driver.