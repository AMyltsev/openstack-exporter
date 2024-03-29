package main

import "time"

// Config struct
type config struct {
	ClusterName string `mapstructure:"cluster_name"`
	HorizonURL  string `mapstructure:"horizon_url"`
	Port        string `mapstructure:"port"`
	Openstack   openstackCred
	AuthURL     string
	ProbSec     time.Duration `mapstructure:"probe_period_sec"`
}

type openstackCred struct {
	Username string
	Password string
}

// Generated by https://quicktype.io

type authOpt struct {
	Auth auth `json:"auth"`
}

type auth struct {
	PasswordCredentials passwordCredentials `json:"passwordCredentials"`
	TenantName          string              `json:"tenantName"`
}

type passwordCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Generated by https://quicktype.io

type authResponse struct {
	Access access `json:"access"`
}

type access struct {
	Token          token            `json:"token"`
	ServiceCatalog []serviceCatalog `json:"serviceCatalog"`
	User           user             `json:"user"`
	Metadata       metadata         `json:"metadata"`
}

type metadata struct {
	IsAdmin int64    `json:"is_admin"`
	Roles   []string `json:"roles"`
}

type serviceCatalog struct {
	Endpoints      []endpoint    `json:"endpoints"`
	EndpointsLinks []interface{} `json:"endpoints_links"`
	Type           string        `json:"type"`
	Name           string        `json:"name"`
}

type endpoint struct {
	AdminURL    *string `json:"adminURL,omitempty"`
	Region      region  `json:"region"`
	InternalURL string  `json:"internalURL"`
	ID          string  `json:"id"`
	PublicURL   string  `json:"publicURL"`
}

type token struct {
	IssuedAt string     `json:"issued_at"`
	Expires  string     `json:"expires"`
	ID       string     `json:"id"`
	Tenant   tenantInfo `json:"tenant"`
	AuditIDS []string   `json:"audit_ids"`
}

type tenantInfo struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type user struct {
	Username   string        `json:"username"`
	RolesLinks []interface{} `json:"roles_links"`
	ID         string        `json:"id"`
	Roles      []role        `json:"roles"`
	Name       string        `json:"name"`
}

type role struct {
	Name string `json:"name"`
}

type region string

const (
	regionOne region = "RegionOne"
)

// Generated by https://quicktype.io

type novaServices struct {
	Services []service `json:"services"`
}

type service struct {
	ID             int64   `json:"id"`
	Binary         string  `json:"binary"`
	DisabledReason *string `json:"disabled_reason"`
	Host           string  `json:"host"`
	State          state   `json:"state"`
	Status         status  `json:"status"`
	UpdatedAt      string  `json:"updated_at"`
	ForcedDown     bool    `json:"forced_down"`
	Zone           string  `json:"zone"`
}

// Generated by https://quicktype.io

type neutronAgents struct {
	Agents []agent `json:"agents"`
}

type agent struct {
	Binary             string         `json:"binary"`
	Description        interface{}    `json:"description"`
	AdminStateUp       bool           `json:"admin_state_up"`
	HeartbeatTimestamp string         `json:"heartbeat_timestamp"`
	Alive              bool           `json:"alive"`
	Topic              string         `json:"topic"`
	Host               string         `json:"host"`
	AgentType          string         `json:"agent_type"`
	CreatedAt          string         `json:"created_at"`
	StartedAt          string         `json:"started_at"`
	ID                 string         `json:"id"`
	Configurations     configurations `json:"configurations"`
}

type configurations struct {
	LogAgentHeartbeats       bool          `json:"log_agent_heartbeats"`
	NovaMetadataPort         *int64        `json:"nova_metadata_port,omitempty"`
	NovaMetadataIP           *string       `json:"nova_metadata_ip,omitempty"`
	MetadataProxySocket      *string       `json:"metadata_proxy_socket,omitempty"`
	InDistributedMode        *bool         `json:"in_distributed_mode,omitempty"`
	ARPResponderEnabled      *bool         `json:"arp_responder_enabled,omitempty"`
	TunnelingIP              interface{}   `json:"tunneling_ip"`
	Devices                  *int64        `json:"devices,omitempty"`
	Extensions               []interface{} `json:"extensions"`
	L2Population             *bool         `json:"l2_population,omitempty"`
	TunnelTypes              []interface{} `json:"tunnel_types"`
	EnableDistributedRouting *bool         `json:"enable_distributed_routing,omitempty"`
	// BridgeMappings            []interface{}        `json:"bridge_mappings,omitempty"`
	RouterID                  *string     `json:"router_id,omitempty"`
	AgentMode                 interface{} `json:"agent_mode,omitempty"`
	GatewayExternalNetworkID  *string     `json:"gateway_external_network_id,omitempty"`
	HandleInternalOnlyRouters *bool       `json:"handle_internal_only_routers,omitempty"`
	UseNamespaces             *bool       `json:"use_namespaces,omitempty"`
	Routers                   *int64      `json:"routers,omitempty"`
	Interfaces                *int64      `json:"interfaces,omitempty"`
	FloatingIPS               *int64      `json:"floating_ips,omitempty"`
	InterfaceDriver           *string     `json:"interface_driver,omitempty"`
	ExternalNetworkBridge     *string     `json:"external_network_bridge,omitempty"`
	ExGwPorts                 *int64      `json:"ex_gw_ports,omitempty"`
	Subnets                   *int64      `json:"subnets,omitempty"`
	DHCPLeaseDuration         *int64      `json:"dhcp_lease_duration,omitempty"`
	DHCPDriver                *string     `json:"dhcp_driver,omitempty"`
	Networks                  *int64      `json:"networks,omitempty"`
	Ports                     *int64      `json:"ports,omitempty"`
}

