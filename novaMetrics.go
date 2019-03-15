package main

func (s novaServices) Metrics() {
	novaServiceState.Reset()
	novaServiceStatus.Reset()
	for _, serv := range s.Services {
		if serv.State == up {
			novaServiceState.WithLabelValues(serv.Binary, serv.Host).Set(1)
		}
		if serv.State == down {
			novaServiceState.WithLabelValues(serv.Binary, serv.Host).Set(0)
		}
		if serv.Status == enabled {
			novaServiceStatus.WithLabelValues(serv.Binary, serv.Host).Set(1)
		}
		if serv.Status == disabled {
			novaServiceStatus.WithLabelValues(serv.Binary, serv.Host).Set(0)
		}
	}

}

func (s hypervisors) Metrics() {
	novaHypervCPUused.Reset()
	novaHypervCPUtotal.Reset()
	novaHyperRAMused.Reset()
	novaHyperRAMtotal.Reset()
	novaHyperInstCount.Reset()
	for _, h := range s.Hypervisors {
		if h.State == down {
			novaHypervCPUtotal.WithLabelValues(h.HypervisorHostname).Set(0)
			novaHypervCPUused.WithLabelValues(h.HypervisorHostname).Set(0)
			novaHyperRAMtotal.WithLabelValues(h.HypervisorHostname).Set(0)
			novaHyperRAMused.WithLabelValues(h.HypervisorHostname).Set(0)
			novaHyperInstCount.WithLabelValues(h.HypervisorHostname).Set(0)
		}
		if h.State == up {
			novaHypervCPUtotal.WithLabelValues(h.HypervisorHostname).Set(float64(h.Vcpus))
			novaHypervCPUused.WithLabelValues(h.HypervisorHostname).Set(float64(h.VcpusUsed))
			novaHyperRAMtotal.WithLabelValues(h.HypervisorHostname).Set(float64(h.MemoryMB))
			novaHyperRAMused.WithLabelValues(h.HypervisorHostname).Set(float64(h.MemoryMBUsed))
			novaHyperInstCount.WithLabelValues(h.HypervisorHostname).Set(float64(h.RunningVms))
		}
	}
}

func (s novaProjectsLimints) Metrics() {
	novaProjectCPUquota.Reset()
	novaProjectCPUusage.Reset()
	novaProjectRAMquota.Reset()
	novaProjectRAMusage.Reset()
	for _, h := range s {
		novaProjectCPUquota.WithLabelValues(h.Project.Name).Set(float64(h.Limits.MaxTotalCores))
		novaProjectCPUusage.WithLabelValues(h.Project.Name).Set(float64(h.Limits.TotalCoresUsed))
		novaProjectRAMquota.WithLabelValues(h.Project.Name).Set(float64(h.Limits.MaxTotalRAMSize))
		novaProjectRAMusage.WithLabelValues(h.Project.Name).Set(float64(h.Limits.TotalRAMUsed))
	}
}

func (s allProjectsANDinstances) Metrics() {
	for _, projecANDinstances := range s {
		// Reset counters before calculation
		for status := range vmStatusCount {
			vmStatusCount[status] = 0
		}
		for _, instance := range projecANDinstances.instances.Servers {
			vmStatusCount[instance.Status]++
		}
		for s, c := range vmStatusCount {
			metric := vmStatusMetrics[s]
			metric.WithLabelValues(projecANDinstances.projects.Name).Set(float64(c))
		}
	}

}

func (s allTenantsUsage) Metrics() {
	novaProjectRootDiskusage.Reset()
	for _, h := range s.TenantUsages {
		local := int64(0)
		for _, l := range h.ServerUsages {
			local += l.LocalGB
		}
		novaProjectRootDiskusage.WithLabelValues(h.name).Set(float64(local))
	}
}
