

import os
import asyncio
import json
import math
from pathlib import Path
from typing import Dict, List, Tuple
from urllib.parse import urlencode, urljoin

from loguru import logger as log
from scrapfly import ScrapeApiResponse, ScrapeConfig, ScrapflyClient

scrapfly = ScrapflyClient(key=os.environ["scp-live-023707dacca84039a5045a8b3af6aa8f"], max_concurrency=5)


def parse_search(result: ScrapeApiResponse) -> Tuple[List[Dict], int]:
    """parse Walmart search results page for product previews"""
    log.debug(f"parsing search page {result.context['url']}")
    data = result.selector.xpath('//script[@id="__NEXT_DATA__"]/text()').get()
    data = json.loads(data)

    total_results = data["props"]["pageProps"]["initialData"]["searchResult"]["itemStacks"][0]["count"]
    results = data["props"]["pageProps"]["initialData"]["searchResult"]["itemStacks"][0]["items"]
    # there are other results types such as ads or placeholders - filter them out:
    results = [result for result in results if result["__typename"] == "Product"]
    log.info(f"parsed {len(results)} search product previews")
    return results, total_results


async def scrape_search(search: str, max_pages: int = 25) -> List[Dict]:
    """scrape walmart search for product previews"""

    def create_search_url(query: str, page=1, sort="price_low") -> str:
        """create url for a single walmart search page"""
        return "https://www.walmart.com/search?" + urlencode(
            {
                "q": query,
                "sort": sort,
                "page": page,
                "affinityOverride": "default",
            }
        )

    log.info(f"searching walmart for {search}")
    first_page = await scrapfly.async_scrape(ScrapeConfig(create_search_url(query=search), country="US", asp=True))
    previews, total_items = parse_search(first_page)

    total_pages = math.ceil(total_items / 40)
    log.info(f"found total {total_pages} pages of results ({total_items} products)")
    if max_pages and total_pages > max_pages:
        total_pages = max_pages

    other_pages = [
        ScrapeConfig(url=create_search_url(query=search, page=i), asp=True, country="US")
        for i in range(2, total_pages + 1)
    ]
    async for result in scrapfly.concurrent_scrape(other_pages):
        previews.extend(parse_search(result)[0])
    log.info(f"parsed total {len(previews)} pages of results ({total_items} products)")
    return previews


def parse_product(result: ScrapeApiResponse):
    """parse walmart product from product page response"""
    data = result.selector.xpath('//script[@id="__NEXT_DATA__"]/text()').get()
    if not data:
        log.error(f"{result.context['url']} has no product data")
    data = json.loads(data)
    _product_raw = data["props"]["pageProps"]["initialData"]["data"]["product"]
    wanted_product_keys = [
        "availabilityStatus",
        "averageRating",
        "brand",
        "id",
        "imageInfo",
        "manufacturerName",
        "name",
        "orderLimit",
        "orderMinLimit",
        "priceInfo",
        "shortDescription",
        "type",
    ]
    product = {k: v for k, v in _product_raw.items() if k in wanted_product_keys}
    reviews_raw = data["props"]["pageProps"]["initialData"]["data"]["reviews"]
    return {"product": product, "reviews": reviews_raw}


async def scrape_products(urls: List[str]):
    """scrape walmart products by urls"""
    log.info(f"scraping {len(urls)} product urls (in chunks of 50)")
    results = []
    to_scrape = [ScrapeConfig(url=url, asp=True, country="US") for url in urls]
    async for result in scrapfly.concurrent_scrape(to_scrape):
        results.append(parse_product(result))
    return results


async def scrape_search_and_products(search: str, max_pages: int = 25):
    """scrape walmart search to find products and then scrape complete product details"""
    search_results = await scrape_search(search, max_pages=max_pages)
    product_urls = [
        urljoin("https://www.walmart.com/", product_preview["canonicalUrl"]) for product_preview in search_results
    ]
    return await scrape_products(product_urls)


if __name__ == "__main__":
    results = scrape_products("https://www.walmart.com/browse/grocery-deals/c2hlbGZfaWQ6MjQ1NTI0NQieie")
    print(results)