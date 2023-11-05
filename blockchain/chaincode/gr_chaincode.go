package main

import ( //경로 수정필요
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Bogo struct {
	Price    string `json:"price"`
	Date     string `json:"date"`
	BookName string `json:"bookname"`
	WalletID string `json:"walletid"`
}

type BogoKey struct {
	Key string
	Idx int
}

type Wallet struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Mileage string `json:"mileage"`
}

type Badge struct {
	Owner       string `json:"owner"`
	Description string `json:"description"`
	Mileage     string `json:"mileage"`
}

type SmartContract struct {
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func generateKey(stub shim.ChaincodeStubInterface) []byte {
	var isFirst = false
	bogokeyAsBytes, err := stub.GetState("latestKey")
	if err != nil {
		fmt.Println(err.Error())
	}
	bogokey := BogoKey{}
	json.Unmarshal(bogokeyAsBytes, &bogokey)
	var tempIdx string
	tempIdx = strconv.Itoa(bogokey.Idx)
	fmt.Println(bogokey)
	fmt.Println("Key is " + strconv.Itoa(len(bogokey.Key)))
	if len(bogokey.Key) == 0 || bogokey.Key == "" {
		isFirst = true
		bogokey.Key = "BG"
	}
	if !isFirst {
		bogokey.Idx = bogokey.Idx + 1
	}
	fmt.Println("Last BogoKey is " + bogokey.Key + " : " + tempIdx)
	returnValueBytes, _ := json.Marshal(bogokey)

	return returnValueBytes
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()

	if function == "initWallet" {
		return s.initWallet(APIstub)
	} else if function == "getWallet" {
		return s.getWallet(APIstub, args)
	} else if function == "setBogo" {
		return s.setBogo(APIstub, args)
	} else if function == "getAllBogo" {
		return s.getAllBogo(APIstub)
	} else if function == "purchaseBogoAndIssueBadge" {
		return s.purchaseBogoAndIssueBadge(APIstub, args)
	}

	fmt.Println("Please check your function: " + function)
	return shim.Error("Unknown function")
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (s *SmartContract) initWallet(APIstub shim.ChaincodeStubInterface) pb.Response {
	seller := Wallet{Name: "Student", ID: "1Q2W3E4R", Mileage: "100"}
	customer := Wallet{Name: "Library", ID: "5T6Y7UBI", Mileage: "200"}

	SellerasJSONBytes, _ := json.Marshal(seller)
	err := APIstub.PutState(seller.ID, SellerasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + seller.Name)
	}

	CustomerasJSONBytes, _ := json.Marshal(customer)
	err = APIstub.PutState(customer.ID, CustomerasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + customer.Name)
	}

	return shim.Success(nil)
}