const (
	dHCPAgent        string = "DHCP agent"
	l3Agent          string = "L3 agent"
	metadataAgent    string = "Metadata agent"
	openVSwitchAgent string = "Open vSwitch agent"
	na               string = "N/A"
	topicDHCPAgent   string = "dhcp_agent"
	topicL3Agent     string = "l3_agent"
)

// Generated by https://quicktype.io

type hypervisors struct {
	Hypervisors []hypervisor `json:"hypervisors"`
}

type hypervisor struct {
	Status             status            `json:"status"`
	Service            hypervisorService `json:"service"`
	VcpusUsed          int64             `json:"vcpus_used"`
	HypervisorType     string            `json:"hypervisor_type"`
	LocalGBUsed        int64             `json:"local_gb_used"`
	Vcpus              int64             `json:"vcpus"`
	HypervisorHostname string            `json:"hypervisor_hostname"`
	MemoryMBUsed       int64             `json:"memory_mb_used"`
	MemoryMB           int64             `json:"memory_mb"`
	CurrentWorkload    int64             `json:"current_workload"`
	State              state             `json:"state"`
	HostIP             string            `json:"host_ip"`
	CPUInfo            string            `json:"cpu_info"`
	RunningVms         int64             `json:"running_vms"`
	FreeDiskGB         int64             `json:"free_disk_gb"`
	HypervisorVersion  int64             `json:"hypervisor_version"`
	DiskAvailableLeast int64             `json:"disk_available_least"`
	LocalGB            int64             `json:"local_gb"`
	FreeRAMMB          int64             `json:"free_ram_mb"`
	ID                 int64             `json:"id"`
}

type hypervisorService struct {
	Host           string  `json:"host"`
	DisabledReason *string `json:"disabled_reason"`
	ID             int64   `json:"id"`
}

type state string

const (
	down state = "down"
	up   state = "up"
)

type status string

const (
	disabled status = "disabled"
	enabled  status = "enabled"
)

// Generated by https://quicktype.io

type images struct {
	Images []glanceImage `json:"images"`
	Schema string        `json:"schema"`
	First  string        `json:"first"`
}

type glanceImage struct {
	ImageState      *string       `json:"image_state,omitempty"`
	ContainerFormat string        `json:"container_format"`
	MinRAM          int64         `json:"min_ram"`
	RamdiskID       interface{}   `json:"ramdisk_id"`
	UpdatedAt       string        `json:"updated_at"`
	File            string        `json:"file"`
	Owner           string        `json:"owner"`
	ID              string        `json:"id"`
	Size            int64         `json:"size"`
	UserID          *string       `json:"user_id,omitempty"`
	ImageType       *string       `json:"image_type,omitempty"`
	Self            string        `json:"self"`
	DiskFormat      string        `json:"disk_format"`
	BaseImageRef    *string       `json:"base_image_ref,omitempty"`
	DirectURL       string        `json:"direct_url"`
	OwnerID         string        `json:"owner_id,omitempty"`
	Status          string        `json:"status"`
	ImageLocation   *string       `json:"image_location,omitempty"`
	Tags            []interface{} `json:"tags"`
	KernelID        interface{}   `json:"kernel_id"`
	Visibility      visibility    `json:"visibility"`
	MinDisk         int64         `json:"min_disk"`
	VirtualSize     interface{}   `json:"virtual_size"`
	InstanceUUID    *string       `json:"instance_uuid,omitempty"`
	Name            string        `json:"name"`
	RHELLicensed    *string       `json:"rhel_licensed,omitempty"`
	Checksum        string        `json:"checksum"`
	CreatedAt       string        `json:"created_at"`
	Protected       bool          `json:"protected"`
	Schema          string        `json:"schema"`
	MuranoImageInfo *string       `json:"murano_image_info,omitempty"`
	Description     *string       `json:"description,omitempty"`
	HypervisorType  *string       `json:"hypervisor_type,omitempty"`
}

