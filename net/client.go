package net

import "log"

type Client struct{}

//
func (client *Client) get(url string) (string, error) {
	log.Print(client)
	return "", nil
}
