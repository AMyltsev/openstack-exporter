package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

func getProjects(authRes authResponse) (projects, error) {
	var t projects
	client := &http.Client{}
	defer client.CloseIdleConnections()
	keyEnd, err := authRes.getEndpoint(identity)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Keystone endpoint didn't find")
		return projects{}, err
	}
	req, err := http.NewRequest(
		"GET", keyEnd+"/tenants", nil,
	)
	req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return projects{}, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Println(err)
		log.Errorln("Cannot decode Projects json")
	}
	return t, err
}
