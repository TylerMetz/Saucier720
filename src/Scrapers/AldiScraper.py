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
    options.page_load_strategy = 'eager' 
    
    options.add_argument("--headless")
    options.add_argument("--disable-gpu")
    options.add_argument("--no-sandbox")
    options.add_argument("--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
    options.add_argument("--disable-geolocation")
    options.add_argument("--window-size=1920x1080")
    options.add_experimental_option("prefs", {"profile.default_content_setting_values.geolocation": 2})

    # Installs driver depending on browser
    driver = webdriver.Chrome(options=options)

    # Open page
    wait = WebDriverWait(driver, 5)
    driver.get(url)

    # Let page load
    time.sleep(3)

    # Need to switch to iframe elements 
    iframe = driver.find_element(By.ID, 'shopLocalPlatformFrame')
    driver.switch_to.frame(iframe)

    # Click search bar 
    iframe_search_bar = driver.find_element(By.CSS_SELECTOR,'#locationInput')
    iframe_search_bar.click()
    iframe_search_bar.send_keys("32601")
    iframe_search_bar.send_keys(Keys.ENTER)

    time.sleep(2)

    # Select store
    iframe_select = driver.find_element(By.CSS_SELECTOR, "#StoreListContainer > div:nth-child(1) > div.Nuep__ButtonsContainer-sc-16nxgew-10.iTUGQX > button")
    iframe_select.click()

    # Allow load & switch back out of iframe
    time.sleep(3)

    # Select the 'Categories' Button
    select_categories = driver.find_element(By.CSS_SELECTOR,"#StyledMenu > a.nav_cat.sc-16w2z54-1-Menu__MenuItem-bfxgCG.JMikp")
    select_categories.click()
    time.sleep(2)

    # Select the 'Chilled' Button to make list appear and first set of deals load
    select_chilled = driver.find_element(By.CSS_SELECTOR, "#FusionApp > div.sc-1jkt2oq-0-content__BaseContentPage-gVrcAy.kkrzRm > div > div > div > div:nth-child(1) > a")
    select_chilled.click()
    scrape_page(driver, 1)
    time.sleep(1)

    # Tab # for each of the following buttons 
    tab_button = [2,3,4,6]

    # Run the scrape on each page
    for tab in tab_button:
        scrape_page(driver, tab)

    driver.quit()
        
# Function that actually scrapes
def scrape_page(driver: webdriver.Chrome, tab: int):
    # Different conditions for anything not the first page
    if tab != 1:
        wait = WebDriverWait(driver, 5)
        select_page = driver.find_element(By.CSS_SELECTOR, "#DeptNaviLeft > ul > li:nth-child(" + str(tab) + ") > a > div")
        select_page.click()

    time.sleep(1)

    # Grab page html
    page_source = driver.page_source

    # Create Beautiful Soup Object
    soup = BeautifulSoup(page_source, "html.parser")

    # Grab image element 
    img_elements = soup.find_all("img", class_="ListingMini_Image")
    div_elements = soup.find_all("div", class_="sc-169d9gp-0-styled__Deal-kAKkkm fwsHHL")

    # Grab all deals and products 
    for img, div in zip(img_elements, div_elements):
        product = img['alt']
        deal = div.find_all('span')[1]
        deal = deal.get_text(strip=True)
        
        print("Product: ", product)
        print("Deal: ", deal)
        print()

def main():
    scrape_aldi()

if __name__ == "__main__":
    main()