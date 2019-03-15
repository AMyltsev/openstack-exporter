package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Statuses of VM
var vmStatuses = []vmStatus{
	statusACTIVE,
	statusERROR,
	statusBUILD,
	statusDELETED,
	statusHARDREBOOT,
	statusMIGRATING,
	statusPASSWORD,
	statusPAUSED,
	statusREBOOT,
	statusREBUILD,
	statusRESCUE,
	statusRESIZE,
	statusREVERTRESIZE,
	statusSHELVED,
	statusSHELVEDOFFLOADED,
	statusSHUTOFF,
	statusSOFTDELETED,
	statusSUSPENDED,
	statusUNKNOWN,
	statusVERIFYRESIZE,
}

var vmStatusMetrics map[vmStatus]*prometheus.GaugeVec
var vmStatusCount map[vmStatus]int

// Nova servers statuses metrics map creation
func regSrvMetrics() {
	mm := make(map[vmStatus]*prometheus.GaugeVec)
	mc := make(map[vmStatus]int)
	for _, s := range vmStatuses {
		m := promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "openstack_project_instances_" + string(s),
			Help: "Count of instances in state " + string(s) + " for project",
		},
			[]string{"project"})
		mm[s] = m
		mc[s] = 0
	}
	vmStatusMetrics = mm
	vmStatusCount = mc
}

var (
	// Openstack API availebility metrics
	osAPIcheck = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_api",
		Help: " API availebility check",
	},
		[]string{"service"})

	// Nova services state metrics
	novaServiceState = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_nova_service_state",
		Help: "Nova services state",
	},
		[]string{"service", "host"})
	novaServiceStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_nova_service_status",
		Help: "Nova services status",
	},
		[]string{"service", "host"})

	// Nova hypervisors metrics
	novaHypervCPUused = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_hypervisor_vcpu_used",
		Help: "Hypervisor vCPU used",
	},
		[]string{"host"})
	novaHypervCPUtotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_hypervisor_vcpu_total",
		Help: "Hypervisor vCPU total",
	},
		[]string{"host"})
	novaHyperRAMused = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_hypervisor_ram_used",
		Help: "Hypervisor RAM used",
	},
		[]string{"host"})
	novaHyperRAMtotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_hypervisor_ram_total",
		Help: "Hypervisor RAM total",
	},
		[]string{"host"})
	novaHyperInstCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_hypervisor_instances_count",
		Help: "Count of instances on hypervisor",
	},
		[]string{"host"})

	// Nova project limits and usage
	novaProjectCPUquota = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_project_vcpu_quota",
		Help: "Project vCPU quota",
	},
		[]string{"project"})
	novaProjectCPUusage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_project_vcpu_usage",
		Help: "Project vCPU usage",
	},
		[]string{"project"})
	novaProjectRAMquota = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_project_ram_quota",
		Help: "Project RAM quota",
	},
		[]string{"project"})
	novaProjectRAMusage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_project_ram_usage",
		Help: "Project RAM usage",
	},
		[]string{"project"})
	novaProjectRootDiskusage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_project_Root_disk_usage",
		Help: "Project Root Disk usage",
	},
		[]string{"project"})

	// Neutron agents state metrics
	neutronAgentState = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_neutron_agent_state",
		Help: "Neutron services state",
	},
		[]string{"service", "host"})
	neutronAgentStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_neutron_agent_status",
		Help: "Neutron services status",
	},
		[]string{"service", "host"})

	// Cinder services state metrics
	cinderServiceState = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_cinder_service_state",
		Help: "Cinder services state",
	},
		[]string{"service", "host"})
	cinderServiceStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_cinder_service_status",
		Help: "Cinder services status",
	},
		[]string{"service", "host"})

	// Cinder project limits and usage
	cinderProjectVOLquota = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_project_volume_quota",
		Help: "Project Volume quota",
	},
		[]string{"project"})
	cinderProjectVOLusage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "openstack_project_volume_usage",
		Help: "Project Volume usage",
	},
		[]string{"project"})

	// Horizon metrics
	horizonAvailable = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "openstack_horizon_available",
		Help: "Status of Horizon availability, test with user login",
	})
)
