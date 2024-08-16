package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	// "github.com/mitchellh/mapstructure"
)

type SortedDiscussion struct {
	DiscussionResults []struct {
		Orders     int    `json:"orders"`
		Username   string `json:"username"`
		Comment    string `json:"comment"`
		ReplyCount int    `json:"reply_count"`
		Replies    []struct {
			Username string `json:"username"`
			Comment  string `json:"comment"`
		} `json:"replies"`
	} `json:"discussion_results"`
}

func Serialize() []byte {
    response, err := http.Get("http://localhost:8080/sample.json")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    // fmt.Println(reflect.TypeOf(responseData))
	var objmap map[string]interface{}
	// var objmapres SortedDiscussion
	if err = json.Unmarshal(responseData, &objmap); err != nil {
		log.Fatal(err)
	}
    if err != nil {
        log.Fatal(err)
    }
	user := objmap["discussion_results"].([]interface{})[0].(map[string]interface{})["username"].(string)
	user = user + "_test_concat"
	objmap["discussion_results"].([]interface{})[0].(map[string]interface{})["username"] = user
	for i := 0; i < len(objmap["discussion_results"].([]interface{})); i++ {
		objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["orders"] = i + 1
		objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["reply_count"] = len(objmap["discussion_results"].([]interface{})[i].(map[string]interface{})["replies"].([]interface{}))
	}
	// mapstructure.Decode(objmap, &objmapres)
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