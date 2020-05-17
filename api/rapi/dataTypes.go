package rapi

// -> https://mholt.github.io/json-to-go/

type Job struct {
	Status     string        `json:"status"`
	ID         int           `json:"id"`
	StartTs    []int         `json:"start_ts"`
	EndTs      []int         `json:"end_ts"`
	ReceivedTs []int         `json:"received_ts"`
	Summary    []string      `json:"summary"`
	OpStatus   []string      `json:"opstatus"`
	Ops        []interface{} `json:"ops"`
}

type JobsBulk []Job

type Node struct {
	Name            string      `json:"name"`
	PinstCnt        int         `json:"pinst_cnt"`
	SinstCnt        int         `json:"sinst_cnt"`
	PinstList       []string    `json:"pinst_list"`
	SinstList       []string    `json:"sinst_list"`
	Dtotal          interface{} `json:"dtotal"`
	Dfree           interface{} `json:"dfree"`
	Ctotal          int         `json:"ctotal"`
	Mtotal          int         `json:"mtotal"`
	Mfree           int         `json:"mfree"`
	Sptotal         interface{} `json:"sptotal"`
	Spfree          interface{} `json:"spfree"`
	Offline         bool        `json:"offline"`
	Drained         bool        `json:"drained"`
	VMCapable       bool        `json:"vm_capable"`
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

type NodesBulk []Node

type Instance struct {
	Name             string        `json:"name"`
	OperState        bool          `json:"oper_state"`
	AdminState       string        `json:"admin_state"`
	Status           string        `json:"status"`
	Pnode            string        `json:"pnode"`
	Snodes           []string      `json:"snodes"`
	Ctime            float64       `json:"ctime"`
	Mtime            float64       `json:"mtime"`
	CustomNicparams  []GanetiNic   `json:"custom_nicparams"`
	DiskNames        []interface{} `json:"disk.names"`
	DiskSizes        []interface{} `json:"disk.sizes"`
	DiskSpindles     []interface{} `json:"disk.spindles"`
	DiskTemplate     string        `json:"disk_template"`
	DiskUsage        int           `json:"disk_usage"`
	DiskUuids        []interface{} `json:"disk.uuids"`
	NetworkPort      int           `json:"network_port"`
	NicBridges       []interface{} `json:"nic.bridges"`
	NicIps           []interface{} `json:"nic.ips"`
	NicLinks         []interface{} `json:"nic.links"`
	NicMacs          []interface{} `json:"nic.macs"`
	NicModes         []interface{} `json:"nic.modes"`
	NicNames         []interface{} `json:"nic.names"`
	NicNetworks      []interface{} `json:"nic.networks"`
	NicNetworksNames []interface{} `json:"nic.networks.names"`
	NicUuids         []interface{} `json:"nic.uuids"`
	OperRAM          interface{}   `json:"oper_ram"`
	OperVcpus        interface{}   `json:"oper_vcpus"`
	Os               string        `json:"os"`
	SerialNo         int           `json:"serial_no"`
	Tags             []string      `json:"tags"`
	UUID             string        `json:"uuid"`
	BeParams         BEParams      `json:"beparams"`
	CustomBeParams   BEParams      `json:"custom_beparams"`
	HvParams         HVParams      `json:"hvparams"`
	CustomHvParams   HVParams      `json:"custom_hvparams"`
}

type BEParams struct {
	AutoBalance    bool `json:"auto_balance,omitempty"`
	SpindleUse     int  `json:"spindle_use,omitempty"`
	Vcpus          int  `json:"vcpus,omitempty"`
	Memory         int  `json:"memory,omitempty"`
	Minmem         int  `json:"minmem,omitempty"`
	AlwaysFailover bool `json:"always_failover,omitempty"`
	Maxmem         int  `json:"maxmem,omitempty"`
}

type HVParams struct {
	Acpi                       bool   `json:"acpi,omitempty"`
	BlockdevPrefix             string `json:"blockdev_prefix,omitempty"`
	BootOrder                  string `json:"boot_order,omitempty"`
	BootloaderArgs             string `json:"bootloader_args,omitempty"`
	BootloaderPath             string `json:"bootloader_path,omitempty"`
	Cdrom2ImagePath            string `json:"cdrom2_image_path,omitempty"`
	CdromDiskType              string `json:"cdrom_disk_type,omitempty"`
	CdromImagePath             string `json:"cdrom_image_path,omitempty"`
	CPUCap                     string `json:"cpu_cap,omitempty"`
	CPUCores                   int    `json:"cpu_cores,omitempty"`
	CPUMask                    string `json:"cpu_mask,omitempty"`
	CPUSockets                 int    `json:"cpu_sockets,omitempty"`
	CPUThreads                 int    `json:"cpu_threads,omitempty"`
	CPUType                    string `json:"cpu_type,omitempty"`
	CPUWeight                  string `json:"cpu_weight,omitempty"`
	Cpuid                      string `json:"cpuid,omitempty"`
	DeviceModel                string `json:"device_model,omitempty"`
	Devices                    string `json:"devices,omitempty"`
	DiskAio                    string `json:"disk_aio,omitempty"`
	DiskCache                  string `json:"disk_cache,omitempty"`
	DiskType                   string `json:"disk_type,omitempty"`
	DropCapabilities           string `json:"drop_capabilities,omitempty"`
	ExtraCgroups               string `json:"extra_cgroups,omitempty"`
	ExtraConfig                string `json:"extra_config,omitempty"`
	FloppyImagePath            string `json:"floppy_image_path,omitempty"`
	InitScript                 string `json:"init_script,omitempty"`
	InitrdPath                 string `json:"initrd_path,omitempty"`
	KernelArgs                 string `json:"kernel_args,omitempty"`
	KernelPath                 string `json:"kernel_path,omitempty"`
	Keymap                     string `json:"keymap,omitempty"`
	KvmExtra                   string `json:"kvm_extra,omitempty"`
	KvmFlag                    string `json:"kvm_flag,omitempty"`
	KvmPath                    string `json:"kvm_path,omitempty"`
	MachineVersion             string `json:"machine_version,omitempty"`
	MemPath                    string `json:"mem_path,omitempty"`
	MigrationCaps              string `json:"migration_caps,omitempty"`
	MigrationDowntime          int    `json:"migration_downtime,omitempty"`
	NicType                    string `json:"nic_type,omitempty"`
	NumTtys                    string `json:"num_ttys,omitempty"`
	Pae                        string `json:"pae,omitempty"`
	PciPass                    string `json:"pci_pass,omitempty"`
	RebootBehavior             string `json:"reboot_behavior,omitempty"`
	RootPath                   string `json:"root_path,omitempty"`
	SecurityDomain             string `json:"security_domain,omitempty"`
	SecurityModel              string `json:"security_model,omitempty"`
	SerialConsole              bool   `json:"serial_console,omitempty"`
	SerialSpeed                int    `json:"serial_speed,omitempty"`
	Soundhw                    string `json:"soundhw,omitempty"`
	SpiceBind                  string `json:"spice_bind,omitempty"`
	SpiceImageCompression      string `json:"spice_image_compression,omitempty"`
	SpiceIPVersion             int    `json:"spice_ip_version,omitempty"`
	SpiceJpegWanCompression    string `json:"spice_jpeg_wan_compression,omitempty"`
	SpicePasswordFile          string `json:"spice_password_file,omitempty"`
	SpicePlaybackCompression   bool   `json:"spice_playback_compression,omitempty"`
	SpiceStreamingVideo        string `json:"spice_streaming_video,omitempty"`
	SpiceTLSCiphers            string `json:"spice_tls_ciphers,omitempty"`
	SpiceUseTLS                bool   `json:"spice_use_tls,omitempty"`
	SpiceUseVdagent            bool   `json:"spice_use_vdagent,omitempty"`
	SpiceZlibGlzWanCompression string `json:"spice_zlib_glz_wan_compression,omitempty"`
	StartupTimeout             string `json:"startup_timeout,omitempty"`
	UsbDevices                 string `json:"usb_devices,omitempty"`
	UsbMouse                   string `json:"usb_mouse,omitempty"`
	UseBootloader              string `json:"use_bootloader,omitempty"`
	UseChroot                  bool   `json:"use_chroot,omitempty"`
	UseLocaltime               bool   `json:"use_localtime,omitempty"`
	UserShutdown               bool   `json:"user_shutdown,omitempty"`
	Vga                        string `json:"vga,omitempty"`
	VhostNet                   bool   `json:"vhost_net,omitempty"`
	VifScript                  string `json:"vif_script,omitempty"`
	VifType                    string `json:"vif_type,omitempty"`
	Viridian                   string `json:"viridian,omitempty"`
	VirtioNetQueues            int    `json:"virtio_net_queues,omitempty"`
	VncBindAddress             string `json:"vnc_bind_address,omitempty"`
	VncPasswordFile            string `json:"vnc_password_file,omitempty"`
	VncTLS                     bool   `json:"vnc_tls,omitempty"`
	VncX509Path                string `json:"vnc_x509_path,omitempty"`
	VncX509Verify              bool   `json:"vnc_x509_verify,omitempty"`
	VnetHdr                    bool   `json:"vnet_hdr,omitempty"`
}

type InstancesBulk []Instance

type InstanceCreate struct {
	Version        int          `json:"__version__"`
	InstanceName   string       `json:"instance_name"`
	ConflictsCheck bool         `json:"conflicts_check"`
	IPCheck        bool         `json:"ip_check"`
	NameCheck      bool         `json:"name_check"`
	WaitForSync    bool         `json:"wait_for_sync"`
	NoInstall      bool         `json:"no_install"`
	DiskTemplate   string       `json:"disk_template"`
	Mode           string       `json:"mode"`
	Hypervisor     string       `json:"hypervisor"`
	Iallocator     string       `json:"iallocator"`
	OsType         string       `json:"os_type"`
	Disks          []GanetiDisk `json:"disks"`
	Nics           []GanetiNic  `json:"nics"`
	HvParams       HVParams     `json:"hvparams"`
	BeParams       BEParams     `json:"beparams"`
}

type InstanceCreateBulk []InstanceCreate

type GanetiNic struct {
	Bridge  string `json:"bridge,omitempty"`
	Name    string `json:"name,omitempty"`
	IP      string `json:"ip,omitempty"`
	Vlan    string `json:"vlan,omitempty"`
	Link    string `json:"link,omitempty"`
	Mode    string `json:"mode,omitempty"`
	Network string `json:"network,omitempty"`
}

type GanetiDisk struct {
	Size int `json:"size"`
}

type JobStatusCount struct {
	Canceled int `json:"canceled"`
	Error    int `json:"error"`
	Pending  int `json:"pending"`
	Queued   int `json:"queued"`
	Success  int `json:"success"`
	Waiting  int `json:"waiting"`
}
