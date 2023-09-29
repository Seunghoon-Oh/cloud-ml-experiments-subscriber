package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-experiments-subscriber/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var cb *circuit.Breaker
var httpClient *circuit.HTTPClient

func SetupExpCircuitBreaker() {
	httpClient, cb = network.GetHttpClient()
}

func CreateExp() {
	if cb.Ready() {
		resp, err := httpClient.Post("http://cloud-ml-experiments-manager.cloud-ml-experiments:8082/exp", "", nil)
		if err != nil {
			fmt.Println(err)
			cb.Fail()
			return
		}
		cb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		fmt.Println(rsData.Data)
		return
	}
}
