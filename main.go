package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Config struct {
	APIKey    string `json:"apiKey"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func SendEmail(config *Config, subject, content string) error {
	from := mail.NewEmail("Sender Name", config.Sender)
	to := mail.NewEmail("Recipient Name", config.Recipient)
	message := mail.NewSingleEmail(from, subject, to, "", content)

	client := sendgrid.NewSendClient(config.APIKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("failed to send email: %s", response.Body)
	}

	return nil
}

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err.Error())
	}

	http.HandleFunc("/webhook", handleWebhookEvent)
	http.HandleFunc("/send-email", func(w http.ResponseWriter, r *http.Request) {
		sendEmail(w, r, config)
	})

	log.Fatal(http.ListenAndServe(":8281", nil))
}

func sendEmail(w http.ResponseWriter, r *http.Request, config *Config) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	subject := r.FormValue("subject")
	content := r.FormValue("content")

	err := SendEmail(config, subject, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Email sent successfully!")
}

type WebhookPayload struct {
	Email     string `json:"email"`
	Event     string `json:"event"`
	SGEventID string `json:"sg_event_id"`
}

func handleWebhookEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payloads []WebhookPayload
	err := json.NewDecoder(r.Body).Decode(&payloads)
	if err != nil {
		log.Printf("Failed to parse webhook payload: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, payload := range payloads {
		// Handle different webhook events
		switch payload.Event {
		case "delivered":
			log.Printf("Email delivered: SG Event ID - %s, Recipient - %s", payload.SGEventID, payload.Email)

		case "open":
			log.Printf("Email opened: SG Event ID - %s, Recipient - %s", payload.SGEventID, payload.Email)

		case "click":
			log.Printf("Link clicked: SG Event ID - %s, Recipient - %s", payload.SGEventID, payload.Email)

		case "bounce":
			log.Printf("Email bounced: SG Event ID - %s, Recipient - %s", payload.SGEventID, payload.Email)

		default:
			log.Printf("Unknown webhook event received: SG Event ID - %s, Recipient - %s, Event - %s", payload.SGEventID, payload.Email, payload.Event)
		}
	}

	w.WriteHeader(http.StatusOK)
}
