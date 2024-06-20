package clients

import (
	"net/http"
	"errors"
	"log"
)

const apiBaseUrl := "https://donnees.montreal.ca/api/"

// Get <limit> latest events through api v3
func getEvents(limit number) (models.Event[], error) {
	
	resp, err := http.Get(apiBaseUrl+"3/action/datastore_search&limit="+limit)
	if err != nil {
		log.err(err)
		return nil, errors.New("Error fetching events data")
	}
	fmt.Println(resp.data)
	return [], nil
}
