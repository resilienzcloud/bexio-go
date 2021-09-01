package bexio

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client is the gateway to the API.
type Client struct {
	restyClient *resty.Client
}

// NewClient creates a new bexio client.
func NewClient(token string) (client *Client) {
	// Create the resty client.
	resty := resty.New()

	// Configure the resty client.
	resty.SetAuthToken(token)
	resty.SetHostURL("https://api.bexio.com")
	resty.SetHeader("Accept", "application/json")

	// Create the client with configured resty client.
	client = &Client{restyClient: resty}

	// Return the new client.
	return client
}

// CreateContact creates a given contact.
func (c *Client) CreateContact(contact Contact) (Contact, error) {
	resp, err := c.restyClient.R().
		SetBody(contact).
		SetResult(&contact).
		Post("/2.0/contact")

	// Check for direct errors.
	if err != nil {
		return contact, err
	}

	// Unless we receive a 201 (Created) we return an error.
	if resp.StatusCode() != 201 {
		return contact, fmt.Errorf(resp.String())
	}

	return contact, nil
}
