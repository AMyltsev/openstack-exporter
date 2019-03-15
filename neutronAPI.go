package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

func getNeutronAgents(authRes authResponse) (neutronAgents, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	h := neutronAgents{}
	netEnd, err := authRes.getEndpoint(network)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Neutron endpoint didn't find")
		return neutronAgents{}, err
	}
	req, err := http.NewRequest(
		"GET", netEnd+"/v2.0/agents", nil,
	)
	req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return neutronAgents{}, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&h); err != nil {
		fmt.Println(err)
		log.Errorln("Cannot decode Neutron agents json")
		return neutronAgents{}, err
	}

	return h, err
}
