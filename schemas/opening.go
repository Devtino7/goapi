package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID        uint      `json: "id"`
	CreatedAt time.Time `json: "createdat"`
	UpdateAt  time.Time `json: "updateat"`
	DeleteAt  time.Time `json: "deleteat, omitepmty"`
	Role      string    `json: "role"`
	Company   string    `json: "company"`
	Location  string    `json: "location"`
	Remote    bool      `json: "Remote"`
	Link      string    `json: "link"`
	Salary    int64     `json: "salary"`
}
