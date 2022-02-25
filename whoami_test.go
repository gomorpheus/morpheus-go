package morpheus_test

import (
	"github.com/gomorpheus/morpheus-go-sdk"
	"testing"
)

func TestWhoami(t *testing.T) {
	client := getTestClient(t)
	resp, err := client.Whoami()
	assertResponse(t, resp, err)
	result := resp.Result.(*morpheus.WhoamiResult)
	assertNotNil(t, result.User)
	assertNotNil(t, result.User.ID)
	assertEqual(t, result.User.Username, testUsername)
}
