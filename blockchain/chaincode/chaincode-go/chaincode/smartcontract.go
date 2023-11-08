package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Bogo struct {
	BogoNo         int `json:"BogoNo"`
	Title          string `json:"Title"`
	Author         string `json:"Author"`
	Publisher      string `json:"Publisher"`
	Report         string `json:"Report"`
	Owner          string `json:"Owner"` //독후감 작성자 - 이메일(=id 역할)입력되도록?
	Approved       bool   `json:"Approved"` //초기 create시 false
}

// AssetCounter는 BogoNo를 저장하는 데 사용
type AssetCounter struct {
	Counter int `json:"Counter"`
}

const bogoKeyFormat = "bogo%d"

// InitLedger는 BogoNo 카운터를 초기화합니다.
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assetCounter := AssetCounter{
		Counter: 0,
	}

	// AssetCounter 구조체를 JSON으로 직렬화
	assetCounterJSON, err := json.Marshal(assetCounter)
	if err != nil {
		return err
	}

	// BogoNo 카운터를 월드 스테이트에 저장
	err = ctx.GetStub().PutState("BogoNoCounter", assetCounterJSON)
	if err != nil {
		return fmt.Errorf("failed to put BogoNoCounter to world state. %v", err)
	}

	return nil
}

// CreateBogo는 Bogo를 생성하고 BogoNo를 자동으로 증가시킵니다.
func (s *SmartContract) CreateBogo(ctx contractapi.TransactionContextInterface, title string, author string, publisher string, report string, owner string) (int, error) {
	
	//var bogoNo int
	
	// BogoNo 카운터를 읽어오기
	assetCounterJSON, err := ctx.GetStub().GetState("BogoNoCounter")
	if err != nil {
		return 0, fmt.Errorf("failed to read BogoNoCounter from world state. %v", err)
	}

	var assetCounter AssetCounter
	if assetCounterJSON != nil {
		err = json.Unmarshal(assetCounterJSON, &assetCounter)
		if err != nil {
			return 0, fmt.Errorf("failed to unmarshal BogoNoCounter. %v", err)
		}
	}
	
	//실제 BogoNo=1부터 시작
	assetCounter.Counter++
	bogoNo := assetCounter.Counter

	// Bogo 생성 (실제 BogoNo=1부터 시작)
	bogo := Bogo{
		BogoNo:    assetCounter.Counter,
		Title:     title,
		Author:    author,
		Publisher: publisher,
		Report:    report,
		Owner:     owner,
		Approved:  false,
	}

	// BogoNo +1 증가
	assetCounter.Counter++

	// Bogo 구조체를 JSON으로 직렬화
	bogoJSON, err := json.Marshal(bogo)
	if err != nil {
		return 0, err
	}

	// 새로운 Bogo를 월드 스테이트에 저장
	err = ctx.GetStub().PutState(fmt.Sprintf(bogoKeyFormat, bogoNo), bogoJSON)
	if err != nil {
		return 0, fmt.Errorf("failed to put Bogo to world state. %v", err)
	}

	//업데이트된 BogoNo 카운터 정보를 JSON으로 직렬화합니다.
	updatedCounterJSON, err := json.Marshal(assetCounter)
	if err != nil {
		return 0, err
	}

	//업데이트된 BogoNo 카운터를 "BogoNoCounter" 키로 원장(world state)에 저장
	err = ctx.GetStub().PutState("BogoNoCounter", updatedCounterJSON)
	if err != nil {
		return 0, fmt.Errorf("failed to update BogoNoCounter in world state. %v", err)
	}

	return bogoNo, nil
}

// BogoNo를 기준으로 특정 bogo 읽어오기
func (s *SmartContract) ReadBogo(ctx contractapi.TransactionContextInterface, bogoNo int) (*Bogo, error) {
    bogoNoKey := fmt.Sprintf(bogoKeyFormat, bogoNo) // BogoNo를 사용하여 키 생성
    bogoJSON, err := ctx.GetStub().GetState(bogoNoKey)
    if err != nil {
        return nil, fmt.Errorf("failed to read Bogo from world state: %v", err)
    }
    if bogoJSON == nil {
        return nil, fmt.Errorf("the Bogo with BogoNo %d does not exist", bogoNo)
    }

    var bogo Bogo
    err = json.Unmarshal(bogoJSON, &bogo)
    if err != nil {
        return nil, err
    }

    return &bogo, nil
}

// ApproveBogo 함수는 도서관 관리자가 독후감을 승인하는 로직을 처리
// ApproveBogo는 Bogo의 Approved 필드를 true로 업데이트합니다.
func (s *SmartContract) ApproveBogo(ctx contractapi.TransactionContextInterface, bogoNo int) error {
    // BogoNo에 해당하는 Bogo 데이터를 읽어오기
    bogoNoKey := fmt.Sprintf(bogoKeyFormat, bogoNo)
    bogoJSON, err := ctx.GetStub().GetState(bogoNoKey)
    if err != nil {
        return fmt.Errorf("failed to read Bogo from world state: %v", err)
    }
    if bogoJSON == nil {
        return fmt.Errorf("the Bogo with BogoNo %d does not exist", bogoNo)
    }

    // 읽어온 Bogo 데이터를 구조체로 unmarshal
    var bogo Bogo
    err = json.Unmarshal(bogoJSON, &bogo)
    if err != nil {
        return err
    }

    // Approved 필드 false-> true로 설정(Bogo에 대해서 승인)
    bogo.Approved = true

    // 업데이트된 Bogo 데이터를 다시 JSON으로 직렬화
    updatedBogoJSON, err := json.Marshal(bogo)
    if err != nil {
        return err
    }

    // 업데이트된 Bogo 데이터를 월드 스테이트에 저장
    err = ctx.GetStub().PutState(bogoNoKey, updatedBogoJSON)
    if err != nil {
        return fmt.Errorf("failed to update Bogo in world state: %v", err)
    }

    return nil
}



// BogoExists returns true when bogo with given bogoNo exists in world state
//주어진 bogoNo에 해당하는 bogo가 월드 스테이트에 존재하는지 확인하는 함수
func (s *SmartContract) BogoExists(ctx contractapi.TransactionContextInterface, bogoNo int) (bool, error) {
	bogoNoKey := fmt.Sprintf(bogoKeyFormat, bogoNo) // BogoNo를 사용하여 키 생성
	bogoJSON, err := ctx.GetStub().GetState(bogoNoKey)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return bogoJSON != nil, nil
}

// GetAllBogos returns all bogos found in world state
func (s *SmartContract) GetAllBogos(ctx contractapi.TransactionContextInterface) ([]*Bogo, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all bogos in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var bogos []*Bogo
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var bogo Bogo
		err = json.Unmarshal(queryResponse.Value, &bogo)
		if err != nil {
			return nil, err
		}
		bogos = append(bogos, &bogo)
	}

	return bogos, nil
}
