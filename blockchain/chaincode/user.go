// 사용자 관련 기능(추가, 조회, 관리 등)
package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type UserChaincode struct {
}

// 사용자 정보 구조
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Badge string `json:"badge"`
	Count int    `json:"count"`
}

func (t *UserChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *UserChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "addUser" {
		return t.addUser(stub, args)
	} else if function == "getUser" {
		return t.getUser(stub, args)
	} else if function == "updateUser" {
		return t.updateUser(stub, args)
	}

	return shim.Error("Invalid function name. Expecting 'addUser', 'getUser', or 'updateUser'")
}

// 사용자 추가
func (t *UserChaincode) addUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {  // 인자의 수를 4로 수정합니다.
		return shim.Error("Incorrect number of arguments. Expecting 4: ID, Name, Badge, Count")
	}

	id := args[0]
	name := args[1]
	badge := args[2]
	count, err := strconv.Atoi(args[3])  // Count를 정수로 변환합니다.
	if err != nil {
		return shim.Error("Invalid Count. It should be an integer.")
	}

	user := &User{
		ID:    id,
		Name:  name,
		Badge: badge,
		Count: count,
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return shim.Error("Error creating user JSON: " + err.Error())
	}

	err = stub.PutState(id, userJSON)
	if err != nil {
		return shim.Error("Failed to add user: " + err.Error())
	}

	return shim.Success(nil)
}
// 사용자 정보 가져오기
func (t *UserChaincode) getUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: ID")
	}

	id := args[0]

	userJSON, err := stub.GetState(id)
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	}
	if userJSON == nil {
		return shim.Error("User not found")
	}

	return shim.Success(userJSON)
}

// 사용자 정보 갱신
func (t *UserChaincode) updateUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {  // 인자의 수를 4로 수정합니다.
		return shim.Error("Incorrect number of arguments. Expecting 4: ID, Name, Badge, Count")
	}

	id := args[0]
	name := args[1]
	badge := args[2]
	count, err := strconv.Atoi(args[3])  // Count를 정수로 변환합니다.
	if err != nil {
		return shim.Error("Invalid Count. It should be an integer.")
	}

	user := &User{
		ID:    id,
		Name:  name,
		Badge: badge,
		Count: count,
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return shim.Error("Error creating user JSON: " + err.Error())
	}

	err = stub.PutState(id, userJSON)
	if err != nil {
		return shim.Error("Failed to update user: " + err.Error())
	}

	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(UserChaincode)); err != nil {
		fmt.Printf("Error starting UserChaincode: %s", err)
	}
}

