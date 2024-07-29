package clients

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"strconv"
	"github.com/waldo121/events-nearby-app/models"
)

const apiBaseUrl = "https://donnees.montreal.ca/api/"
const resourceId = "6decf611-6f11-4f34-bb36-324d804c9bad" // events in csv format

type ApiResponse struct {
	Result Result `json:"result"`
}

type Result struct  {
 Events []models.Event `json:"records"`
}

// Get <limit> latest events through api v3
func GetEvents(limit int) []models.Event {
	var response ApiResponse
	limitString := strconv.Itoa(limit)
	resp, err := http.Get(apiBaseUrl+"3/action/datastore_search?resource_id="+resourceId+"&limit="+limitString)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	return response.Result.Events
}
