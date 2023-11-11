package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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
	BogoNo    int    `json:"BogoNo"`
	Title     string `json:"Title"`
	Author    string `json:"Author"`
	Publisher string `json:"Publisher"`
	Report    string `json:"Report"`
	Owner     string `json:"Owner"`    //독후감 작성자 - 이메일(=id 역할)입력되도록?
	Approved  bool   `json:"Approved"` //초기 create시 false
}

// AssetCounter는 BogoNo를 저장하는 데 사용
type AssetCounter struct {
	Counter int `json:"Counter"`
}

type Badge struct {
	BadgeNo     int    `json:"BadgeNo"`     //배지 고유 번호
	Name        string `json:"Name"`        //배지 이름
	IssueDate   string `json:"IssueDate"`   //발급일 (YYYY.MM.DD)
	IssuerId    string `json:"IssuerId"`    //발급자 ID
	IssuerName  string `json:"IssuerName"`  //발급자명(부경대 도서관 관리자)
	Recipient   string `json:"Recipient"`   //배지 받는 사람(학생 ID)
	Image       string `json:"Image"`       //배지 이미지
	Level       string `json:"Level"`       //배지 단계
	Description string `json:"Description"` //배지 설명
}

const bogoKeyFormat = "bogo%04d"
const badgeKeyFormat = "badge%04d"
const DDMMYYYY = "2006-01-02"

// InitLedger
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

// CreateBogo: Bogo 생성
func (s *SmartContract) CreateBogo(ctx contractapi.TransactionContextInterface, title string, author string, publisher string, report string, owner string) (int, error) {

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

	// BogoNo +1 증가
	assetCounter.Counter++
	bogoNo := assetCounter.Counter

	// Bogo 생성 (BogoNo=1부터 시작)
	bogo := Bogo{
		BogoNo:    bogoNo,
		Title:     title,
		Author:    author,
		Publisher: publisher,
		Report:    report,
		Owner:     owner,
		Approved:  false,
	}

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

// ApproveBogo: 도서관 관리자가 독후감을 승인/Bogo의 Approved 필드를 true로 업데이트.
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
// 주어진 bogoNo에 해당하는 bogo가 월드 스테이트에 존재하는지 확인하는 함수
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

// GetBogosByOwner 함수: owner를 기준으로 bogo 조회
func (s *SmartContract) GetBogosByOwner(ctx contractapi.TransactionContextInterface, owner string) ([]*Bogo, error) {
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

		// Owner 필드가 원하는 값과 일치하는 경우에만 추가
		if bogo.Owner == owner {
			bogos = append(bogos, &bogo)
		}
	}

	return bogos, nil
}

// GetUnapprovedBogos 함수: Approved가 false인 Bogo들의 목록을 조회.
func (s *SmartContract) GetUnapprovedBogos(ctx contractapi.TransactionContextInterface) ([]*Bogo, error) {
	// 모든 Bogo 조회
	allBogos, err := s.GetAllBogos(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all bogos: %v", err)
	}

	// Approved가 false인 Bogo들을 필터링
	var unapprovedBogos []*Bogo
	for _, bogo := range allBogos {
		if !bogo.Approved {
			unapprovedBogos = append(unapprovedBogos, bogo)
		}
	}

	return unapprovedBogos, nil
}

////////////////////////
//////배지 관련 함수/////
////////////////////////

