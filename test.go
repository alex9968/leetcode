package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

type SelfValidateFunc func() bool

var  funcList = []SelfValidateFunc {
	fff,
	fff2,

}


func main() {
	a := []string{
		"/yxct/s/managementHomepage/initGraphicalReport.json",
		"/yxct-portal/s/investFund/findInvestFundList.json?_v=1578381339028",
		"/yxct-portal/s/investFund/findInvestFundList.jpg?_v=1578381339028",
		"/yxct-portal/s/investFund/findInvestFundList.png",
		"/yxct-portal/s/investFund/findInvestFundList",
	}

	for _, v := range a {
		fmt.Println(v, filterUri(v))
	}
}


	func filterUri(uri string) (bool){
		suffixList := []string{".jpg", ".pdf", ".png", ".css", ".less", ".sass", ".scss",".gif", ".svg", ".woff"}
		ext := filepath.Ext(uri)
		for _,v := range suffixList {
			if strings.Contains(ext, v) {
				return false
			}
		}
		return true //不包含静态资源
	}

func fff() bool {
	return true
}
func fff2()  bool{
	return true
}


