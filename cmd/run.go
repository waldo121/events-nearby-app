package cmd

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/waldo121/events-nearby-app/clients"
	"github.com/waldo121/events-nearby-app/models"
	"github.com/waldo121/events-nearby-app/util"
)

func Execute() {
	var eventsNumber int
	var cost string
	var eventType string
	flag.IntVar(&eventsNumber, "nbEvents", 10, "number of events to retrieve from the api")
	flag.StringVar(&cost, "cost", "", "Event cost")
	flag.StringVar(&eventType, "type", "", "Event type")
	flag.Parse()
	var result []models.Event = clients.GetEvents(eventsNumber)

	if len(cost) != 0 {
		result = util.Filter(result, func(Element models.Event) bool {
			return Element.Cost == cost
		})
	}

	if len(eventType) != 0 {
		result = util.Filter(result, func(Element models.Event) bool {
			return Element.Type == eventType
		})
	}
	s, err := json.MarshalIndent(result, "", " ")
	if err == nil {
		fmt.Println(string(s))
	}

}