type visibility string

const (
	private visibility = "private"
	public  visibility = "public"
)

const (
	nova     string = "nova"
	neutron  string = "neutron"
	glance   string = "glance"
	cinder   string = "cinderv2"
	keystone string = "keystone"
)
const (
	volumev2 string = "volumev2"
	image    string = "image"
	compute  string = "compute"
	identity string = "identity"
	network  string = "network"
)

// Generated by https://quicktype.io

type cinderServices struct {
	Services []cinderService `json:"services"`
}

type cinderService struct {
	Status         status      `json:"status"`
	Binary         string      `json:"binary"`
	Zone           string      `json:"zone"`
	State          state       `json:"state"`
	UpdatedAt      string      `json:"updated_at"`
	Host           string      `json:"host"`
	DisabledReason interface{} `json:"disabled_reason"`
}

// Generated by https://quicktype.io
type novaProjectsLimints []novaProjectLimits

type novaProjectLimits struct {
	Project tenant
	Limits  absolutNovaLimits
}

type novaLimit struct {
	Limits novaLimitsClass `json:"limits"`
}

type novaLimitsClass struct {
	Rate     []interface{}     `json:"rate"`
	Absolute absolutNovaLimits `json:"absolute"`
}

type absolutNovaLimits struct {
	MaxPersonality          int64 `json:"maxPersonality"`
	TotalServerGroupsUsed   int64 `json:"totalServerGroupsUsed"`
	MaxImageMeta            int64 `json:"maxImageMeta"`
	MaxServerMeta           int64 `json:"maxServerMeta"`
	MaxTotalKeypairs        int64 `json:"maxTotalKeypairs"`
	MaxPersonalitySize      int64 `json:"maxPersonalitySize"`
	MaxSecurityGroupRules   int64 `json:"maxSecurityGroupRules"`
	MaxServerGroups         int64 `json:"maxServerGroups"`
	TotalCoresUsed          int64 `json:"totalCoresUsed"`
	TotalRAMUsed            int64 `json:"totalRAMUsed"`
	TotalInstancesUsed      int64 `json:"totalInstancesUsed"`
	MaxSecurityGroups       int64 `json:"maxSecurityGroups"`
	TotalFloatingIpsUsed    int64 `json:"totalFloatingIpsUsed"`
	MaxTotalCores           int64 `json:"maxTotalCores"`
	MaxServerGroupMembers   int64 `json:"maxServerGroupMembers"`
	MaxTotalFloatingIps     int64 `json:"maxTotalFloatingIps"`
	TotalSecurityGroupsUsed int64 `json:"totalSecurityGroupsUsed"`
	MaxTotalInstances       int64 `json:"maxTotalInstances"`
	MaxTotalRAMSize         int64 `json:"maxTotalRAMSize"`
}

// Generated by https://quicktype.io

type projects struct {
	TenantsLinks []interface{} `json:"tenants_links"`
	Tenants      []tenant      `json:"tenants"`
}

type tenant struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

// Generated by https://quicktype.io
type cinderProjectsLimints []cinderProjectLimits

type cinderProjectLimits struct {
	Project tenant
	Limits  quotaSet
}

type tenantQuota struct {
	QuotaSet quotaSet `json:"quota_set"`
}

type quotaSet struct {
	SnapshotsDefault backupGigabytes `json:"snapshots_default"`
	GigabytesDefault backupGigabytes `json:"gigabytes_default"`
	Gigabytes        backupGigabytes `json:"gigabytes"`
	BackupGigabytes  backupGigabytes `json:"backup_gigabytes"`
	VolumesDefault   backupGigabytes `json:"volumes_default"`
	Snapshots        backupGigabytes `json:"snapshots"`
	Volumes          backupGigabytes `json:"volumes"`
	Backups          backupGigabytes `json:"backups"`
	ID               string          `json:"id"`
}

type backupGigabytes struct {
	Reserved int64 `json:"reserved"`
	Limit    int64 `json:"limit"`
	InUse    int64 `json:"in_use"`
}

type allProjectsANDinstances []projectANDinstances

type projectANDinstances struct {
	projects  tenant
	instances allProjectInstanses
}

// Generated by https://quicktype.io

type allProjectInstanses struct {
	Servers []server `json:"servers"`
}

