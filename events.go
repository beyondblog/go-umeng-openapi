package umeng

import (
	"encoding/json"
)

type Events struct {
	*Client
}

func NewEvents(config *Config) *Apps {
	return &Apps{
		Client: &Client{
			Config: config,
		},
	}
}

type GetGroupListOutput struct {
	Count       int64  `json:"count"`
	Name        string `json:"name"`
	GroupId     string `json:"group_id"`
	DisplayName string `json:"display_name"`
}

type GetGroupListInput struct {
	AppKey     string `url:"appkey"` //app 标识
	StartDate  string `url:"start_date,omitempty"`
	EndDate    string `url:"end_date,omitempty"`
	PeriodType string `url:"period_type,omitempty"`
	Versions   string `url:"versions,omitempty"` //版本号
}

//获取自定义事件Group 列表
func (c *Events) GetGroupList(in *GetGroupListInput) (out []*GetGroupListOutput, err error) {
	body, err := c.get("/events/group_list", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

type GetDataOutput struct {
	Data struct {
		All []int64 `json:"all"`
	} `json:"data"`

	Dates []string `json:"dates,omitempty"`
}

type GetDataInput struct {
	AppKey     string `url:"appkey"` //app 标识
	StartDate  string `url:"start_date,omitempty"`
	EndDate    string `url:"end_date,omitempty"`
	PeriodType string `url:"period_type,omitempty"`
	GroupId    string `url:"group_id,omitempty"`
	Type       string `url:"type,omitempty"`
}

//获取事件消息数
func (c *Events) GetData(in *GetDataInput) (out *GetDataOutput, err error) {
	body, err := c.get("/events/daily_data", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
