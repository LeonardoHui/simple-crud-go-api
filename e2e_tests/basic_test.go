package e2etests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddBook_addsBookSuccessfuly(t *testing.T) {
	jsonStr := []byte(`{
		"title":"harry potter and philosopher's stone",
		"Author": "J.K.Rowling", 
		"description": "It is a story about Harry Potter"
		}`)
	req, err := http.NewRequest("POST", "http://localhost:4000/books", bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "201 Created", resp.Status)
	assert.Contains(t, string(body), "Created")
	assert.Equal(t, nil, err)
}
