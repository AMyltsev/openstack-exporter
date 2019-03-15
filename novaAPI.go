package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/prometheus/common/log"
)

func getNovaServices(authRes authResponse) (novaServices, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	h := novaServices{}
	novaEnd, err := authRes.getEndpoint(compute)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Nova endpoint didn't find")
		return novaServices{}, err
	}
	req, err := http.NewRequest(
		"GET", novaEnd+"/os-services", nil,
	)
	req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return novaServices{}, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&h); err != nil {
		fmt.Println(err)
		log.Errorln("Cannot decode Nova services json")
	}

	return h, err
}

func getNovaHypervisors(authRes authResponse) (hypervisors, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	h := hypervisors{}
	novaEnd, err := authRes.getEndpoint(compute)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Nova endpoint didn't find")
		return hypervisors{}, err
	}
	req, err := http.NewRequest(
		"GET", novaEnd+"/os-hypervisors/detail", nil,
	)
	req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return hypervisors{}, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&h); err != nil {
		fmt.Println(err)
		log.Errorln("Cannot decode Nova hypervisors json")
	}

	return h, err
}

func getNovaLimits(authRes authResponse) (novaProjectsLimints, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	h := novaProjectsLimints{}
	nl := novaLimit{}
	novaEnd, err := authRes.getEndpoint(compute)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Nova endpoint didn't find")
		return novaProjectsLimints{}, err
	}
	projects, err := getProjects(authRes)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Can't get project list")
		return novaProjectsLimints{}, err
	}

	for _, p := range projects.Tenants {
		pl := novaProjectLimits{}
		req, err := http.NewRequest(
			"GET", novaEnd+"/limits?tenant_id="+p.ID, nil,
		)
		req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return novaProjectsLimints{}, err
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&nl); err != nil {
			fmt.Println(err)
			log.Errorln("Cannot decode Nova limits json")
		}
		pl.Limits = nl.Limits.Absolute
		pl.Project = p
		h = append(h, pl)
	}
	return h, err
}

func getServersInfo(authRes authResponse) (allProjectsANDinstances, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	h := allProjectsANDinstances{}
	novaEnd, err := authRes.getEndpoint(compute)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Nova endpoint didn't find")
		return allProjectsANDinstances{}, err
	}
	projects, err := getProjects(authRes)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Can't get project list")
		return allProjectsANDinstances{}, err
	}

	for _, p := range projects.Tenants {
		i := allProjectInstanses{}
		pai := projectANDinstances{}
		req, err := http.NewRequest(
			"GET", novaEnd+"/servers/detail?all_tenants=true&tenant_id="+p.ID, nil,
		)
		req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return allProjectsANDinstances{}, err
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&i); err != nil {
			fmt.Println(err)
			log.Errorln("Cannot decode \"All projects and Instances\" json")
			return allProjectsANDinstances{}, err
		}
		pai.projects = p
		pai.instances = i
		h = append(h, pai)
	}
	return h, err
}

func getAllProjectsUsage(authRes authResponse) (allTenantsUsage, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	h := allTenantsUsage{}
	novaEnd, err := authRes.getEndpoint(compute)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Nova endpoint didn't find")
		return allTenantsUsage{}, err
	}
	projects, err := getProjects(authRes)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Can't get project list")
		return allTenantsUsage{}, err
	}

	req, err := http.NewRequest(
		"GET", novaEnd+"/os-simple-tenant-usage?detailed=1", nil,
	)
	req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return allTenantsUsage{}, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&h); err != nil {
		fmt.Println(err)
		log.Errorln("Cannot decode Nova projects usage json")
	}
	for i, tu := range h.TenantUsages {
		for _, t := range projects.Tenants {
			if strings.TrimSpace(tu.TenantID) == strings.TrimSpace(t.ID) {
				h.TenantUsages[i].name = t.Name
			}
		}
	}
	return h, err
}
