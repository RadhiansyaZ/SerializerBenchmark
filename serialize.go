package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Serialize() []byte {
	// Variable declaration to handle the API response
	var objmap map[string]interface{}

	// Fetching the API response
	response, err := http.Get("http://localhost:8080/sample.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Reading the API response
	responseData, err := ioutil.ReadAll(response.Body)
	if err = json.Unmarshal(responseData, &objmap); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Modifying the API response
	user := objmap["discussion_results"].([]interface{})[0].(map[string]interface{})["username"].(string)
	user = user + "_test_concat"
	objmap["discussion_results"].([]interface{})[0].(map[string]interface{})["username"] = user
	for i := 0; i < len(objmap["discussion_results"].([]interface{})); i++ {
		objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["orders"] = i + 1
		objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["reply_count"] = len(objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["replies"].([]interface{}))
	}

	// Serializing the modified API response
	finalres, err := json.Marshal(objmap)
	if err != nil {
		panic(err)
	}

	return finalres

}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(Serialize())
}

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
	Serialize()
}
