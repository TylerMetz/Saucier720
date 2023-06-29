import re
from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.keys import Keys
import sys

def scrape_walmart():
    url = "https://walmart.com/shop/deals/food/foodrollbacks"
    
    if sys.platform.startswith('win'):
        driverPath = "SeleniumDrivers/chromedriver_win32/chromedriver.exe"
    else:
        driverPath = "SeleniumDrivers/chromedriver_mac64/chromedriver"
    
    while True:
        # Set up Selenium options
        options = Options()
        options.add_argument("--headless")  # Run Chrome in headless mode
        options.add_argument("--disable-gpu")
        options.add_argument("--no-sandbox")
        options.add_argument("--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

        # Start Selenium webdriver
        driver = webdriver.Chrome(executable_path = driverPath, options=options)
        
        # Open the webpage
        driver.get(url)
        
        # Wait for the page to load completely (increase the timeout if needed)
        wait = WebDriverWait(driver, 10)
        # Get the page source
        page_source = driver.page_source
        
        #print(page_source)
        
        # Create BeautifulSoup object to parse the page source
        soup = BeautifulSoup(page_source, "html.parser")
        
        # Extract the desired data from the soup object
        # Modify the code below according to your specific requirements
        
         # Example: Extract all product names and prices from the home page
        products = soup.find_all("span", {"class": "normal dark-gray mb0 mt1 lh-title f6 f5-l lh-copy"})
        prices = soup.find_all("div", {"class": "flex flex-wrap justify-start items-center lh-title mb1"})
            
        for product, price in zip(products, prices):
            # product name
            try:
                newStr = re.search(r'^([^0-9]+)', product.text.strip()).group(1).strip().rstrip(',')
                if newStr:
                    print("Product:", newStr)
            except:
                print("Product:", product.text.strip())
            
            # product price
            print("Price: " + "$" + re.findall(r'\$([\d.]+)', price.text.strip())[0])
            print()
        
        # Check if the next page button is present
        next_link = soup.find("a", {"aria-label": "Next Page"})
        if next_link:
            next_url = next_link["href"]
            url = "https://walmart.com" + next_url
            # Close the Selenium webdriver
            driver.quit()
        else:
            break
        
def main():
    scrape_walmart()

if __name__ == "__main__":
    main()