// IssueBadge 함수 : 배지를 발급
func (s *SmartContract) IssueBadge(ctx contractapi.TransactionContextInterface, recipient string, issuerId string) (int, error) {
	// 도서관 조직의 유저가 가진 bogo 개수 조회
	bogos, err := s.GetBogosByOwnerAndApproval(ctx, recipient)
	if err != nil {
		return 0, fmt.Errorf("failed to get bogos for recipient: %v", err)
	}

	// 총 보고서 개수 계산
	totalBogos := len(bogos)

	// Existing badges for the recipient
	existingBadges, err := s.GetBadgesByRecipient(ctx, recipient)
	if err != nil {
		return 0, fmt.Errorf("failed to get existing badges for recipient: %v", err)
	}

	// 배지를 owner 기준으로 읽어와서 가장 높은 단계 확인
	highestBadgeLevel := getHighestBadgeLevel(existingBadges)

	// 적절한 배지 단계 결정
	var badgeLevel string
	var badgeImage string
	var badgeDescription string

	if totalBogos >= 1 && totalBogos <= 5 {
		if highestBadgeLevel == 0 {
			badgeLevel = "1"
			badgeImage = "https://example.com/badge_image_1"
			badgeDescription = "This is a level 1 badge for book reports. Issued when the number of approved book reports is 1 to 5."
		} else {
			return 0, fmt.Errorf("recipient already has a badge of level %d or higher", highestBadgeLevel)
		}
	} else if totalBogos >= 6 && totalBogos <= 10 {
		if highestBadgeLevel == 1 {
			badgeLevel = "2"
			badgeImage = "https://example.com/badge_image_2"
			badgeDescription = "This is a level 2 badge for book reports. Issued when the number of approved book reports is 6 to 10."
		} else {
			return 0, fmt.Errorf("recipient already has a badge of level %d or higher", highestBadgeLevel)
		}
	} else if totalBogos >= 11 && totalBogos <= 15 {
		if highestBadgeLevel == 1 {
			badgeLevel = "3"
			badgeImage = "https://example.com/badge_image_3"
			badgeDescription = "This is a level 3 badge for book reports. Issued when the number of approved book reports is 11 to 15."
		} else {
			return 0, fmt.Errorf("recipient already has a badge of level %d or higher", highestBadgeLevel)
		}
	} else {
		// 적절한 조건에 해당하지 않는 경우 배지 발급하지 않음
		return 0, fmt.Errorf("no badge issuance condition satisfied or no more badges to issue")
	}

	// 새로운 배지 생성
	assetCounterJSON, err := ctx.GetStub().GetState("BadgeNoCounter")
	if err != nil {
		return 0, fmt.Errorf("failed to read BadgeNoCounter from world state: %v", err)
	}

	var badgeCounter AssetCounter
	if assetCounterJSON != nil {
		err = json.Unmarshal(assetCounterJSON, &badgeCounter)
		if err != nil {
			return 0, fmt.Errorf("failed to unmarshal BadgeNoCounter: %v", err)
		}
	}

	// BadgeNo +1 증가
	badgeCounter.Counter++
	badgeNo := badgeCounter.Counter

	// 현재 날짜를 YYYY-MM-DD 형식으로 포맷팅
	issueDate := time.Now().UTC().Format(DDMMYYYY)

	// 새로운 배지 생성
	badge := Badge{
		BadgeNo:     badgeNo,
		Name:        fmt.Sprintf("Book Report Badge Lv %s.", badgeLevel),
		IssueDate:   issueDate,
		IssuerId:    issuerId,
		IssuerName:  fmt.Sprintf("부경대 도서관 관리자"),
		Recipient:   recipient,
		Image:       badgeImage,
		Level:       badgeLevel,
		Description: badgeDescription,
	}

	// Badge 구조체를 JSON으로 직렬화
	badgeJSON, err := json.Marshal(badge)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal Badge: %v", err)
	}

	// 새로운 Badge를 월드 스테이트에 저장
	err = ctx.GetStub().PutState(fmt.Sprintf(badgeKeyFormat, badgeNo), badgeJSON)
	if err != nil {
		return 0, fmt.Errorf("failed to put Badge to world state: %v", err)
	}

	// 업데이트된 BadgeNo 카운터 정보를 JSON으로 직렬화
	updatedCounterJSON, err := json.Marshal(badgeCounter)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal updated BadgeNoCounter: %v", err)
	}

	// 업데이트된 BadgeNo 카운터를 "BadgeNoCounter" 키로 원장에 저장
	err = ctx.GetStub().PutState("BadgeNoCounter", updatedCounterJSON)
	if err != nil {
		return 0, fmt.Errorf("failed to update BadgeNoCounter in world state: %v", err)
	}

	return badgeNo, nil
}

// GetBogosByOwnerAndApproval 함수: owner를 기준으로 bogo 조회 (Approval 상태에 따라 필터링)
func (s *SmartContract) GetBogosByOwnerAndApproval(ctx contractapi.TransactionContextInterface, owner string) ([]*Bogo, error) {
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

		// 필터링: owner가 일치하고 Approved 상태가 true인 Bogo만 반환
		if bogo.Owner == owner && bogo.Approved {
			bogos = append(bogos, &bogo)
		}
	}

	return bogos, nil
}


// getHighestBadgeLevel 함수: 가장 높은 단계의 배지 레벨을 찾아서 반환
func getHighestBadgeLevel(badges []*Badge) int {
	var highestLevel int

	for _, badge := range badges {
		level, err := strconv.Atoi(badge.Level)
		if err == nil && level > highestLevel {
			highestLevel = level
		}
	}

	return highestLevel
}

// GetBadgeByBadgeNo 함수: badgeNo를 기준으로 특정 배지를 조회
func (s *SmartContract) GetBadgeByBadgeNo(ctx contractapi.TransactionContextInterface, badgeNo int) (*Badge, error) {
	badgeKey := fmt.Sprintf(badgeKeyFormat, badgeNo)
	badgeJSON, err := ctx.GetStub().GetState(badgeKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read badge from world state: %v", err)
	}

	if badgeJSON == nil {
		return nil, fmt.Errorf("badge %d does not exist", badgeNo)
	}

	var badge Badge
	err = json.Unmarshal(badgeJSON, &badge)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal badge: %v", err)
	}

	return &badge, nil
}

// GetBadgesByRecipient 함수: recipient를 기준으로 배지 조회
func (s *SmartContract) GetBadgesByRecipient(ctx contractapi.TransactionContextInterface, recipient string) ([]*Badge, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var badges []*Badge
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var badge Badge
		err = json.Unmarshal(queryResponse.Value, &badge)
		if err != nil {
			return nil, err
		}

		// 필터링: recipient가 일치하는 Badge만 반환
		if badge.Recipient == recipient {
			badges = append(badges, &badge)
		}
	}

	return badges, nil
}
