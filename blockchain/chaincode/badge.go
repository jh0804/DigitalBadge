// 배지 관련 코드
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Badge struct {
	ID          string `json:"id"`
	Issuer      string `json:"issuer"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

type User struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	BadgeIDs  []string `json:"badgeIds"`
	Count     int      `json:"count"`  // Count 필드 추가
}

type BadgeContract struct {
	contractapi.Contract
}

func (c *BadgeContract) Init(ctx contractapi.TransactionContextInterface) error {
	// 초기 데이터베이스 설정 또는 초기 데이터 로드를 수행
	return nil
}

func (c *BadgeContract) Invoke(ctx contractapi.TransactionContextInterface) error {
	function, args := ctx.GetStub().GetFunctionAndParameters()

	if function == "IssueBadge" {
		return c.IssueBadge(ctx, args[0], args[1], args[2], args[3])
	} else if function == "GetUserBadges" {
		return c.GetUserBadges(ctx, args[0])
	} else if function == "UpdateUserCount" {
		return c.UpdateUserCount(ctx, args[0], args[1])
	} else if function == "GetBadge" {
		return c.GetBadge(ctx, args[0])
	}

	return fmt.Errorf("Invalid function name: %s", function)
}

// 배지 발급 함수
func (c *BadgeContract) IssueBadge(ctx contractapi.TransactionContextInterface, badgeID, issuer, description, ownerID string) error {
	isAdmin, err := c.checkUserRole(ctx, "admin")
	if err != nil {
		return err
	}
	if !isAdmin {
		return errors.New("Only admins can issue badges")
	}

	userKey := "USER_" + ownerID
	userJSON, err := ctx.GetStub().GetState(userKey)
	if err != nil {
		return err
	}
	if userJSON == nil {
		return fmt.Errorf("User with ID %s does not exist", ownerID)
	}

	var user User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return err
	}

	// Count가 5의 배수일 때 새로운 배지 발급
	if user.Count%5 == 0 && user.Count != 0 {
		badge := Badge{
			ID:          badgeID,
			Issuer:      issuer,
			Description: description,
			Owner:       ownerID,
		}

		badgeJSON, err := json.Marshal(badge)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(badgeID, badgeJSON)
		if err != nil {
			return err
		}

		user.BadgeIDs = append(user.BadgeIDs, badgeID)

		// Count를 증가시킴
		user.Count++
		userJSON, err = json.Marshal(user)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(userKey, userJSON)
		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("User's count (%d) is not a multiple of 5", user.Count)
}

// 사용자 배지 목록 조회
func (c *BadgeContract) GetUserBadges(ctx contractapi.TransactionContextInterface, userID string) ([]*Badge, error) {
	// Check if the client has the 'admin' role
	isAdmin, err := c.checkUserRole(ctx, "admin")
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errors.New("Only admins can get user badges")
	}

	userKey := "USER_" + userID
	userJSON, err := ctx.GetStub().GetState(userKey)
	if err != nil {
		return nil, err
	}
	if userJSON == nil {
		return nil, fmt.Errorf("User with ID %s does not exist", userID)
	}

	var user User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, err
	}

	var badges []*Badge
	for _, badgeID := range user.BadgeIDs {
		badge, err := c.GetBadge(ctx, badgeID)
		if err != nil {
			return nil, err
		}
		badges = append(badges, badge)
	}

	return badges, nil
}

// 사용자의 Count 정보 업데이트 함수 추가
func (c *BadgeContract) UpdateUserCount(ctx contractapi.TransactionContextInterface, userID string, countStr string) error {
	isAdmin, err := c.checkUserRole(ctx, "admin")
	if err != nil {
		return err
	}
	if !isAdmin {
		return errors.New("Only admins can update user count")
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return errors.New("Invalid Count. It should be an integer.")
	}

	userKey := "USER_" + userID
	userJSON, err := ctx.GetStub().GetState(userKey)
	if err != nil {
		return err
	}
	if userJSON == nil {
		return fmt.Errorf("User with ID %s does not exist", userID)
	}

	var user User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return err
	}

	// Count 정보 업데이트
	user.Count = count
	userJSON, err = json.Marshal(user)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(userKey, userJSON)
	if err != nil {
		return err
	}

	return nil
}

// 사용자의 배지 조회
func (c *BadgeContract) GetBadge(ctx contractapi.TransactionContextInterface, badgeID string) (*Badge, error) {
	badgeJSON, err := ctx.GetStub().GetState(badgeID)
	if err != nil {
		return nil, err
	}
	if badgeJSON == nil {
		return nil, fmt.Errorf("Badge with ID %s does not exist", badgeID)
	}

	var badge Badge
	err = json.Unmarshal(badgeJSON, &badge)
	if err != nil {
		return nil, err
	}

	return &badge, nil
}

// 사용자 역할 확인(관리자, 일반사용자)
func (c *BadgeContract) checkUserRole(ctx contractapi.TransactionContextInterface, expectedRole string) (bool, error) {
	mspID, err := ctx.GetStub().GetMSPID()
	if err != nil {
		return false, err
	}

	// Compare the client's MSP ID with the expected role
	return mspID == expectedRole, nil
}
