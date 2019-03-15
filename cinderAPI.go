package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

func getCinderServices(authRes authResponse) (cinderServices, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	h := cinderServices{}
	netEnd, err := authRes.getEndpoint(volumev2)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Cinder endpoint didn't find")
		return cinderServices{}, err
	}
	req, err := http.NewRequest(
		"GET", netEnd+"/os-services", nil,
	)
	req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return cinderServices{}, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&h); err != nil {
		fmt.Println(err)
		log.Errorln("Cannot decode cinder services json")
	}

	return h, err
}

func getCinderLimits(authRes authResponse) (cinderProjectsLimints, error) {
	client := &http.Client{}
	h := cinderProjectsLimints{}
	nl := tenantQuota{}
	cinderEnd, err := authRes.getEndpoint(volumev2)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Cinder endpoint didn't find")
		return cinderProjectsLimints{}, err
	}
	projects, err := getProjects(authRes)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Can't get project list")
		return cinderProjectsLimints{}, err
	}

	for _, p := range projects.Tenants {
		pl := cinderProjectLimits{}
		req, err := http.NewRequest(
			"GET", cinderEnd+"/os-quota-sets/"+p.ID+"?usage=True", nil,
		)
		req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return cinderProjectsLimints{}, err
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&nl); err != nil {
			fmt.Println(err)
			log.Errorln("Cannot decode cinder limits json")
			return cinderProjectsLimints{}, err
		}
		pl.Limits = nl.QuotaSet
		pl.Project = p
		h = append(h, pl)
	}
	return h, err
}
