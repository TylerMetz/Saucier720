import re
import time
from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver import ActionChains
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.keys import Keys
from webdriver_manager.chrome import ChromeDriverManager

import sys

def scrape_publix():
    # Get to publix website 
    url = "https://www.publix.com/savings/weekly-ad/view-all"
    # Driver options
    if sys.platform.startswith('win'):
        driverPath = "SeleniumDrivers/chromedriver_win32/chromedriver.exe"
    else:
        driverPath = "SeleniumDrivers/chromedriver_mac64/chromedriver"

    # Set up Selenium options 
    options = Options()
    #options.add_argument("--headless")
    options.add_argument("--disable-gpu")
    options.add_argument("--no-sandbox")
    options.add_argument("--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
    options.add_argument("--disable-geolocation")
    options.add_experimental_option("prefs", {"profile.default_content_setting_values.geolocation": 2})

    driver=webdriver.Chrome(service=Service(ChromeDriverManager().install()),options=options)

    # Open page
    wait = WebDriverWait(driver, 10)
    driver.maximize_window()
    driver.get(url)
    
    # Select location 
    location_button = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, "#main > div.savings-content-wrapper.skeleton-spacer > div > div.savings-container.full-bleed > div > div > button > span")))
    
    # Click button and allow time to load 
    location_button.click()
    time.sleep(5)

    # Click search bar
    search_bar = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR,"#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div.search-container > form > input[type=search]")))
    search_bar.click()
    search_bar.send_keys("32601")
    search_bar.send_keys(Keys.ENTER)
    time.sleep(5)

    # Click first store
    store_button = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR,"#\\31 560 > div > div > div.buttons-links > div.p-button-group__wrapper.buttons-wrapper > div > button")))
    store_button.click()
    time.sleep(10)


    page_source = driver.page_source

    # Wait for page to load 
    print(driver.page_source)

    soup = BeautifulSoup(page_source, "html.parser")


    

def main():
    scrape_publix()

if __name__ == "__main__":
    main()
