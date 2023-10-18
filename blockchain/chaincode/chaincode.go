package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"time"
)

// 배지 구조 예시
type Badge struct {
	ID          string    `json:"id"`
	Issuer      string    `json:"issuer"`
	Description string    `json:"description"`
	IssuedAt    time.Time `json:"issuedAt"`
	Signature   string    `json:"signature"`
}

// 사용자 정보 구조 예시
type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Badges   []Badge `json:"badges"`
}


type BadgeChaincode struct {
}


func (t *BadgeChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("BadgeChaincode has been initialized")
	return shim.Success(nil)
}


func (t *BadgeChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoking function: " + function)

	switch function {
	case "AddUser":
		return t.AddUser(stub, args)
	case "UserList":
		return t.UserList(stub)
	case "UpdateUserInfo":
		return t.UpdateUserInfo(stub, args)
	case "RegisterBadge":
		return t.RegisterBadge(stub, args)
	case "IssueBadge":
		return t.IssueBadge(stub, args)
	default:
		return shim.Error("Invalid function name")
	}
}

// 사용자 등록
func (t *BadgeChaincode) AddUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: ID and Name")
	}

	userID := args[0]
	userName := args[1]

	// 중복 등록 확인
	userAsBytes, err := stub.GetState(userID)
	if err != nil {
		return shim.Error("Failed to check if the user already exists: " + err.Error())
	}
	if userAsBytes != nil {
		return shim.Error("User with ID " + userID + " already exists")
	}

	// 사용자 정보를 상태 데이터베이스에 저장
	user := User{
		ID:     userID,
		Name:   userName,
		Badges: []Badge{}, // 새 사용자는 배지를 보유하지 않음
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 사용자 정보 저장
	err = stub.PutState(userID, userJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}


// 사용자 목록 확인
func (t *BadgeChaincode) UserList(stub shim.ChaincodeStubInterface) pb.Response {
	// 모든 사용자의 목록을 저장할 슬라이스 생성
	var userList []User

	// 전체 키 범위로부터 사용자 정보를 가져와서 슬라이스에 추가
	resultsIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return shim.Error("Failed to retrieve user list: " + err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Failed to iterate over user list: " + err.Error())
		}

		var user User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			return shim.Error("Failed to unmarshal user data: " + err.Error())
		}

		// 배지 정보 등을 필요에 따라 가공할 수 있습니다.
		// 여기에서는 배지 정보를 제외하고 사용자 정보만 포함시킵니다.
		user.Badges = nil

		// 슬라이스에 사용자 정보 추가
		userList = append(userList, user)
	}

	// 사용자 목록을 JSON 형태로 변환하여 반환
	userListJSON, err := json.Marshal(userList)
	if err != nil {
		return shim.Error("Failed to marshal user list to JSON: " + err.Error())
	}

	return shim.Success(userListJSON)
}


// 사용자 정보 갱신
func (t *BadgeChaincode) UpdateUserInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: ID and New Name")
	}

	userID := args[0]
	newName := args[1]

	// 기존 사용자 정보를 조회
	userAsBytes, err := stub.GetState(userID)
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	}
	if userAsBytes == nil {
		return shim.Error("User with ID " + userID + " does not exist")
	}

	var user User
	err = json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return shim.Error("Failed to unmarshal user data: " + err.Error())
	}

	// 사용자 정보 갱신
	user.Name = newName

	// 갱신된 사용자 정보를 상태 데이터베이스에 저장
	userJSON, err := json.Marshal(user)
	if err != nil {
		return shim.Error("Failed to marshal user data: " + err.Error())
	}
	err = stub.PutState(userID, userJSON)
	if err != nil {
		return shim.Error("Failed to update user: " + err.Error())
	}

	return shim.Success(nil)
}


// 배지 등록
func (t *BadgeChaincode) RegisterBadge(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 배지 등록 로직
	// ...
	return shim.Success(nil)
}

// 배지 발급
func (t *BadgeChaincode) IssueBadge(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 배지 발급 로직
	// ...
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(BadgeChaincode))
	if err != nil {
		fmt.Printf("Error starting BadgeChaincode: %s", err)
	}
}
