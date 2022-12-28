package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	secretKey string
	baseUrl   *url.URL
	client    *http.Client
}

type User struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Email      string `json:"email"`
}

type UserUpdate struct {
	Reference  string `json:"reference"`
	Identifier string `json:"identifier"`
}

type UserDelete struct {
	Reference string `json:"reference"`
}

type UserGet struct {
	Reference string `json:"reference"`
}

type LinkGet struct {
	LinkId string `json:"linkid"`
}

type LinkCharge struct {
	Amount uint64 `json:"amount"`
	Remark string `json:"remark"`
}

type ChargeAuthorization struct {
	Event string `json:"event"`
}

func NewClient(secretKey string) *Client {
	baseURL, _ := url.Parse("https://api.thepeer.co")

	return &Client{
		secretKey: secretKey,
		baseUrl:   baseURL,
		client: &http.Client{
			Transport: &http.Transport{},
		},
	}
}

func (c *Client) ValidateSignature(payload []byte, signature string) bool {

	h := hmac.New(sha1.New, []byte(c.secretKey))
	h.Write(payload)

	expectedSignatureString := hex.EncodeToString(h.Sum(nil))

	return signature == expectedSignatureString
}

func (c *Client) IndexUser(name string, identifier string, email string) []byte {
	user := User{
		Name:       name,
		Identifier: identifier,
		Email:      email,
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	req, err := c.MakeHTTPRequest(c.baseUrl, "POST", "/users", data)
	if err != nil {
		fmt.Println("could not make the request to the end point")
	}

	return req
}

func (c *Client) UpdateUser(reference string, identifier string) ([]byte, error) {
	user := UserUpdate{
		Reference:  reference,
		Identifier: identifier,
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	req, err := c.MakeHTTPRequest(c.baseUrl, "PUT", "/users/"+reference, data)

	if err != nil {
		fmt.Println("could not make the request to the end point")
	}

	return req, nil
}

func (c *Client) DeleteUser(reference string) ([]byte, error) {
	user := UserDelete{
		Reference: reference,
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	req, err := c.MakeHTTPRequest(c.baseUrl, "DELETE", "/users/"+reference, data)
	if err != nil {
		fmt.Println("could not get a response")
	}

	return req, nil
}

func (c *Client) GetUser(reference string) ([]byte, error) {
	user := UserGet{
		Reference: reference,
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	req, err := c.MakeHTTPRequest(c.baseUrl, "GET", "/users/"+reference, data)
	if err != nil {
		fmt.Println("could not make the request to the /users/reference end point")
	}

	return req, nil
}

func (c *Client) GetLink(linkId string) ([]byte, error) {
	link := LinkGet{
		LinkId: linkId,
	}

	data, err := json.Marshal(link)
	if err != nil {
		panic(err)
	}

	req, err := c.MakeHTTPRequest(c.baseUrl, "GET", "/link/"+linkId, data)
	if err != nil {
		fmt.Println("could not make the request to the /users/reference end point")
	}

	return req, nil
}

func (c *Client) ChargeLink(linkId string, amount uint64, remark string) []byte {
	link := LinkCharge{
		Amount: amount,
		Remark: remark,
	}

	data, err := json.Marshal(link)
	if err != nil {
		panic(err)
	}

	req, err := c.MakeHTTPRequest(c.baseUrl, "POST", "/link/"+linkId+"/charge", data)
	if err != nil {
		fmt.Println("could not make the request to the /users/reference end point")
	}

	return req
}

func (c *Client) AuthorizeCharge(reference string, event string) ([]byte, error) {
	charge := ChargeAuthorization{
		Event: event,
	}

	data, err := json.Marshal(charge)
	if err != nil {
		panic(err)
	}

	req, err := c.MakeHTTPRequest(c.baseUrl, "POST", "/authorization/"+reference, data)

	if err != nil {
		fmt.Println("could not get a response")
	}

	return req, nil
}
