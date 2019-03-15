package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

func osAuth(conf config) (authResponse, error) {
	var auth authOpt
	var authRes authResponse
	auth.Auth.TenantName = conf.Openstack.Username
	auth.Auth.PasswordCredentials.Username = conf.Openstack.Username
	auth.Auth.PasswordCredentials.Password = conf.Openstack.Password
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(auth)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Cannot encode auth json")
	}
	res, err := http.Post(conf.AuthURL+"tokens", "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Println(err)
		return authResponse{}, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&authRes); err != nil {
		fmt.Println(err)
		log.Fatalln("Cannot decode auth json")
	}
	if authRes.Access.Token.ID == "" {
		err := errors.New("Auth does not have token")
		return authRes, err
	}
	return authRes, err
}
