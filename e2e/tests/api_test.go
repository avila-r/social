package e2e_test

import (
	"net/http"
	"testing"

	"github.com/avila-r/social/e2e"
	"github.com/stretchr/testify/assert"
)

var (
	url = "http://" + e2e.Env.Get("API_URL")
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