type server struct {
	Status                           vmStatus                           `json:"status"`
	Updated                          string                             `json:"updated"`
	HostID                           string                             `json:"hostId"`
	OSEXTSRVATTRHost                 string                             `json:"OS-EXT-SRV-ATTR:host"`
	Addresses                        map[string][]address               `json:"addresses"`
	Links                            []link                             `json:"links"`
	KeyName                          *string                            `json:"key_name"`
	OSEXTSTSTaskState                interface{}                        `json:"OS-EXT-STS:task_state"`
	OSEXTSTSVMState                  string                             `json:"OS-EXT-STS:vm_state"`
	OSEXTSRVATTRInstanceName         string                             `json:"OS-EXT-SRV-ATTR:instance_name"`
	OSSRVUSGLaunchedAt               string                             `json:"OS-SRV-USG:launched_at"`
	OSEXTSRVATTRHypervisorHostname   string                             `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	Flavor                           flavor                             `json:"flavor"`
	ID                               string                             `json:"id"`
	SecurityGroups                   []securityGroup                    `json:"security_groups"`
	OSSRVUSGTerminatedAt             interface{}                        `json:"OS-SRV-USG:terminated_at"`
	OSEXTAZAvailabilityZone          string                             `json:"OS-EXT-AZ:availability_zone"`
	UserID                           string                             `json:"user_id"`
	Name                             string                             `json:"name"`
	Created                          string                             `json:"created"`
	TenantID                         string                             `json:"tenant_id"`
	OSDCFDiskConfig                  string                             `json:"OS-DCF:diskConfig"`
	OSExtendedVolumesVolumesAttached []oSExtendedVolumesVolumesAttached `json:"os-extended-volumes:volumes_attached"`
	AccessIPv4                       string                             `json:"accessIPv4"`
	AccessIPv6                       string                             `json:"accessIPv6"`
	Progress                         *int64                             `json:"progress,omitempty"`
	OSEXTSTSPowerState               int64                              `json:"OS-EXT-STS:power_state"`
	ConfigDrive                      string                             `json:"config_drive"`
	Metadata                         vmMetadata                         `json:"metadata"`
	// Image                            *vmImage                           `json:"image"`

}

type address struct {
	OSEXTIPSMACMACAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	Version            int64  `json:"version"`
	Addr               string `json:"addr"`
	OSEXTIPSType       string `json:"OS-EXT-IPS:type"`
}

type flavor struct {
	ID    string `json:"id"`
	Links []link `json:"links"`
}

type link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type vmMetadata struct {
}

type oSExtendedVolumesVolumesAttached struct {
	ID string `json:"id"`
}

type securityGroup struct {
	Name string `json:"name"`
}

type vmStatus string

const (
	statusACTIVE           vmStatus = "ACTIVE"
	statusERROR            vmStatus = "ERROR"
	statusBUILD            vmStatus = "BUILD"
	statusDELETED          vmStatus = "DELETED"
	statusHARDREBOOT       vmStatus = "HARD_REBOOT"
	statusMIGRATING        vmStatus = "MIGRATING"
	statusPASSWORD         vmStatus = "PASSWORD"
	statusPAUSED           vmStatus = "PAUSED"
	statusREBOOT           vmStatus = "REBOOT"
	statusREBUILD          vmStatus = "REBUILD"
	statusRESCUE           vmStatus = "RESCUE"
	statusRESIZE           vmStatus = "RESIZE"
	statusREVERTRESIZE     vmStatus = "REVERT_RESIZE"
	statusSHELVED          vmStatus = "SHELVED"
	statusSHELVEDOFFLOADED vmStatus = "SHELVED_OFFLOADED"
	statusSHUTOFF          vmStatus = "SHUTOFF"
	statusSOFTDELETED      vmStatus = "SOFT_DELETED"
	statusSUSPENDED        vmStatus = "SUSPENDED"
	statusUNKNOWN          vmStatus = "UNKNOWN"
	statusVERIFYRESIZE     vmStatus = "VERIFY_RESIZE"
)

// type vmImage struct {
// 	Flavor *flavor
// 	String *string
// }

// Generated by https://quicktype.io

type allTenantsUsage struct {
	TenantUsages []tenantUsage `json:"tenant_usages"`
}

type tenantUsage struct {
	name               string
	TotalMemoryMBUsage float64       `json:"total_memory_mb_usage"`
	TotalVcpusUsage    float64       `json:"total_vcpus_usage"`
	Start              string        `json:"start"`
	TenantID           string        `json:"tenant_id"`
	Stop               string        `json:"stop"`
	ServerUsages       []serverUsage `json:"server_usages"`
	TotalHours         float64       `json:"total_hours"`
	TotalLocalGBUsage  float64       `json:"total_local_gb_usage"`
}

type serverUsage struct {
	InstanceID string      `json:"instance_id"`
	Uptime     int64       `json:"uptime"`
	StartedAt  string      `json:"started_at"`
	EndedAt    interface{} `json:"ended_at"`
	MemoryMB   int64       `json:"memory_mb"`
	TenantID   string      `json:"tenant_id"`
	State      string      `json:"state"`
	Hours      float64     `json:"hours"`
	Vcpus      int64       `json:"vcpus"`
	Flavor     string      `json:"flavor"`
	LocalGB    int64       `json:"local_gb"`
	Name       string      `json:"name"`
}
