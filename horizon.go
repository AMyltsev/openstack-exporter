package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/prometheus/common/log"
)

func loginHorizon(conf config) error {
	horizonURL := conf.HorizonURL + "horizon/auth/login/"
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest(
		"GET", horizonURL, nil,
	)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Can't make GET request to Horizon")
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Can't do GET request to Horizon")
		return err
	}
	defer res.Body.Close()
	u, _ := url.Parse(horizonURL)
	csrfToken := getCookieValue(client.Jar.Cookies(u), "csrftoken")
	str := "csrfmiddlewaretoken=" + csrfToken + "&username=" + conf.Openstack.Username + "&password=" + conf.Openstack.Password
	r := strings.NewReader(str)
	req, err = http.NewRequest(
		"POST", horizonURL, r,
	)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Can't make POST request to Horizon")
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", horizonURL)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Errorln("Can't do POST request to Horizon")
		return err
	}

	sessionID := getCookieValue(res.Cookies(), "sessionid")
	if sessionID == "" {
		err := errors.New("There is not session ID")
		return err
	}
	return err
}

func getCookieValue(coockies []*http.Cookie, str string) string {
	for _, v := range coockies {
		if v.Name == str {
			return v.Value
		}
	}
	return ""
}
