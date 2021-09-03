package repository

type rapiGroupResponse struct {
	UUID        string        `json:"uuid"`
	Name        string        `json:"name"`
	Tags        []interface{} `json:"tags"`
	NodeCnt     int           `json:"node_cnt"`
	AllocPolicy string        `json:"alloc_policy"`
	NodeList    []string      `json:"node_list"`
	SerialNo    int           `json:"serial_no"`
}

type rapiGroupsResponse []rapiGroupResponse
