package morpheus_test

import (
	"testing"
	"github.com/gomorpheus/morpheus-go-sdk"
)

func TestWhoami(t *testing.T) {
	client := getTestClient()
	resp, err := client.Whoami()
	assertResponse(t, resp, err)
	result := resp.Result.(*morpheus.WhoamiResult)
	assertNotNil(t, result.User)
	assertNotNil(t, result.User.ID)
	assertEqual(t, result.User.Username, testUsername)
}
