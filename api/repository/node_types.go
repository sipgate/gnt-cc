package repository

type rapiNodeResponse struct {
	Name            string      `json:"name"`
	PinstCnt        int         `json:"pinst_cnt"`
	SinstCnt        int         `json:"sinst_cnt"`
	PinstList       []string    `json:"pinst_list"`
	SinstList       []string    `json:"sinst_list"`
	Dtotal          int         `json:"dtotal"`
	Dfree           int         `json:"dfree"`
	Ctotal          int         `json:"ctotal"`
	Mtotal          int         `json:"mtotal"`
	Mfree           int         `json:"mfree"`
	Sptotal         interface{} `json:"sptotal"`
	Spfree          interface{} `json:"spfree"`
	Offline         bool        `json:"offline"`
	Drained         bool        `json:"drained"`
	VMCapable       bool        `json:"vm_capable"`
	Master          bool        `json:"master"`
	MasterCandidate bool        `json:"master_candidate"`
	MasterCapable   bool        `json:"master_capable"`
	Mnode           int         `json:"mnode"`
	Cnodes          int         `json:"cnodes"`
	Ctime           float64     `json:"ctime"`
	Mtime           float64     `json:"mtime"`
	SerialNo        int         `json:"serial_no"`
	Pip             string      `json:"pip"`
	Sip             string      `json:"sip"`
	UUID            string      `json:"uuid"`
	Csockets        int         `json:"csockets"`
	Role            string      `json:"role"`
	Tags            []string    `json:"tags"`
	Group           string      `json:"group"`
	GroupUUID       string      `json:"group.uuid"`
	Cnos            int         `json:"cnos"`
	Ndparams        struct {
		Ovs              bool   `json:"ovs"`
		SSHPort          int    `json:"ssh_port"`
		OvsLink          string `json:"ovs_link"`
		SpindleCount     int    `json:"spindle_count"`
		ExclusiveStorage bool   `json:"exclusive_storage"`
		CPUSpeed         int    `json:"cpu_speed"`
		OvsName          string `json:"ovs_name"`
		OobProgram       string `json:"oob_program"`
	} `json:"ndparams"`
}
