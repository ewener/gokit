/******************************************************
# DESC       : ip
# MAINTAINER : yamei
# EMAIL      : daixw@ecpark.cn
# DATE       : 2020/2/24
******************************************************/
package backend

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"fmt"
	"strings"
)

type IPInfo struct {
	Code int `json:"code"`
	Data IP  `json:"data"`
}

type IP struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
}

func init() {
	for _, b := range []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"} {
		if _, block, err := net.ParseCIDR(b); err == nil {
			privateBlocks = append(privateBlocks, block)
		}
	}
}

var (
	privateBlocks []*net.IPNet
)

func isPrivateIP(ip net.IP) bool {
	for _, priv := range privateBlocks {
		if priv.Contains(ip) {
			return true
		}
	}
	return false
}

// 获取公网ip
func GetExternalIp() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

// 淘宝api查询公网地址信息
func TabaoAPI(ip string) *IPInfo {
	url := "http://ip.taobao.com/service/getIpInfo.php?ip="
	url += ip

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var result IPInfo
	if err := json.Unmarshal(out, &result); err != nil {
		return nil
	}

	return &result
}

// 判断是否是公网ip
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		// 排除私有网段
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// 通过dns服务器8.8.8.8:80获取使用的ip
func GetPublicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

func GetLocalIP() (string, error) {
	faces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	var addr net.IP
	for _, face := range faces {
		if !isValidNetworkInterface(face) {
			continue
		}

		addrs, err := face.Addrs()
		if err != nil {
			return "", err
		}

		if ipv4, ok := getValidIPv4(addrs); ok {
			addr = ipv4
			if isPrivateIP(ipv4) {
				return ipv4.String(), nil
			}
		}
	}

	if addr == nil {
		return "", errors.New("cannot get local IP")
	}

	return addr.String(), nil
}

func isValidNetworkInterface(face net.Interface) bool {
	if face.Flags&net.FlagUp == 0 {
		// interface down
		return false
	}

	if face.Flags&net.FlagLoopback != 0 {
		// loopback interface
		return false
	}

	if strings.Contains(strings.ToLower(face.Name), "docker") {
		return false
	}

	return true
}

func getValidIPv4(addrs []net.Addr) (net.IP, bool) {
	for _, addr := range addrs {
		var ip net.IP

		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		ip = ip.To4()
		if ip == nil {
			// not an valid ipv4 address
			continue
		}

		return ip, true
	}
	return nil, false
}

// 获取外部ip
// 1.优先外网ip
// 2.其次局域网ip
// 3.1、2都没获取到则返回错误
func GetIp() (string, []string, error) {
	addrs, err := net.InterfaceAddrs()
	localIps := make([]string, 0)
	pubicIp := ""
	if err != nil {
		return pubicIp, localIps, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			ip := ipnet.IP
			if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
				continue
			}
			// 找到外网ip直接返回
			if IsPublicIP(ip) && pubicIp != "" {
				pubicIp = ip.String()
			}
			localIps = append(localIps, ip.String())
		}
	}
	if len(localIps) == 0 {
		return pubicIp, localIps, errors.New(fmt.Sprintf("connected to the network?"))
	}
	return pubicIp, localIps, nil
}
