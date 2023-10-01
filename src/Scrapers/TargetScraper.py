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

def scrape_target():
    # Get to test website 
    url = "https://www.target.com/c/grocery-deals/-/N-k4uyq"
    
    # Set up Selenium options 
    options = Options()
    options.page_load_strategy = 'eager' # might have to delete bc lazy loading

    options.add_argument("--headless")
    options.add_argument("--disable-gpu")
    options.add_argument("--no-sandbox")
    options.add_argument("--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
    options.add_argument("--disable-geolocation")
    options.add_experimental_option("prefs", {"profile.default_content_setting_values.geolocation": 2})

    # Installs driver depending on browser
    driver = webdriver.Chrome(options=options)
    
    # Open page
    wait = WebDriverWait(driver, 10)
    driver.maximize_window()
    driver.get(url)
    driver.execute_script("document.body.style.zoom='25%'")
    
    # Give time to load
    time.sleep(3)
    scrape_page(driver, True)
    
    time.sleep(1)
    # Get every other set of deals, stop when page not clickable
    while True:
        try:                                                                                        
            load_more = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR,"#pageBodyContainer > div > div:nth-child(1) > div > div:nth-child(10) > div > div > div > div > div.styles__ProductListGridFadedLoading-sc-u8zdb1-0 > div.styles__StyledRow-sc-wmoju4-0.cbPxiq > div > div.Pagination__RootDiv-sc-sq3l8r-4.bdkChJ > div:nth-child(3) > button")))
            load_more.click()
            time.sleep(1)
            scrape_page(driver, False)
        except:
            break

# Function that actually scrapes
def scrape_page(driver: webdriver.Chrome, startPage: bool):
    # Special conditions when it begins
    if startPage:
        driver.execute_script(f"window.scrollBy(0, 1000);")
    
    # Need to back up when it is any page but 1
    else:
        driver.execute_script("document.body.style.zoom='25%'")
        driver.execute_script(f"window.scrollBy(0, -500);")

    time.sleep(1)

    # Grab page html
    page_source = driver.page_source

    # Create Beautiful Soup Object
    soup = BeautifulSoup(page_source,"html.parser")

    # Grab all cards
    product_cards = soup.find_all('div', {'class': 'hCeGXD'})

    # Grab the product and deal from within each card
    for card in product_cards:
        try:
            product = card.find(attrs={'data-test': 'product-title'}).text
            if card.find('div', class_='styles__Truncate-sc-1wcknu2-0 hcXfd'):
                deal = card.find('div', class_='styles__Truncate-sc-1wcknu2-0 hcXfd').text
            else:
                deal = card.find('div', class_='h-text-red').text 

            print("Product: ", product)
            print("Deal: ", deal)
            print()
        except Exception as e:
            print("Error:", str(e))
            continue

    # Scroll down and select next page, must zoom back in in order to see button
    driver.execute_script(f"window.scrollBy(0, 800);")
    driver.execute_script("document.body.style.zoom='100%'")

    
def main():
    scrape_target()

if __name__ == "__main__":
    main()