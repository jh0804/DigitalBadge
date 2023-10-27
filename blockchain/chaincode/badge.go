// 배지 관련 코드
package main

import (
	"encoding/json"
	"errors"
	"fmt"

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
}

type BadgeContract struct {
	contractapi.Contract
}

func (c *BadgeContract) Init(ctx contractapi.TransactionContextInterface) error {
	// 초기 데이터베이스 설정 또는 초기 데이터 로드를 수행
	return nil
}

// 배지 발급
func (c *BadgeContract) IssueBadge(ctx contractapi.TransactionContextInterface, badgeID, issuer, description, ownerID string) error {
	// Check if the client has the 'admin' role
	isAdmin, err := c.checkUserRole(ctx, "admin")
	if err != nil {
		return err
	}
	if !isAdmin {
		return errors.New("Only admins can issue badges")
	}

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

	// Add badgeID to the user's badge list
	userKey := "USER_" + ownerID
	userJSON, err := ctx.GetStub().GetState(userKey)
	if err != nil {
		return err
	}

	var user User
	if userJSON == nil {
		user = User{
			ID:        ownerID,
			Name:      "Sample User", // 사용자의 이름 또는 다른 정보를 추가할 수 있습니다.
			BadgeIDs:  []string{badgeID},
		}
	} else {
		err = json.Unmarshal(userJSON, &user)
		if err != nil {
			return err
		}

		user.BadgeIDs = append(user.BadgeIDs, badgeID)
	}

	userJSON, err = json.Marshal(user)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(userKey, userJSON)
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
