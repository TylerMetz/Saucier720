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

def scrape_target():
    # Get to test website 
    url = "https://www.target.com/c/grocery-deals/-/N-k4uyq"
        
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

    # Let initial page load
    time.sleep(10)
    
    # Accumulate all page HTML
    all_html = ""

    while True:
        # Wait for the "Next Page" button to be clickable
        next_button = driver.find_element(By.CSS_SELECTOR, "#pageBodyContainer > div > div:nth-child(1) > div > div:nth-child(13) > div > div.styles__ProductListGridFadedLoading-sc-u8zdb1-0 > div.styles__StyledRow-sc-wmoju4-0.ftXYPI > div > div.styles__RootDiv-sc-l17a0m-5.hgWYOr > div:nth-child(3) > button")
        WebDriverWait(driver, timeout=10).until(lambda d : next_button.is_displayed())

        # If the "Next Page" button is disabled, stop the loop
        if not next_button.is_enabled():
            break

        # get page html
        all_html += driver.page_source

        # click the next page button
        next_button.click()
        next_button = None
        time.sleep(2) # let next page load

    # Close the WebDriver
    driver.quit()

    # Create a Beautiful Soup object from all accumulated HTML
    soup = BeautifulSoup(all_html, 'html.parser')

    # Extract the data and remove white space
    products = [product.get_text(strip=True) for product in soup.find_all("a", class_="styles__StyledLink-sc-vpsldm-0 styles__StyledTitleLink-sc-14ktig2-1 cbOry csOImU h-display-block h-text-bold h-text-bs")]
    prices = [price.get_text(strip=True) for price in soup.find_all("div", class_="h-padding-r-tiny")]

    # Print the extracted data
    for product, price in zip(products, prices):
        print("Product:", product)
        print("Price:", price)
        print()

def main():
    scrape_target()

if __name__ == "__main__":
    main()