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

def scrape_aldi():
    # Get to test website 
    url = "https://www.aldi.us/en/weekly-specials/our-weekly-ads/"
    
    # Set up Selenium options 
    options = Options()
    #options.add_argument("--headless")
    options.add_argument("--disable-gpu")
    options.add_argument("--no-sandbox")
    options.add_argument("--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
    options.add_argument("--disable-geolocation")
    options.add_experimental_option("prefs", {"profile.default_content_setting_values.geolocation": 2})

    # Installs driver depending on browser
    driver=webdriver.Chrome(service=Service(ChromeDriverManager(version='114.0.5735.90').install()),options=options)

    # Open page
    wait = WebDriverWait(driver, 10)
    driver.maximize_window()
    driver.get(url)

    # Let page load
    time.sleep(5)

    # Need to switch to iframe elements 
    iframe = driver.find_element(By.ID, 'shopLocalPlatformFrame')
    driver.switch_to.frame(iframe)

    # Click search bar 
    iframe_search_bar = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR,'#locationInput')))
    iframe_search_bar.click()
    iframe_search_bar.send_keys("32601")
    iframe_search_bar.send_keys(Keys.ENTER)

    time.sleep(2)

    # Select store
    iframe_select = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, "#StoreListContainer > div:nth-child(1) > div.Nuep__ButtonsContainer-sc-16nxgew-10.iTUGQX > button")))
    iframe_select.click()

    # Allow load
    time.sleep(5)

def main():
    scrape_aldi()

if __name__ == "__main__":
    main()