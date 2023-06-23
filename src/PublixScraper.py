import re
from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.keys import Keys
import sys

def scrape_publix():
    # Get to publix website 
    url = "https://www.publix.com/savings/weekly-ad/view-all"
    # Driver options
    if sys.platform.startswith('win'):
        driverPath = "SeleniumDrivers/chromedriver_win32/chromedriver.exe"
    else:
        driverPath = "SeleniumDrivers/chromedriver_mac64/chromedriver"

    while True:
        # Set up Selenium options 
        options = Options()
        options.add_argument("--headless")
        options.add_argument("--disable-gpu")
        options.add_argument("--no-sandbox")
        options.add_argument("--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")

        # Start Selenium webdriver
        driver = webdriver.Chrome(executable_path = driverPath, options = options)

        # Open page
        driver.get(url)

        # Wait for page to load 
        wait = WebDriverWait(driver, 10)
        print("Waiting for page to load")
        
        # Get page source
        page_source = driver.page_source
        

def main():
    scrape_publix()

if __name__ == "__main__":
    main()