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

def scrape_costco():
    # Get to test website 
    url = "https://www.costco.com/grocery-household.html?keyword=OFF&dept=All&sortBy=item_page_views+desc"
    
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
    
    # Accumulate all page HTML
    all_html = ""

    while True:
        # Wait for the "Next Page" button to be clickable
        
        try:
            next_button = WebDriverWait(driver, 5).until(EC.element_to_be_clickable((By.CSS_SELECTOR, '#search-results > div.product-list.grid > nav > div > div.paging.col-xs-12 > ul > li.forward > a')))
            
            # get page html
            all_html += driver.page_source

            # Find the "Next Page" button and click it
            next_button.click()

            # Re-locate the next_button element after the page navigation
            next_button = WebDriverWait(driver, 5).until(EC.presence_of_element_located((By.CSS_SELECTOR, '#search-results > div.product-list.grid > nav > div > div.paging.col-xs-12 > ul > li.forward > a')))
            
        except Exception as e:
            # Break the loop if the "Next Page" button is not clickable (end of pagination)
            break
        
        


    # Close the WebDriver
    driver.quit()

    # Create a Beautiful Soup object from all accumulated HTML
    soup = BeautifulSoup(all_html, 'html.parser')

    # Extract the data and remove white space
    products = [a.get_text(strip=True) for span in soup.find_all("span", class_="description") for a in span.find_all('a')]
    prices = [price.get_text(strip=True) for price in soup.find_all("div", class_="price")]

    # Print the extracted data
    for product, price in zip(products, prices):
        print("Product:", product)
        print("Price:", price)
        print()

def main():
    scrape_costco()

if __name__ == "__main__":
    main()