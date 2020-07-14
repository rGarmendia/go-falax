package main

import (
	"fmt"

	"github.com/kr/pretty"
)

func main() {

	subnets := []string{"10.145.,31.64/26",
		"10.145.12.80/28",
		"10.145.15.0/27",
		"10.145.14.128/27",
		"10.145.64.240/28",
		"10.145.64.112/28",
		"10.145.68.112/28",
		"10.145.64.144/28",
		"10.145.117.224/27",
		"10.145.15.128/28",
		"10.145.122.32/27",
		"10.145.123.32/27",
		"10.145.64.208/28",
		"10.145.14.0/27",
		"10.145.122.128/27",
		"10.145.64.16/28",
		"10.145.64.0/28",
		"10.145.12.32/27",
		"10.145.123.128/28",
		"10.145.69.32/28",
		"10.145.14.96/27",
		"10.145.14.64/27",
		"10.145.68.16/28",
		"10.145.65.64/28",
		"10.145.122.160/27",
		"10.145.122.192/27",
		"10.145.12.0/27",
		"10.145.123.0/27",
		"10.145.120.160/27",
		"10.145.13.176/28",
		"10.145.12.64/28",
		"10.145.69.224/27",
		"10.145.15.96/27",
		"10.145.15.160/27",
		"10.145.118.224/28",
		"10.145.70.80/28",
		"10.145.118.32/28",
		"10.145.73.16/28",
		"10.145.118.16/28",
		"10.145.15.32/27",
		"10.145.112.144/28",
		"10.145.112.48/28",
		"10.145.113.16/28",
		"10.145.12.224/28",
		"10.145.12.240/28",
		"10.145.113.208/28",
		"10.145.73.208/28",
		"10.145.118.128/28",
		"10.145.112.176/28",
		"10.145.74.48/28",
		"10.145.15.208/28",
		"10.145.14.208/28",
		"10.145.72.112/28",
		"10.145.75.208/28",
		"10.145.114.80/28",
		"10.145.74.80/28",
		"10.145.121.32/27",
		"10.145.73.64/28",
		"10.145.115.192/28",
		"10.145.123.160/27",
		"10.145.75.192/28",
		"10.145.64.192/28",
		"10.145.12.96/28",
		"10.145.64.224/28",
		"10.145.113.192/28",
		"10.145.118.208/28",
		"10.145.13.32/28",
		"10.145.73.192/28",
		"10.145.65.0/28",
		"10.145.113.224/28",
		"10.145.123.192/27",
		"10.145.13.128/27",
		"10.145.12.128/28",
		"10.145.15.224/28",
		"10.145.72.32/28",
		"10.145.13.192/27",
		"10.145.14.160/27",
		"10.145.113.32/28",
		"10.145.118.192/28",
		"10.145.119.0/27",
		"10.145.124.32/27",
		"10.145.118.64/28",
		"10.145.112.224/28",
		"10.145.13.64/28",
		"10.145.119.32/28",
		"10.145.74.208/28",
		"10.145.124.96/27",
		"10.145.114.128/28",
		"10.145.74.192/28",
		"10.145.112.208/28",
		"10.145.118.112/28",
		"10.145.120.0/27",
		"10.145.117.208/28",
		"10.145.72.176/28",
		"10.145.114.144/28",
		"10.145.122.224/27",
		"10.145.119.64/27",
		"10.145.74.96/28",
		"10.145.113.80/28",
		"10.145.119.96/27",
		"10.145.69.208/28",
		"10.145.119.160/28",
		"10.145.112.64/28",
		"10.145.114.32/28",
		"10.145.119.128/27",
		"10.145.13.16/28",
		"10.145.12.144/28",
		"10.145.75.0/26",
		"10.145.73.80/28",
		"10.145.65.192/28",
		"10.145.65.208/28",
		"10.145.121.96/27",
		"10.145.65.224/27",
		"10.145.119.192/27",
		"10.145.112.192/28",
		"10.145.112.32/28",
		"10.145.113.96/28",
		"10.145.65.112/28",
		"10.145.114.16/28",
		"10.145.112.96/28",
		"10.145.66.0/28",
		"10.145.66.64/26",
		"10.145.14.240/28",
		"10.145.124.64/27",
		"10.145.66.128/28",
		"10.145.114.48/28",
		"10.145.114.64/28",
		"10.145.12.160/28",
		"10.145.113.240/28",
		"10.145.14.192/28",
		"10.145.72.48/28",
		"10.145.73.96/28",
		"10.145.12.176/28",
		"10.145.74.16/28",
		"10.145.73.240/28",
		"10.145.120.192/27",
		"10.145.123.96/27",
		"10.145.123.144/28",
		"10.145.74.0/28",
		"10.145.72.0/28",
		"10.145.65.16/28",
		"10.145.13.0/28",
		"10.145.64.96/28",
		"10.145.67.128/28",
		"10.145.15.240/28",
		"10.145.14.224/28",
		"10.145.124.160/27",
		"10.145.123.224/27",
		"10.145.72.208/28",
		"10.145.65.32/28",
		"10.145.124.224/27",
		"10.145.66.176/28",
		"10.145.66.208/28",
		"10.145.69.64/27",
		"10.145.124.192/27",
		"10.145.67.32/28",
		"10.145.67.48/28",
		"10.145.14.32/27",
		"10.145.67.160/27",
		"10.145.118.144/28",
		"10.145.13.112/28",
		"10.145.65.128/28",
		"10.145.114.160/28",
		"10.145.114.176/28",
		"10.145.67.224/27",
		"10.145.13.96/28",
		"10.145.67.144/28",
		"10.145.74.32/28",
		"10.145.112.80/28",
		"10.145.113.112/28",
		"10.145.74.64/28",
		"10.145.66.224/27",
		"10.145.72.16/28",
		"10.145.116.128/27",
		"10.145.121.160/27",
		"10.145.73.32/28",
		"10.145.12.192/28",
		"10.145.64.128/28",
		"10.145.64.64/28",
		"10.145.121.0/27",
		"10.145.64.48/28",
		"10.145.121.64/27",
		"10.145.120.32/27",
		"10.145.72.128/28",
		"10.145.12.112/28",
		"10.145.115.0/28",
		"10.145.66.32/27",
		"10.145.73.112/28",
		"10.145.118.0/28",
		"10.145.66.16/28",
		"10.145.15.192/28",
		"10.145.64.160/27",
		"10.145.126.0/26",
		"10.145.72.64/28",
		"10.145.75.144/28",
		"10.145.15.144/28",
		"10.145.15.64/27",
		"10.145.75.128/28",
		"10.145.125.0/26",
		"10.145.121.224/27",
		"10.145.67.64/26",
		"10.145.121.128/27",
		"10.145.68.224/27",
		"10.145.123.64/27",
		"10.145.75.160/28",
		"10.145.112.0/28",
		"10.145.73.176/28",
		"10.145.69.0/28",
		"10.145.113.48/28",
		"10.145.72.160/28",
		"10.145.117.0/26",
		"10.145.75.176/28",
		"10.145.69.128/27",
		"10.145.116.192/26",
		"10.145.72.96/28",
		"10.145.68.64/27",
		"10.145.119.224/27",
		"10.145.120.128/27",
		"10.145.75.224/28",
		"10.145.70.192/27",
		"10.145.65.144/28",
		"10.145.70.112/28",
		"10.145.112.240/28",
		"10.145.13.224/27",
		"10.145.116.0/26",
		"10.145.112.112/28",
		"10.145.113.0/28",
		"10.145.68.128/27",
		"10.145.13.160/28",
		"10.145.126.176/28",
		"10.145.115.144/28",
		"10.145.118.176/28",
		"10.145.67.192/27",
		"10.145.75.240/28",
		"10.145.72.80/28",
		"10.145.71.224/27",
		"10.145.120.224/27",
		"10.145.126.128/28",
		"10.145.69.96/27",
		"10.145.68.192/27",
		"10.145.119.48/28",
		"10.145.115.64/28",
		"10.145.74.224/27",
		"10.145.71.0/26",
		"10.145.119.176/28"}

	type myType struct {
		a, b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("%# v", pretty.Formatter(x))

	_ = subnets

	// arr := strings.Split("10.145.68.48/27", "/")
	// host := arr[0]
	// mask, _ := strconv.Atoi(arr[1])
	// subnetToValid := ipsubnet.SubnetCalculator(host, mask)
	// _ = subnetToValid
	// ip, valid, _ := net.ParseCIDR("10.145.68.48/27")

	// fmt.Println(valid)
	// fmt.Println(ip)
	// fmt.Println(valid.IP)
	// fmt.Println(valid.IP.Equal(ip))

	// for _, subnetStr := range subnets {
	// 	_, subnetMap, _ := net.ParseCIDR(subnetStr)

	// 	if subnetMap.Contains(net.ParseIP(subnetToValid.GetIPAddress())) || subnetMap.Contains(net.ParseIP(subnetToValid.GetBroadcastAddress())) {
	// 		return errors.New("This subnet " + subnet + " is Overlaped by " + subnetStr)
	// 	}
	// }

}
