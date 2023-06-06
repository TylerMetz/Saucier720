import requests
from bs4 import BeautifulSoup

def scrape_walmart():
    url = "https://www.walmart.com/shop/deals/food/foodrollbacks"
    headers = {
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"}
    
    response = requests.get(url, headers=headers)
    soup = BeautifulSoup(response.content, "html.parser")
    
    # Extract the desired data from the soup object
    # Modify the code below according to your specific requirements
    
    # Example: Extract all product names and prices from the home page
    products = soup.find_all("span", {"class": "normal dark-gray mb0 mt1 lh-title f6 f5-l lh-copy"})
    prices = soup.find_all("span", {"class": "w_iUH7"})
    
    if len(prices) == 0 & len(products) == 0:
        print("No products found")
    
    for product, price in zip(products, prices):
        print("Product:", product.text.strip())
        print("Price:", price.text.strip())
        print()
        
def main():
    scrape_walmart()

if __name__ == "__main__":
    main()
