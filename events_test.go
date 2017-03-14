package umeng

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetGroupList(t *testing.T) {
	c := New(NewConfig("YOU EMAIL", "PASSWORD"))

	in := &GetGroupListInput{
		AppKey:    "YOU APP KEY",
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
		AppKey:    "YOU APP KEY",
		StartDate: "2017-03-10",
		EndDate:   "2017-03-10",
		GroupId:   "YOU APP GROUP ID",
	}
	result, err := c.Events.GetData(in)
	t.Log(result)
	assert.NoError(t, err)
}
