package main

// import (
// 	"fmt"

// 	"github.com/sunzeyong/go-best-practices/chain_of_responsibility"
// )

// func main() {
// 	TestChainResp()

// }

// func TestChainResp() {
// 	patientHealthHandler := chain_of_responsibility.Start{}
// 	p := &chain_of_responsibility.Patient{
// 		Name: "123",
// 	}

// 	patientHealthHandler.SetNext(&chain_of_responsibility.Reception{}).
// 		SetNext(&chain_of_responsibility.Clinic{}).
// 		SetNext(&chain_of_responsibility.Reception{}).
// 		SetNext(&chain_of_responsibility.Pharmacy{})

// 	if err := patientHealthHandler.Execute(p); err != nil {
// 		fmt.Println("fail", err)
// 		return
// 	}
// 	fmt.Println("success")
// }
