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
from selenium.common.exceptions import TimeoutException
from webdriver_manager.chrome import ChromeDriverManager

def scrape_kroger():
    # Get to test website 
    url = "https://www.kroger.com/weeklyad/shoppable"

    # Set up Selenium options 
    options = Options()
    options.page_load_strategy = 'eager'

    # options.add_argument("--headless")
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
    time.sleep(8)

    """
    # close pop-up
    try:
        popup = driver.find_element(By.CSS_SELECTOR, "#kds-Modal-ln80cktr > button")
        popup.click()
    except:
        print("no popup")
    """
        
    # click location menu
    open_loc = driver.find_element(By.CSS_SELECTOR, "#QuickLinksContainerv2 > div > div.KrogerHeader-ItemV2.KrogerHeader-ModalitySelectorV2.flex.self-center.h-full > div > span > button")
    open_loc.click()
    time.sleep(5)
    # click zipcode
    # open_zip = driver.find_element(By.CSS_SELECTOR, "#root > div > div.Page-outer-block.stack-base > div.ReactModalPortal > div > div > div:nth-child(2) > div > div > button")
    # open_zip.click()
    # time.sleep(5)
    # enter zipcode
    search_bar = driver.find_element(By.CSS_SELECTOR,"#root > div > div.Page-outer-block.stack-base > div.ReactModalPortal > div > div > div:nth-child(2) > form > div > div.PostalCodeSearchBox-inputWrapper.PostalCodeSearchBox-wrapperNew.flex-1 > label > div > input")
    search_bar.click()
    search_bar.send_keys(Keys.BACKSPACE * 6)
    search_bar.send_keys("75080")
    search_button = driver.find_element(By.CSS_SELECTOR,"#root > div > div.Page-outer-block.stack-base > div.ReactModalPortal > div > div > div:nth-child(2) > form > div > div.PostalCodeSearchBox-inputWrapper.PostalCodeSearchBox-wrapperNew.flex-1 > label > div > button")
    time.sleep(100)
    search_button.click()
    time.sleep(10)
    # select store
    change_store = driver.find_element(By.CSS_SELECTOR, "#root > div > div.Page-outer-block.stack-base > div.ReactModalPortal > div > div > div.pb-8 > div:nth-child(5) > div > div > div:nth-child(2) > div.flex.flex-col > div > button")
    change_store.click()
    select_store = driver.find_element(By.CSS_SELECTOR, "#root > div > div.Page-outer-block.stack-base > div.ReactModalPortal > div > div > div.ModalitySelector--StoreSelectionMenu.overflow-y-scroll > div > div:nth-child(1) > div.StoreSelectionMenu-StoreButtonWrapper.flex.flex-row.sm\:flex-col.justify-between.sm\:self-end.mb-auto > div.StoreSelectionMenu-StartButton.text-right.flex.flex-col.sm\:mt-4.self-end.-mt-32 > button")
    select_store.click()
    time.sleep(8)

    # Create a Beautiful Soup object from HTML
    soup = BeautifulSoup(driver.page_source, 'html.parser')

    # Extract the data and remove white space
    products = [span.get_text(strip=True) for span in soup.find_all('span', class_='kds-Text--m SWA-OmniDealDescription2Lines SWA-Clamp2Lines')]
    prices = [div['aria-label'] for div in soup.find_all('div', class_='kds-Text--l SWA-OmniPriceHeading font-heavy font-secondary pl-16 pr-4 truncate mb-8 mt-0 font-bold')]

    # Remove Coupons
    products = [product for product in products if not product.startswith('Save')]

    # Print the extracted data
    for product, price in zip(products, prices):
        if price != "FREE":
            print("Product:", re.sub(r'^\$\d+\.\d+\s', '', product))
            print("Price:", price)
            print()

def main():
    scrape_kroger()

if __name__ == "__main__":
    main()