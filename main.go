package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService("src/SeleniumDrivers/chromedriver_mac64/chromedriver", 9515, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	err = wd.Get("https://www.publix.com/savings/weekly-ad/view-all")
	if err != nil {
		panic(err)
	}

	chooseStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div.savings-content-wrapper.skeleton-spacer > div > div.savings-container.full-bleed > div > div > button > span")
	if err != nil {
		panic(err)
	}
	err = chooseStoreButton.Click()
	if err != nil {
		panic(err)
	} else{
		fmt.Println("landing page button selected")
	}

	inputBox, err := wd.FindElement(selenium.ByCSSSelector, "#main > div:nth-child(5) > div > div > div.content.no-padding > div.p-store-locator > div > div > div > form > input[type=search]")
	if err != nil {
		panic(err)
	}
	err = inputBox.SendKeys("32601")
	if err != nil {
		panic(err)
	} else{
		fmt.Println("zip inputed")
	}

	searchStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div:nth-child(5) > div > div > div.content.no-padding > div.p-store-locator > div > div > div.search-container > form > button")
	if err != nil {
		panic(err)
	}
	err = searchStoreButton.Click()
	if err != nil {
		panic(err)
	} else{
		fmt.Println("search button pressed")
	}
	time.Sleep(20 * time.Second) // wait for page to load

	chooseStoreResult, err := wd.FindElement(selenium.ByCSSSelector, "#\\31 560 > div > div > div.buttons-links > div.p-button-group__wrapper.buttons-wrapper > div > button")
	if err != nil {
		panic(err)
	}
	err = chooseStoreResult.Click()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("store selected")
	}
	time.Sleep(5 * time.Second) // wait for page to load

	listViewButton, err := wd.FindElement(selenium.ByCSSSelector, "#listing")
	if err != nil {
		panic(err)
	}
	err = listViewButton.Click()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("list view enabled")
	}

	time.Sleep(20 * time.Second) // wait for page to load

	// Keeps hitting "Load More" button until all of the data is loaded
	moreLoadingNeeded := true;
	// triggers page to load more
	for moreLoadingNeeded{
		loadMoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div.savings-content-wrapper > div > div.savings-container > div.card-loader.savings-content.search-results-section.-coupons > div.button-container > button")
		if err != nil{
			moreLoadingNeeded = false
		} else{
			_ = loadMoreButton.Click()
			time.Sleep(3 * time.Second)
		}
	}

	html, err := wd.PageSource()
	if err != nil {
		panic(err)
	}

	fmt.Println(html)
}
