// 파일 경로 수정필요, 패키징 필요
package main

import (
	"fmt"
	// user.go 파일의 경로수정필요 "github.com/your-package-path/user"
	// badge.go 파일의 경로수정필요 "github.com/your-package-path/badge"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	err := shim.Start(&user.UserChaincode{}, &badge.BadgeContract{})
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
