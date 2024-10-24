package models

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type ID struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`
}

func (base *ID) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return
}

type Agent struct {
	ID
	Name           string  `json:"name"`
	UpdateInterval float64 `json:"update_interval"`
}

type CPU struct {
	ID
	AgentId   string `gorm:"foreignKey:AgentId" json:"agent_id"`
	Number    string `json:"number"`    // index
	Usage     string `json:"usage"`     // percentage
	Date      string `json:"date"`      // local date
	Timestamp string `json:"timestamp"` // UNIX
}

type RAM struct {
	ID
	AgentId   string `gorm:"foreignKey:AgentId" json:"agent_id"`
	Total     string `json:"total"`     // GB
	Used      string `json:"used"`      // GB
	Usage     string `json:"usage"`     // percentage
	Date      string `json:"date"`      // local date
	Timestamp string `json:"timestamp"` // UNIX
}

type SSD struct {
	ID
	AgentId   string `gorm:"foreignKey:AgentId" json:"agent_id"`
	Total     string `json:"total"`     // GB
	Used      string `json:"used"`      // GB
	Usage     string `json:"usage"`     // percentage
	Date      string `json:"date"`      // local date
	Timestamp string `json:"timestamp"` // UNIX
}

var AllModels = []interface{}{
	&Agent{},
	&CPU{},
	&RAM{},
	&SSD{},
}

// JSON UNMARSHALLING

type TimeStamp struct {
	Date      string `json:"date"`      // local date
	Timestamp string `json:"timestamp"` // UNIX
}

type Package struct {
	Time TimeStamp `json:"time"`
	Cpu  []CPU     `json:"cpu"`
	Ram  RAM       `json:"ram"`
	Ssd  SSD       `json:"ssd"`
}
