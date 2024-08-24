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

type Result struct {
	Discussions []Comment `json:"discussion_results"`
}

type Comment struct {
	Orders     int     `json:"orders"`
	Username   string  `json:"username"`
	Comment    string  `json:"comment"`
	ReplyCount int     `json:"reply_count"`
	Replies    []Reply `json:"replies"`
}

func (c *Comment) CountReplies() {
	c.ReplyCount = len(c.Replies)
}

type Reply struct {
	Username string `json:"username"`
	Comment  string `json:"comment"`
}

func SerializeToStruct() []byte {
	start := time.Now()

	// jsoniter impl:
	// can be interchangeable by comment line below and import "encoding/json"
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	// Fetching the API response
	response, err := http.Get("http://localhost:8080/sample.json")
	if err != nil {
		log.Fatal(err)
	}

	// Reading the API response
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Variable declaration to handle the API response
	var objmap Result
	if err = json.Unmarshal(responseData, &objmap); err != nil {
		log.Fatal(err)
	}

	// Modifying the API response
	for i := range objmap.Discussions {
		objmap.Discussions[i].Orders = i + 1
		objmap.Discussions[i].CountReplies()
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
	// Serialize()
	SerializeToStruct()
}
