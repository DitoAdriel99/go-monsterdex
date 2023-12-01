package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestGetBreedsHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:9999/breeds/list/all")
	s.Equal(http.StatusOK, r.StatusCode)
}
func (s *EndToEndSuite) TestGetSubBreedHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:9999/breed/hound/list")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndToEndSuite) TestCreateDogHandler() {
	client := &http.Client{}
	payload := []byte(`{
		"name": "lala",
		"breeds": "hound",
		"gender": "female",
		"image": "httppp"
	}`)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:9999/dogs", bytes.NewBuffer(payload))
	if err != nil {
		s.Fail("Failed to create request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		s.Fail("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.Fail("Failed to read response body:", err)
		return
	}
	s.Equal(http.StatusCreated, resp.StatusCode)
	fmt.Println("Response body:", string(body))

}
