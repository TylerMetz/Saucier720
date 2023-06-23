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
        options = Options()
        options.add_argument("--headless")
        options.add_argument("--disable-gpu")
        options.add_argument("--no-sandbox")


def main():
    scrape_publix()

if __name__ == "__main__":
    main()