package BackendPkg

import (
	"fmt"
	"runtime"
	"time"
	"strconv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Scraper struct {
	Store                GroceryStore
	TimeLastDealsScraped time.Time
	PublixDeals			[]FoodItem
	WalmartDeals 		[]FoodItem
}

func (s *Scraper) Scrape() {
	// scrapes data for all stores

	// scraped publix deals
	//s.PublixScrapeDeals()
	fmt.Println("Publix Deals Scraped!")

	// scraped walmart deals
	s.WalmartScrapeDeals()
	fmt.Println("Walmart Deals Scraped!")

	// saves current time to ref later
	s.TimeLastDealsScraped = time.Now()
}

func (s *Scraper) PublixScrapeDeals() {
	// init chrome driver
	opts := []selenium.ServiceOption{}
	if runtime.GOOS == "windows" {
		service, err := selenium.NewChromeDriverService("SeleniumDrivers/chromedriver_win32/chromedriver.exe", 9515, opts...)
		if err != nil {
			panic(err)
		}
		defer service.Stop()
	} else {
		service, err := selenium.NewChromeDriverService("SeleniumDrivers/chromedriver_mac64/chromedriver", 9515, opts...)
		if err != nil {
			panic(err)
		}
		defer service.Stop()
	}

	// init headless browser
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

	// run headless chrome browser
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// open desired page
	err = wd.Get("https://www.publix.com/savings/weekly-ad/view-all")
	if err != nil {
		panic(err)
	}

	// select "Choose a Store button"
	chooseStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div.savings-content-wrapper.skeleton-spacer > div > div.savings-container.full-bleed > div > div > button > span")
	if err != nil {
		panic(err)
	}
	err = chooseStoreButton.Click()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println("landing page button selected")
	}

	alternateLayout := false
	// input desired zipcode
	inputBox, err := wd.FindElement(selenium.ByCSSSelector, "#main > div:nth-child(5) > div > div > div.content.no-padding > div.p-store-locator > div > div > div > form > input[type=search]")
	if err != nil {
		alternateLayout = true //checks if the alternate windowed version is running (runs on some networks with the window on the right side of the screen)
	}
	if alternateLayout == false {
		time.Sleep(10 * time.Second)
		err = inputBox.SendKeys(s.Store.ZipCode)
		if err != nil {
			panic(err)
		} else {
			//fmt.Println("zip inputed")
		}

		// search for stores
		searchStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div:nth-child(5) > div > div > div.content.no-padding > div.p-store-locator > div > div > div.search-container > form > button")
		if err != nil {
			panic(err)
		}
		err = searchStoreButton.Click()
		if err != nil {
			panic(err)
		} else {
			//fmt.Println("search button pressed")
		}
	} else {
		alternateLayoutThree := false
		// sets the input box of the alternate window as the input box
		inputBoxTwo, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div > form > input[type=search]")
		if err != nil {
			alternateLayoutThree = true //checks if the alternate windowed version is running (runs on some networks with the window on the right side of the screen)
		}
		if !alternateLayoutThree {
			time.Sleep(10 * time.Second)
			err = inputBoxTwo.SendKeys(s.Store.ZipCode)
			// search for stores
			searchStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div > form > button")
			if err != nil {
				panic(err)
			}
			err = searchStoreButton.Click()
			if err != nil {
				panic(err)
			} else {
				//fmt.Println("search button pressed")
			}
		} else {
			inputBoxThree, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div.search-container > form > input[type=search]")
			time.Sleep(10 * time.Second)
			err = inputBoxThree.SendKeys("32601")
			if err != nil {
				fmt.Println("not third layout")
			}
			// search for stores
			searchStoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#navBar > div > div.navigation-bar-main > div > div > div.navigation-section.top > div.user-navigation > div > div > div.navigation-sidebar-container > div.navigation-sidebar-body > div > div > div > div.search-container > form > button")
			if err != nil {
				panic(err)
			}
			err = searchStoreButton.Click()
			if err != nil {
				panic(err)
			} else {
				//fmt.Println("search button pressed")
			}
		}
	}
	time.Sleep(20 * time.Second) // wait for page to load

	// select the first store from the results list
	chooseStoreResult, err := wd.FindElement(selenium.ByCSSSelector, "#\\31 560 > div > div > div.buttons-links > div.p-button-group__wrapper.buttons-wrapper > div > button")
	if err != nil {
		panic(err)
	}
	err = chooseStoreResult.Click()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println("store selected")
	}
	time.Sleep(5 * time.Second) // wait for page to load

	// switch to list view
	listViewButton, err := wd.FindElement(selenium.ByCSSSelector, "#listing")
	if err != nil {
		panic(err)
	}
	err = listViewButton.Click()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println("list view enabled")
	}

	time.Sleep(20 * time.Second) // wait for page to load

	// Keeps hitting "Load More" button until all of the data is loaded
	moreLoadingNeeded := true
	// triggers page to load more
	for moreLoadingNeeded {
		loadMoreButton, err := wd.FindElement(selenium.ByCSSSelector, "#main > div.savings-content-wrapper > div > div.savings-container > div.card-loader.savings-content.search-results-section.-coupons > div.button-container > button")
		if err != nil {
			moreLoadingNeeded = false
		} else {
			_ = loadMoreButton.Click()
			time.Sleep(3 * time.Second)
		}
	}

	// stores expanded deals page html
	html, err := wd.PageSource()
	if err != nil {
		panic(err)
	}

	// organize and store publix deals as a slice of FoodItems
	s.PublixDeals = s.Store.OrganizeDeals(html)

}

