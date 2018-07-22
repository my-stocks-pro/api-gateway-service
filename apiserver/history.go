package apiserver

import (
	"os"
	"fmt"
	"net/http"
)

//const (
//	approvedHistoryURL = "http://127.0.0.1:8006/approved/history"
//	earningsHistoryURL = "http://127.0.0.1:8007/earnings/history"
//)


func (s *TypeServer)Request (url string) (*http.Response, error) {
	req, errReq := http.NewRequest("GET", url, nil)
	if errReq != nil {
		return nil, errReq
	}

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	return resp, nil
}

func (s *TypeServer) HistoryALL() {
	history := os.Getenv("HISTORYALL")

	if history == "1" {

		fmt.Println("Get ALL history data")

		func(url string) {
			resp, err := s.Request(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(resp)
		}(fmt.Sprintf("http://%s:%s/approved/history",
			s.Config.Hosts["approved-history-service"],
			s.Config.Ports["approved-history-service"]))

		func(url string) {
			resp, err := s.Request(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(resp)
		}(fmt.Sprintf("http://%s:%s/approved/history",
			s.Config.Hosts["earnings-history-service"],
			s.Config.Ports["earnings-history-service"]))

	}
}

func (s *TypeServer) HistoryByMonth() {
	history := os.Getenv("HISTORYBYMONTH")

	if history == "1" {

	}
}
