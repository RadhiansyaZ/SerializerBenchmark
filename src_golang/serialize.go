package main

import (
	// normal impl:"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	// jsoniter impl:
	jsoniter "github.com/json-iterator/go"
)

func Serialize() []byte {
	start := time.Now()
	// Variable declaration to handle the API response
	var objmap map[string]interface{}
	// jsoniter impl:
	var jsonit = jsoniter.ConfigCompatibleWithStandardLibrary

	// Fetching the API response
	response, err := http.Get("http://localhost:8080/sample.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Reading the API response
	responseData, err := ioutil.ReadAll(response.Body)
	// jsoniter impl:
	// normal impl: if err = json.Unmarshal(responseData, &objmap); err != nil {
	if err = jsonit.Unmarshal(responseData, &objmap); err != nil {

		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Modifying the API response
	for i := 0; i < len(objmap["discussion_results"].([]interface{})); i++ {
		objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["orders"] = i + 1
		objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["reply_count"] = len(objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["replies"].([]interface{}))
	}

	// Serializing the modified API response
	// normal impl: finalres, err := json.Marshal(objmap)
	// jsoniter impl:
	finalres, err := jsonit.Marshal(objmap)
	if err != nil {
		panic(err)
	}

	// Code to measure
	duration := time.Since(start)

	os.WriteFile("out.txt", []byte("duration: "+duration.String()), 0666)

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
