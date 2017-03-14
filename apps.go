package umeng

import (
	"encoding/json"
)

type Apps struct {
	*Client
}

func NewApps(config *Config) *Apps {
	return &Apps{
		Client: &Client{
			Config: config,
		},
	}
}

// GetAppsOutput request output.
type GetAppsOutput struct {
	Name      string `json:"name"`
	Categroy  string `json:"category"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Platform  string `json:"platform"`
	Appkey    string `json:"appkey"`
}

type GetAppBaseDataOutput struct {
	ToDayActiveUsers     int64 `json:"today_active_users"`     //今日活跃用户
	YesterDayActiveUsers int64 `json:"yesterday_active_users"` //今日活跃用户
	ToDayLaunches        int64 `json:"today_launches"`         //今日启动次数
	YesterdayLaunches    int64 `json:"yesterday_launches"`     //昨日启动次数
	ToDayNewUsers        int64 `json:"today_new_users"`        //今日新增用户
	YesterDayNewUsers    int64 `json:"yesterday_new_users"`    //昨日新增用户
}

//获取apps列表
func (c *Apps) GetApps() (out []*GetAppsOutput, err error) {
	body, err := c.get("/apps", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

//获取基础数据
func (c *Apps) GetBaseData() (out *GetAppBaseDataOutput, err error) {

	body, err := c.get("/apps/base_data", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
