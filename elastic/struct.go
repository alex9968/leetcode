package main

type WebLog struct {
	ID         string   `json:"_id"`
	HttpMethod string   `json:"http_method"`
	Uri        string   `json:"uri"`
	Stem       string   `json:"uri-stem"`
	Param      string   `json:"uri-query"`
	BodyLength string   `json:"body_length"`
	ClientIp   string   `json:"remote_ip"`
	ClientUser string   `json:"remote_user"`
	Time       int64    `json:"time"`
	StatusCode int64    `json:"status_code"`
	Referer    string   `json:"referer"`
	UserAgent  string   `json:"user_agent"`
	AttackRule []string `json:"attack"`
	Location   string   `json:"location"`
	Request    string   `json:"request,omitempty"`
	LogicUser  string   `json:"logic_user,omitempty"`
}

type Event struct {
	System struct {
		EventID     int64 `json:"event_id"`
		TimeCreated int64 `json:"TimeCreated"`
		Security    interface{}
	} `json:"System"`
	EventData map[string]interface{}
}
type WinLog struct {
	AttackRule []string `json:"attack_rule"`
	Event      Event
}

type EsReq struct {
	Indexs       []string
	From         int64
	Size         int64
	Fields       map[string]interface{}
	SubFieldName []string
	StartTime    int64
	EndTime      int64
	Count        int64
	IsAttack     bool
}

type EsRes struct {
	Took   int64
	Size   int64
	Result map[string]interface{}
	Hits   string
}

type RespBucket struct {
	TotalDocCount int64 `json:"total_doc_count"`
	Took          int   `json:"took"`
	Group         Group `json:"group"`
}

type Group struct {
	DocCountErrorUpperBound int64     `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int64     `json:"sum_other_doc_count"`
	Buckets                 []*Bucket `json:"buckets"`
}

type Bucket struct {
	Key      string                 `json:"key"`
	DocCount int64                  `json:"doc_count"`
	Percent  string                 `json:"percent"`
	Group    *Group                 `json:"group,omitempty"`
	Location map[string]interface{} `json:"location,omitempty"`
	Threat   interface{}            `json:"threat,omitempty"`
}
