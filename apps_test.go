package umeng

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func client() *Client {
	return New(NewConfig("YOU_EMAIL", "YOU_PASSWORD"))
}

func TestAppss_GetApps(t *testing.T) {
	c := client()
	result, err := c.Apps.GetApps()
	t.Log(result[0])
	assert.NoError(t, err)
}

func TestAppss_GetBaseData(t *testing.T) {
	c := client()
	result, err := c.Apps.GetBaseData()
	t.Log(result)
	assert.NoError(t, err)
}
