package model

type ResFlg struct {
	Status int    `json:"status"`
	Result string `json:"result"`
	User_ID    uint      `json:"user_id"`
	UserName   string    `json:"user_name"`
	Introduce  string    `json:"introduce"`
}

type ResOshiData struct {
	Status     int
	Result     string
	OshiId     int
	OshiName   string
	Birthday   string
	OshiMeet   string
	OshiLike1  string
	OshiLike2  string
	OshiLike3  string
	Free_Space string
	Interest   string
}
