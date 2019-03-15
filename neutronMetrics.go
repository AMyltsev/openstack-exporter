package main

func (s neutronAgents) Metrics() {
	neutronAgentState.Reset()
	neutronAgentStatus.Reset()
	for _, serv := range s.Agents {
		if serv.Alive == true {
			neutronAgentState.WithLabelValues(serv.Binary, serv.Host).Set(1)
		}
		if serv.Alive == false {
			neutronAgentState.WithLabelValues(serv.Binary, serv.Host).Set(0)
		}
		if serv.Alive == true {
			neutronAgentStatus.WithLabelValues(serv.Binary, serv.Host).Set(1)
		}
		if serv.Alive == false {
			neutronAgentStatus.WithLabelValues(serv.Binary, serv.Host).Set(0)
		}
	}

}