func (s *SmartContract) getWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	walletAsBytes, err := APIstub.GetState(args[0])

	if err != nil {
		fmt.Println(err.Error())
		return shim.Error(err.Error())
	}

	wallet := Wallet{}
	json.Unmarshal(walletAsBytes, &wallet)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	buffer.WriteString("{\"Name\":\"")
	buffer.WriteString(wallet.Name)
	buffer.WriteString("\",\"ID\":\"")
	buffer.WriteString(wallet.ID)
	buffer.WriteString("\",\"Mileage\":\"")
	buffer.WriteString(wallet.Mileage)
	buffer.WriteString("\"}")

	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) setBogo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	var bogokey = BogoKey{}
	json.Unmarshal(generateKey(APIstub), &bogokey)
	keyidx := strconv.Itoa(bogokey.Idx)
	fmt.Println("Key : " + bogokey.Key + ", Idx : " + keyidx)

	var bogo = Bogo{Price: args[0], Date: args[1], BookName: args[2], WalletID: args[3]}

	bogoAsJSONBytes, err := json.Marshal(bogo)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal bogo: %s", err.Error()))
	}

	var keyString = bogokey.Key + keyidx

	fmt.Println("bogokey is " + keyString)
	err = APIstub.PutState(keyString, bogoAsJSONBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record bogo: %s", err.Error()))
	}

	bogokeyAsBytes, err := json.Marshal(bogokey)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal bogokey: %s", err.Error()))
	}

	err = APIstub.PutState("latestKey", bogokeyAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update latestKey: %s", err.Error()))
	}

	return shim.Success(nil)
}
func (s *SmartContract) getAllBogo(APIstub shim.ChaincodeStubInterface) pb.Response {
	bogokeyAsBytes, _ := APIstub.GetState("latestKey")
	bogokey := BogoKey{}
	json.Unmarshal(bogokeyAsBytes, &bogokey)
	idxStr := strconv.Itoa(bogokey.Idx + 1)

	var startKey = "BG0"
	var endKey = bogokey.Key + idxStr

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	for resultsIter.HasNext() {
		queryResponse, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true

	}
	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) purchaseBogoAndIssueBadge(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// 1. Bogo 정보 가져오기
	bogoAsBytes, err := APIstub.GetState(args[2])
	if err != nil {
		return shim.Error(err.Error())
	}
	bogo := Bogo{}
	err = json.Unmarshal(bogoAsBytes, &bogo)
	if err != nil {
		return shim.Error("Failed to unmarshal Bogo data")
	}

	// 2. 판매자(seller)와 구매자(customer)의 지갑 정보 가져오기
	sellerAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get seller's wallet")
	}
	customerAsBytes, err := APIstub.GetState(args[1])
	if err != nil {
		return shim.Error("Failed to get customer's wallet")
	}

	// 3. 지갑 정보를 구조체로 언마샬
	sellerWallet := Wallet{}
	customerWallet := Wallet{}
	err = json.Unmarshal(sellerAsBytes, &sellerWallet)
	if err != nil {
		return shim.Error("Failed to unmarshal seller's wallet data")
	}
	err = json.Unmarshal(customerAsBytes, &customerWallet)
	if err != nil {
		return shim.Error("Failed to unmarshal customer's wallet data")
	}

	// 4. Bogo 가격을 정수로 변환
	bogoPrice, err := strconv.Atoi(bogo.Price)
	if err != nil {
		return shim.Error("Failed to convert Bogo price to integer")
	}

	// 5. 구매자의 마일리지 확인
	customerMileage, err := strconv.Atoi(customerWallet.Mileage)
	if err != nil {
		return shim.Error("Failed to convert customer's mileage to integer")
	}

	// 6. Bogo 구매에 필요한 마일리지 확인
	if customerMileage < bogoPrice {
		return shim.Error("Insufficient mileage to purchase Bogo")
	}

	// 7. 판매자와 구매자의 마일리지 업데이트
	sellerMileage, err := strconv.Atoi(sellerWallet.Mileage)
	if err != nil {
		return shim.Error("Failed to convert seller's mileage to integer")
	}

	sellerMileage += bogoPrice
	customerMileage -= bogoPrice

	sellerWallet.Mileage = strconv.Itoa(sellerMileage)
	customerWallet.Mileage = strconv.Itoa(customerMileage)

	// 8. 업데이트된 지갑 정보 저장
	sellerWalletAsBytes, err := json.Marshal(sellerWallet)
	if err != nil {
		return shim.Error("Failed to marshal seller's wallet data")
	}

	customerWalletAsBytes, err := json.Marshal(customerWallet)
	if err != nil {
		return shim.Error("Failed to marshal customer's wallet data")
	}

	err = APIstub.PutState(args[0], sellerWalletAsBytes)
	if err != nil {
		return shim.Error("Failed to update seller's wallet")
	}

	err = APIstub.PutState(args[1], customerWalletAsBytes)
	if err != nil {
		return shim.Error("Failed to update customer's wallet")
	}

	// 9. 판매자에게 Badge 발급 조건 확인
	totalMileage := sellerMileage + customerMileage
	if totalMileage%5 == 0 {
		// 10. Badge 발급
		badge := Badge{
			Owner:       args[0], // 판매자의 ID
			Description: "Achieved 5x total mileage",
			Mileage:     strconv.Itoa(totalMileage),
		}
		badgeAsBytes, err := json.Marshal(badge)
		if err != nil {
			return shim.Error("Failed to marshal badge data")
		}

		badgeKey := "Badge_" + args[0]
		err = APIstub.PutState(badgeKey, badgeAsBytes)
		if err != nil {
			return shim.Error("Failed to issue badge")
		}
	}

	return shim.Success(nil)
}
