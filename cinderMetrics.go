package main

func (s cinderServices) Metrics() {
	cinderServiceState.Reset()
	cinderServiceStatus.Reset()
	for _, serv := range s.Services {
		if serv.State == up {
			cinderServiceState.WithLabelValues(serv.Binary, serv.Host).Set(1)
		}
		if serv.State == down {
			cinderServiceState.WithLabelValues(serv.Binary, serv.Host).Set(0)
		}
		if serv.Status == enabled {
			cinderServiceStatus.WithLabelValues(serv.Binary, serv.Host).Set(1)
		}
		if serv.Status == disabled {
			cinderServiceStatus.WithLabelValues(serv.Binary, serv.Host).Set(0)
		}
	}

}

func (s cinderProjectsLimints) Metrics() {
	cinderProjectVOLquota.Reset()
	cinderProjectVOLusage.Reset()
	for _, h := range s {
		cinderProjectVOLquota.WithLabelValues(h.Project.Name).Set(float64(h.Limits.Gigabytes.Limit))
		cinderProjectVOLusage.WithLabelValues(h.Project.Name).Set(float64(h.Limits.Gigabytes.InUse))
	}
}
