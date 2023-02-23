package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		seleniumPath    = "src/SeleniumDrivers/selenium-server-4.8.1.jar"
		geckoDriverPath = "src/SeleniumDrivers/chromedriver_mac64/chromedriver"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Go to the Publix website
	if err := wd.Get("https://www.publix.com/savings/weekly-ad/view-all"); err != nil {
		panic(err)
	}

	// Click the "Choose a store" button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, ".p-button button--default button--lg button--primary")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

    // KEY ENTRY

    // tab down to search bar
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.TabKey)

    // enter zip code, gnv by default
    selenium.KeyDownAction("3")
    selenium.KeyDownAction("2")
    selenium.KeyDownAction("6")
    selenium.KeyDownAction("0")
    selenium.KeyDownAction("1")

    // go to search button and select
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.EnterKey)

    // tab down to "Choose Store" button and hit enter
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.TabKey)
    selenium.KeyDownAction(selenium.EnterKey)

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "p-grid-item")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))


}