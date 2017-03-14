package umeng

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetGroupList(t *testing.T) {
	c := New(NewConfig("YOU EMAIL", "PASSWORD"))

	in := &GetGroupListInput{
		AppKey:    "583bedb6e88bad684200140b",
		StartDate: "2017-03-10",
		EndDate:   "2017-03-10",
	}
	result, err := c.Events.GetGroupList(in)
	t.Log(result)
	assert.NoError(t, err)
}

func Test_GetGroupData(t *testing.T) {
	c := New(NewConfig("YOU EMAIL", "PASSWORD"))

	in := &GetDataInput{
		AppKey:    "583bedb6e88bad684200140b",
		StartDate: "2017-03-10",
		EndDate:   "2017-03-10",
		GroupId:   "585357ae5312dd614800031f",
	}
	result, err := c.Events.GetData(in)
	t.Log(result)
	assert.NoError(t, err)
}
