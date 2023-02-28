package BackendPkg

import (
	"fmt"
	"io/ioutil"
	"new/http"

	"github.com/gorilla/mux"
)

func TestingJSON(testItems []interface{}, expectedOutput string) string {
	testRouter := BackendPkg.Router{
		Name:             "test1",
		ItemsToBeEncoded: testItems,
	}
	testRouter.Rout()
	//pull data from local host port
	resp, err := http.Get("http://localhost:8080/api/Pantry")
	if err != nil {
		fmt.Fprint("Error: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if body == expectedOutput {
		return "Test Passed"
	} else {
		return "Fuck you"
	}

}
