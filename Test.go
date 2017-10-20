package main

import (
	"./src/ABEservice"
)

var ABES ABEservice.ABEService

func init(){

}


func main(){
	//现在封装成ABEService
	ABES = ABEservice.ServiceInit("PolicyDB")
	ABES.Query("policy")
	ap := []string{"abc","123","ABE"}
	ABES.PropertyUpdate("33",ap)
	ABES.Query("policy")
	//数据库关闭
	ABES.Close()
}
