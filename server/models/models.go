package models

type Agent struct {
	ID             string  `gorm:"type:string;primaryKey" json:"id"`
	Name           string  `json:"name"`
	UpdateInterval float64 `json:"update_interval"`
}

type CPU struct {
	ID        string `gorm:"type:string;primaryKey" json:"id"`
	AgentID   string `gorm:"type:string;foreignKey:AgentID" json:"agent_id"`
	Number    string `json:"number"`    // thread number
	Usage     string `json:"usage"`     // percentage
	Date      string `json:"date"`      // local date
	Timestamp int64  `json:"timestamp"` // UNIX
}

type RAM struct {
	ID        string `gorm:"type:string;primaryKey" json:"id"`
	AgentID   string `gorm:"type:string;foreignKey:AgentID" json:"agent_id"`
	Total     string `json:"total"`     // GB
	Used      string `json:"used"`      // GB
	Usage     string `json:"usage"`     // percentage
	Date      string `json:"date"`      // local date
	Timestamp int64  `json:"timestamp"` // UNIX
}

type SSD struct {
	ID        string `gorm:"type:string;primaryKey" json:"id"`
	AgentID   string `gorm:"type:string;foreignKey:AgentID" json:"agent_id"`
	Total     string `json:"total"`     // GB
	Used      string `json:"used"`      // GB
	Usage     string `json:"usage"`     // percentage
	Date      string `json:"date"`      // local date
	Timestamp int64  `json:"timestamp"` // UNIX
}

type AgentData struct {
	AgentName string `json:"agent_name"`
	CPUs      []CPU  `json:"cpus"`
	RAM       []RAM  `json:"ram"`
	SSD       []SSD  `json:"ssd"`
}

type ClientPackage struct {
	Agents []AgentData `json:"agents"`
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
	Timestamp int64  `json:"timestamp"` // UNIX
}

type Package struct {
	AgentName string    `json:"agent_name"`
	Time      TimeStamp `json:"time"`
	Cpu       []CPU     `json:"cpu"`
	Ram       RAM       `json:"ram"`
	Ssd       SSD       `json:"ssd"`
}
