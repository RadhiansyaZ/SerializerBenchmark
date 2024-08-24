package main

import (
	// normal impl:
	// "encoding/json"

	"io"
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
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	// Fetching the API response
	response, err := http.Get("http://localhost:8080/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Reading the API response
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(responseData, &objmap); err != nil {
		log.Fatal(err)
	}

	// Modifying the API response
	for i := 0; i < len(objmap["discussion_results"].([]interface{})); i++ {
		objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["orders"] = i + 1
		objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["reply_count"] = len(objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["replies"].([]interface{}))
	}

	// Serializing the modified API response
	finalres, err := json.Marshal(objmap)
	if err != nil {
		log.Fatal(err)
	}

	// Code to measure
	duration := time.Since(start)

	os.WriteFile("out.txt", []byte("duration: "+duration.String()), 0666)

	return finalres

}

func main() {
	Serialize()
}
