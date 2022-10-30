package api

import (
	"bytes"
	"encoding/json"
	"go-user-service/data"
	"io/ioutil"
	"net/http"
)

func (app *Config) getUserService() (*data.Records, error) {
	url := "<NETLIFY FUNCTION GET URL>"

	records := data.Records{}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &records)
	if err != nil {
		return nil, err
	}

	return &records, nil
}

func (app *Config) createUserService(newUser data.User) (*data.CreateResponse, error) {
	url := "<NETLIFY FUNCTION CREATE URL>"

	response := data.CreateResponse{}
	jsonData := data.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		PhoneNumber: newUser.PhoneNumber,
	}

	postBody, _ := json.Marshal(jsonData)
	bodyData := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", bodyData)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
