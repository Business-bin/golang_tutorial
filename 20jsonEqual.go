package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func main() {
	var aList, bList interface{}
	aMapStr := `[{"channel_2":"","channel_5":"128","channel_6":"","idx":0,"ip_address":"0.0.0.0","mac_address":"98:77:e7:03:1a:c4","mac_address_2":"","mac_address_5":"","mac_address_6":"","model_name":"AR1344","name":"AR1344","parents_mac_address":"","rssi":"-10","station":[{"ip_address":"","mac_address":"12:23:30:82:D2:88","name":"12:23:30:82:D2:88","rssi":"-50","upper_connecting_frequency_band":"5GHz"}],"type":"Slave","upper_access_point":"16","upper_connecting_frequency_band":"5GHz" },{ "channel_2": "" ,"channel_5":"128","channel_6":"","idx":1,"ip_address":"192.168.0.21","mac_address":"98:77:e7:03:1a:4c","mac_address_2":"","mac_address_5":"98:77:E7:03:1A:4F","mac_address_6":"","model_name":"AR1344","name":"AR1344","parents_mac_address":"","rssi":"-50","station":[{"ip_address":"","mac_address":"12:23:30:82:d2:88","name":"12:23:30:82:d2:88","rssi":"-50","upper_connecting_frequency_band":"5GHz"}],"type":"Master","upper_access_point":"","upper_connecting_frequency_band":""}]`
	bMapStr := `[{"channel_5":"128","channel_6":"","idx":1,"ip_address":"192.168.0.21","mac_address":"98:77:e7:03:1a:4c","mac_address_2":"","mac_address_5":"98:77:E7:03:1A:4F","mac_address_6":"","model_name":"AR1344","name":"AR1344","parents_mac_address":"","rssi":"-50","station":[{"ip_address":"","mac_address":"12:23:30:82:d2:88","name":"12:23:30:82:d2:88","rssi":"-50","upper_connecting_frequency_band":"5GHz"}],"type":"Master","upper_access_point":"","upper_connecting_frequency_band":"","channel_2":""},{"channel_2":"","channel_5":"128","channel_6":"","idx":0,"ip_address":"0.0.0.0","mac_address":"98:77:e7:03:1a:c4","mac_address_2":"","mac_address_5":"","mac_address_6":"","model_name":"AR1344","name":"AR1344","parents_mac_address":"","rssi":"-10","station":[{"ip_address":"","mac_address":"12:23:30:82:D2:88","name":"12:23:30:82:D2:88","rssi":"-50","upper_connecting_frequency_band":"5GHz"}],"type":"Slave","upper_access_point":"16","upper_connecting_frequency_band":"5GHz"}]`

	//aMapStr := `{"channel_2":"","channel_6":"","idx":1,"ip_address":"192.168.0.21","mac_address_2":"","mac_address_5":"98:77:E7:03:1A:4F","mac_address_6":"","model_name":"AR1344","name":"AR1344","rssi":"-50","stations":[{"ip_address":"","mac_address":"56:38:f4:49:61:30",	"name":"56:38:f4:49:61:30","rssi":"-50","upper_connecting_frequency_band":"5GHz"}],"type":"Master","upper_access_point":"","upper_connecting_frequency_band":"","channel_5":"128"}`
	//bMapStr := `{"channel_2":"","channel_5":"128","channel_6":"","idx":1,"ip_address":"192.168.0.21","mac_address_2":"","mac_address_5":"98:77:E7:03:1A:4F","mac_address_6":"","model_name":"AR1344","name":"AR1344","rssi":"-50","stations":[{"ip_address":"","name":"56:38:f4:49:61:30","rssi":"-50","mac_address":"56:38:f4:49:61:30","upper_connecting_frequency_band":"5GHz"}],"type":"Master","upper_access_point":"","upper_connecting_frequency_band":""}`

	//fmt.Println("111111111111 aList = ", aList)
	//fmt.Println("111111111111 bList = ", bList)
	json.Unmarshal([]byte(aMapStr), &aList)
	json.Unmarshal([]byte(bMapStr), &bList)
	fmt.Println("222222222222 aList = ", aList)
	fmt.Println("222222222222 bList = ", bList)

	// 배열 안의 map의 순서 바뀌면 false 리턴
	//eq, err := JSONByEqual([]byte(aMapStr), []byte(bMapStr))
	//fmt.Println("333333333333 aMapStr=bMapStr\t", eq, "with error", err)
	//apEq := reflect.DeepEqual(aList, bList)
	//fmt.Println("4444444444444 apEq = ", apEq)

	// 배열 안의 map의 순서 달라도 비교
	startTime1 := time.Now()
	apEq2 := JSONEqual(aList, bList)
	elapsedTime1 := time.Since(startTime1)
	fmt.Println("5555555555555 실행시간 = ", elapsedTime1.Seconds(), " apEq2 = ", apEq2)

	startTime2 := time.Now()
	apEq3 := reflect.DeepEqual(aList, bList)
	elapsedTime2 := time.Since(startTime2)
	fmt.Println("6666666666666 실행시간 = ", elapsedTime2.Seconds(), " apEq3 = ", apEq3)
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONByEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// Equal checks equality between 2 Body-encoded data.
func JSONEqual(vx, vy interface{}) bool {
	if reflect.TypeOf(vx) != reflect.TypeOf(vy) {
		return false
	}

	switch x := vx.(type) {
	case map[string]interface{}:
		y := vy.(map[string]interface{})

		if len(x) != len(y) {
			return false
		}

		for k, v := range x {
			val2 := y[k]

			if (v == nil) != (val2 == nil) {
				return false
			}

			if !JSONEqual(v, val2) {
				return false
			}
		}

		return true
	case []interface{}:
		y := vy.([]interface{})

		if len(x) != len(y) {
			return false
		}

		var matches int
		flagged := make([]bool, len(y))
		for _, v := range x {
			for i, v2 := range y {
				if JSONEqual(v, v2) && !flagged[i] {
					matches++
					flagged[i] = true

					break
				}
			}
		}

		return matches == len(x)
	default:
		return vx == vy
	}
}
