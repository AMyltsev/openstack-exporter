package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

func getGlanceImages(authRes authResponse) (images, error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	h := images{}
	netEnd, err := authRes.getEndpoint(image)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Glance endpoint didn't find")
		return images{}, err
	}
	req, err := http.NewRequest(
		"GET", netEnd+"/v2/images", nil,
	)
	req.Header.Add("X-Auth-Token", authRes.Access.Token.ID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return images{}, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&h); err != nil {
		fmt.Println(err)
		log.Errorln("Cannot decode Glance images json")
	}

	return h, err
}
