package services

import (
	"log"
	"mailcast-worker/configuration"

	"github.com/go-resty/resty/v2"
)

var clientResty = resty.New()

func SendWaMessage(payload map[string]interface{}) {
	log.Printf("Start Sending WhatsApp messages... ")

	// Send the POST request with headers
	resp, err := clientResty.R().
		SetHeaders(map[string]string{
			// "Authorization": "defaultDS-49434e96f251d2ff",
			// "x-api-key":     "23b964f4c543ccdc",
			// "jwt":           util.JWT,
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"token":        configuration.CONFIG.DaisiApiToken,
		}).
		SetBody(payload).
		Post(configuration.CONFIG.DaisiApiUrl)

	if err != nil {
		log.Fatalf("Error occurred while sending message: %v", err)
	}

	log.Println("--------- Start Message ---------")
	log.Println("Request payload: ", payload)
	log.Println("--------- End Message ---------")

	log.Println("Response Status:", resp.Status())
	log.Println("Response Body:", resp.String())

	log.Printf("End Sending WhatsApp messages... ")
}
