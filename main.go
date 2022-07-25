package main

import "github.com/Qurb0n0vSh2oo2/wallet/pkg/wallet"

func main(){
	svc := &wallet.Service{}
	wallet.RegisterAccount(svc, "+992918703330")
}