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

def scrape_costco():
    # Get to test website 
    url = "https://www.costco.com/grocery-household.html?keyword=OFF&dept=All&sortBy=item_page_views+desc"
    
    # Set up Selenium options 
    options = Options()
    options.add_argument("--headless")
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
    
    #while True:
        # Get the page source
    page_source = driver.page_source
    
    # Create BeautifulSoup object to parse the page source
    soup = BeautifulSoup(page_source, "html.parser")
    
    # Extract the data and remove white space
    products = [a.get_text(strip=True) for span in soup.find_all("span", class_="description") for a in span.find_all('a')]  
    prices = [price.get_text(strip=True) for price in soup.find_all("div", class_="price")]
    
    print(len(products))
    print(len(prices))
    
    for product, price in zip(products, prices):
        print("Product:", product)
        print("Price:", price)
        print()


def main():
    scrape_costco()

if __name__ == "__main__":
    main()