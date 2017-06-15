package sql

/*
map[string]interface {}{"client_ip":"172.19.0.5", "client_server":"",
"client_proc":"", "@timestamp":"2017-06-13T20:30:20.262Z",
"query":"SELECT `id`, `name`, `iso2`, `iso3`, `grc`, `linked` FROM `countries` AS `countries`;",
"responsetime":0, "method":"SELECT", "bytes_out":8787,
"mysql":map[string]interface {}{"insert_id":0, "num_rows":258, "num_fields":6,
	"iserror":false, "error_code":0, "error_message":"", "affected_rows":0},
"bytes_in":90, "client_port":41234, "type":"mysql", "port":3306,
"beat":map[string]interface {}{"hostname":"bnp-watv2", "version":"5.4.1", "name":"bnp-watv2"},
"@metadata":map[string]interface {}{"type":"mysql", "beat":"packetbeat"},
"status":"OK", "proc":"", "direction":"in", "path":"wat.countries", "ip":"10.20.125.140", "server":""}

*/

//easyjson:json
type Packetbeat struct {
	Type string `json:"type"`
}

//easyjson:json
type Mysql struct {
	ClientIP     string `json:"client_ip"`
	ClientServer string `json:"client_server"`
	ClientProc   string `json:"client_proc"`
	TimeStamp    string `json:"@timestamp"`
	Query        string `json:"query"`
	Method       string `json:"method"`
	BytesOut     int    `json:"bytes_out"`
	BytesIn      int    `json:"bytes_in"`
	Mysql        struct {
		InsertId     int  `json:"insert_id"`
		NumRows      int  `json:"num_rows"`
		NumFields    int  `json:"num_fields"`
		IsError      bool `json:"is_error"`
		ErrorCode    int  `json:error_code`
		AffectedRows int  `json:affected_rows`
	} `json:"mysql"`
	ClientPort int    `json:"client_port"`
	Type       string `json:"type"`
	Port       int    `json:"port"`
	Beat       struct {
		Hostname string `json:"hostname"`
		Version  string `json:"version"`
		Name     string `json:"name"`
	} `json:"beat"`
	Metadata struct {
		Type string `json:"type"`
		Beat string `json:"beat"`
	} `json:"@metadata"`
	ResponseTime float64 `json:"response_time"`
	Status       string  `json:"status"`
	IP           string  `json:"ip"`
	Proc         string  `json:"proc"`
	Direction    string  `json:"direction"`
	Path         string  `json:"path"`
	Server       string  `json:server`
}
