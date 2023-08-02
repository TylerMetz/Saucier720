import re
import time
from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager
from selenium_stealth import stealth

def scrape_walmart():
    url = "https://walmart.com/shop/deals/food/foodrollbacks"
    
    # Set up Selenium options
    options = Options()
    options.add_argument("--headless")  # Run Chrome in headless mode
    options.add_argument("--disable-gpu")
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-geolocation")
    options.add_experimental_option("excludeSwitches", ["enable-automation"])
    options.add_experimental_option('useAutomationExtension', False)
    options.add_experimental_option("prefs", {"profile.default_content_setting_values.geolocation": 2})

    # Start Selenium webdriver
    driver=webdriver.Chrome(service=Service(ChromeDriverManager(version='114.0.5735.90').install()),options=options)

    # Set stealth to be undetectable 
    stealth(driver,
       user_agent = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36',
       languages=["en-US", "en"],
       vendor="Google Inc.",
       platform="Win32",
       webgl_vendor="Intel Inc.",
       renderer="Intel Iris OpenGL Engine",
       fix_hairline=True,
       )

    # Open the webpage
    driver.get(url)

    
    time.sleep(10)
    # Wait for the page to load completely (increase the timeout if needed)
    wait = WebDriverWait(driver, 10)
    page_count = 1
    div_count = 1
    while True: 
        time.sleep(5)
        # Get the page source
        page_source = driver.page_source
        
        # Create BeautifulSoup object to parse the page source
        soup = BeautifulSoup(page_source, "html.parser")

        # Get the total amount of pages 
        if page_count == 1:
            total_pages_list = [div.text for div in soup.find_all("div", class_="sans-serif ph1 pv2 w4 h4 lh-copy border-box br-100 b--solid mh2-m db tc no-underline gray bg-white b--white-90")]
            total_pages = int(total_pages_list[0])
        
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
            

        # Page instances when the search bar is different sizes
        if page_count < 3:
            next_button = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, "#maincontent > main > div > div:nth-child(3) > div > div > div:nth-child(2) > nav > ul > li:nth-child(7) > a")))
            next_button.click()
            page_count += 1
        elif page_count == total_pages:
            driver.close()
        elif page_count < 4 or page_count == total_pages - 2 or page_count == total_pages - 1:
            next_button = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, "#maincontent > main > div > div:nth-child(3) > div > div > div:nth-child(2) > nav > ul > li:nth-child(8) > a")))
            next_button.click()
            page_count += 1
        elif page_count >= 4: 
            next_button = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, "#maincontent > main > div > div:nth-child(3) > div > div > div:nth-child(2) > nav > ul > li:nth-child(9) > a")))
            next_button.click()
            page_count += 1

        
def main():
    scrape_walmart()

if __name__ == "__main__":
    main()
