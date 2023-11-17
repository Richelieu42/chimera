package gaodeKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"testing"
)

func TestClient_GetIp(t *testing.T) {
	client, err := NewClient("b15c36bf1df4c272e92f3f1875a127f1")
	if err != nil {
		panic(err)
	}

	//{
	//	ip := "10.0.9.141"
	//	ipInfo, err := client.GetIpInfo(ip)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(jsonKit.MarshalIndentToString(ipInfo, "", "    "))
	//}
	//{
	//	ip := "218.90.174.146"
	//	ipInfo, err := client.GetIpInfo(ip)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(jsonKit.MarshalIndentToString(ipInfo, "", "    "))
	//}
	{
		ip := "1.0.16.0"
		ipInfo, err := client.GetIpInfo(ip)
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonKit.MarshalIndentToString(ipInfo, "", "    "))
	}

	/*
		{
		    "province": "局域网",
		    "city": "",
		    "adcode": "",
		    "rectangle": ""
		} <nil>
		{
		    "province": "江苏省",
		    "city": "无锡市",
		    "adcode": "320200",
		    "rectangle": "120.1788533,31.4648817;120.4605818,31.68307651"
		} <nil>
	*/
}
