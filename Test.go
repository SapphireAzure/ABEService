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
	ABES.Encrypt("33","qwer",ap)
	ABES.Query("policy")


	ABES.Close()
}
