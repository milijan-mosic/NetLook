package models

import database "github.com/milijan-mosic/net-look/server/database"

type Agent struct {
	database.ID
	Name           string  `json:"name"`
	UpdateInterval float64 `json:"update_interval"`
}

type CPU struct {
	database.ID
	AgentId   string `gorm:"foreignKey:AgentId" json:"agent_id"`
	Number    string `json:"number"`    // index
	Usage     string `json:"usage"`     // percentage
	Date      string `json:"date"`      // local date
	Timestamp string `json:"timestamp"` // UNIX
}

type RAM struct {
	database.ID
	AgentId   string `gorm:"foreignKey:AgentId" json:"agent_id"`
	Total     string `json:"total"`     // GB
	Used      string `json:"used"`      // GB
	Usage     string `json:"usage"`     // percentage
	Date      string `json:"date"`      // local date
	Timestamp string `json:"timestamp"` // UNIX
}

type SSD struct {
	database.ID
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