func (s *Scraper) WalmartScrapeDeals() {
	// init chrome driver
	opts := []selenium.ServiceOption{}
	if runtime.GOOS == "windows" {
		service, err := selenium.NewChromeDriverService("SeleniumDrivers/chromedriver_win32/chromedriver.exe", 9515, opts...)
		if err != nil {
			panic(err)
		}
		defer service.Stop()
	} else {
		service, err := selenium.NewChromeDriverService("SeleniumDrivers/chromedriver_mac64/chromedriver", 9515, opts...)
		if err != nil {
			panic(err)
		}
		defer service.Stop()
	}

	// init headless browser
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
	
	// run headless chrome browser
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Open the Walmart webpage
	if err := wd.Get("https://www.walmart.com/browse/grocery-deals/c2hlbGZfaWQ6MjQ1NTI0NQieie"); err != nil {
		panic(err)
	}

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// Find the wrapper containing all the items
	wrapper, err := wd.FindElement(selenium.ByCSSSelector, "#maincontent > main > div > div > div > div > div.w-100.relative-m.pl4.pr4.flex.pt2 > div.relative.w-80 > div > section > div")
	if err != nil {
		panic(err)
	}

	// Find all the items within the wrapper
	items, err := wrapper.FindElements(selenium.ByCSSSelector, ".search-result-gridview-item-wrapper")
	if err != nil {
		panic(err)
	}

	// Iterate over the items and extract the information
	var foodItems []FoodItem
	for _, item := range items {
		nameElem, _ := item.FindElement(selenium.ByCSSSelector, "span span")
		name, _ := nameElem.Text()

		priceElem, _ := item.FindElement(selenium.ByCSSSelector, "div.flex.flex-wrap.justify-start.items-center.lh-title.mb1.mb0 > div.b.black.lh-copy.f5.f4-l")
		price, _ := priceElem.Text()
		salePrice, _ := strconv.ParseFloat(price, 64)

		detailsElem, _ := item.FindElement(selenium.ByCSSSelector, "div.gray.mr1.f7.f6-l")
		details, _ := detailsElem.Text()

		foodItem := FoodItem{
			Name:        name,
			SalePrice:   salePrice,
			SaleDetails: details,
		}
		foodItems = append(foodItems, foodItem)
	}

	// store deals in scraper class
	s.WalmartDeals = foodItems
}