package e2e_test

import (
	"net/http"
	"testing"

	application "github.com/avila-r/social"
	"github.com/stretchr/testify/assert"
)

var (
	url = "http://" + application.Env.Get("SERVER_URL") + "/verify"
)

func Test_App(t *testing.T) {
	assert := assert.New(t)

	response, err := http.Get(url)

	if err != nil {
		t.Errorf(
			"GET request in %v throws an error - %v",
			url, err.Error(),
		)
	}

	defer response.Body.Close()

	assert.Equal(
		http.StatusOK, response.StatusCode,
	)
}
