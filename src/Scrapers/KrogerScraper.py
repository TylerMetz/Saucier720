

import re
import time
from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.keys import Keys
from webdriver_manager.chrome import ChromeDriverManager

def scrape_kroger():
    # Get to test website 
    url = "https://www.kroger.com/weeklyad/shoppable"
    
    # Set up Selenium options 
    options = Options()
    #options.add_argument("--headless")
    options.add_argument("--disable-gpu")
    options.add_argument("--no-sandbox")
    options.add_argument("--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
    options.add_argument("--disable-geolocation")
    options.add_argument("--window-size=1920x1080")
    options.add_experimental_option("prefs", {"profile.default_content_setting_values.geolocation": 2})

    # Installs driver depending on browser
    driver=webdriver.Chrome(service=Service(ChromeDriverManager(version='114.0.5735.90').install()),options=options)

    # Open page
    wait = WebDriverWait(driver, 5)
    driver.get(url)

    # Let page load
    time.sleep(3)


def main():
    scrape_kroger()

if __name__ == "__main__":
    main()