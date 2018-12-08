package main

import "fmt"

type ProxyType struct {
}


func NewProxy() ProxyType {
	return ProxyType{}
}

func Proxy1(service string) {
	proxy := map[string]func(){
		"earnings": func() {
			fmt.Println("earnings")
		},
		"approved": func() {
			fmt.Println("approved")
		},
		"rejected": func() {
			fmt.Println("approved")
		},
	}

	proxy[service]()
}
