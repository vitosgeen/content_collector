package smartproxy

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const lenProxySplit = 4

type ISmartProxyFile interface {
	GetProxy(index int) (string, error)
	GetProxyRandom() (string, error)
	GetProxyRandomSmartProxy() (*SmartProxy, error)
	ParseFile() ([]*SmartProxy, error)
}

type SmartProxy struct {
	Username string
	Host     string
	Password string
	Port     string
}

type SmartProxyFileTxt struct {
	pathTextFile string
	proxies      []*SmartProxy
}

func NewSmartProxy(pathTextFile string) ISmartProxyFile {
	return &SmartProxyFileTxt{
		pathTextFile: pathTextFile,
	}
}

func (s *SmartProxyFileTxt) GetProxy(index int) (string, error) {
	if index >= len(s.proxies) || index < 0 {
		return "", fmt.Errorf("index out of range")
	}

	proxy := s.proxies[index]
	proxyString := fmt.Sprintf("%s:%s@%s:%s", proxy.Username, proxy.Password, proxy.Host, proxy.Port)
	return proxyString, nil
}

func (s *SmartProxyFileTxt) GetProxyRandom() (string, error) {
	lenProxies := len(s.proxies)
	if lenProxies == 0 {
		return "", fmt.Errorf("no proxies available")
	}

	randomIndex := rand.Intn(lenProxies - 1)
	randomProxy := s.proxies[randomIndex]
	proxyString := fmt.Sprintf("%s:%s@%s:%s", randomProxy.Username, randomProxy.Password, randomProxy.Host, randomProxy.Port)

	return proxyString, nil
}

func (s *SmartProxyFileTxt) GetProxyRandomSmartProxy() (*SmartProxy, error) {
	lenProxies := len(s.proxies)
	if lenProxies == 0 {
		return nil, fmt.Errorf("no smart proxies available")
	}

	randomIndex := rand.Intn(lenProxies - 1)
	randomProxy := s.proxies[randomIndex]

	return randomProxy, nil
}

func (s *SmartProxyFileTxt) ParseFile() ([]*SmartProxy, error) {
	proxies := []*SmartProxy{}
	file, err := os.Open(s.pathTextFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		smartProxy := &SmartProxy{}
		line := scanner.Text()
		split := strings.FieldsFunc(line, split)
		if len(split) != lenProxySplit {
			return nil, fmt.Errorf("invalid proxy format")
		}
		smartProxy.Username = split[0]
		smartProxy.Password = split[1]
		smartProxy.Host = split[2]
		smartProxy.Port = split[3]
		proxies = append(proxies, smartProxy)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	s.proxies = proxies
	fmt.Println("Proxies loaded", len(s.proxies), proxies[0])
	return proxies, nil
}

func split(r rune) bool {
	return r == ':' || r == '@'
}

func (sp *SmartProxy) String() string {
	return fmt.Sprintf("%s:%s@%s:%s", sp.Username, sp.Password, sp.Host, sp.Port)
}
