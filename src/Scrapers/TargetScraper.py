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
    options.add_argument("--headless")
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
    driver.execute_script("document.body.style.zoom='25%'")
    
    # Give time to load
    time.sleep(20)
    scrapePage(driver)
    
    # Get every other set of deals
    while True:
        try:                                                                                        
            load_more = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR,"#pageBodyContainer > div > div:nth-child(1) > div > div:nth-child(10) > div > div > div > div > div.styles__ProductListGridFadedLoading-sc-u8zdb1-0 > div.styles__StyledRow-sc-wmoju4-0.cbPxiq > div > div.Pagination__RootDiv-sc-sq3l8r-4.bdkChJ > div:nth-child(3) > button")))
            load_more.click()
            time.sleep(10)
            page_source = page_source + driver.page_source
        except:
            break

def scrapePage(driver: webdriver.Chrome):
    driver.execute_script(f"window.scrollBy(0, 1000);")
    time.sleep(5)
    page_source = driver.page_source

    soup = BeautifulSoup(page_source,"html.parser")
    #products = soup.find_all(attrs={'data-test': 'product-title'})
    #deals = soup.find_all('div', {'class': 'h-display-flex'})

    product_cards = soup.find_all('div', {'class': 'hCeGXD'})

    for card in product_cards:
        
        product = card.find('a', class_='styles__StyledTitleLink-sc-14ktig2-1').text
        if card.find('div', class_='styles__Truncate-sc-1wcknu2-0 hcXfd'):
            deal = card.find('div', class_='styles__Truncate-sc-1wcknu2-0 hcXfd').text
        else:
            deal = card.find('div', class_='h-text-red').text 

        print("Product: ", product)
        print("Deal: ", deal)
        print()

    

    time.sleep(5)

    
def main():
    scrape_target()

if __name__ == "__main__":
    main()