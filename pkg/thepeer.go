package pkg

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type User struct {
	name       string
	identifier string
	email      string
}

const API_LINK = "https://api.thepeer.co"

func ValidateSignature(payload []byte, signature interface{}, secretKey []byte) bool {

	// Generate the HMAC signature
	h := hmac.New(sha1.New, secretKey)
	h.Write(payload)

	expectedSignatureString := hex.EncodeToString(h.Sum(nil))

	return signature == expectedSignatureString
}

func IndexUser(name string, identifier string, email string) {
	user := User{
		name:       name,
		identifier: identifier,
		email:      email,
	}

	body, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	response, err := http.NewRequest("POST", API_LINK+"/users", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	// Look at using secret key globally
	secretKey := ""
	response.Header.Add("x-api-key", secretKey)
	response.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(response)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		panic(res.Status)
	}
}

func UpdateUser(reference string, identifier string) {

}
