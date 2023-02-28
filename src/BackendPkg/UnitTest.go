package BackendPkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func TestingPortOutput(testItems []interface{}) string {

	// convert items to be translated into json
	itemsInJson, _ := json.Marshal(testItems)

	// create a router to output items to the port
	testRouter := Router{
		Name:             "test1",
		ItemsToBeEncoded: testItems,
	}

	// display item on port in background
	go testRouter.Rout()

	//pull data from local host port
	resp, err := http.Get("http://localhost:8080/api/Pantry")
	if err != nil {
		fmt.Println("Error!")
	}
	defer resp.Body.Close()

	// make website data a string
	body, err := ioutil.ReadAll(resp.Body)
	portValue := string(body)

	// compare converted data
	if portValue == string(itemsInJson) {
		return "Test Passed"
	} else {
		return "Test not passed"
	}

}
