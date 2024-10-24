package models

type CPU struct {
	Number string `json:"number"` // index
	Usage  string `json:"usage"`  // percentage
}

type RAM struct {
	Total string `json:"total"` // GB
	Used  string `json:"used"`  // GB
	Usage string `json:"usage"` // percentage
}

type SSD struct {
	Total string `json:"total"` // GB
	Used  string `json:"used"`  // GB
	Usage string `json:"usage"` // percentage
}

type TimeStamp struct {
	Date      string `json:"date"`      // local date
	Timestamp string `json:"timestamp"` // UNIX
}

type Package struct {
	AgentName string    `json:"agent_name"`
	Time      TimeStamp `json:"time"`
	Cpu       []CPU     `json:"cpu"`
	Ram       RAM       `json:"ram"`
	Ssd       SSD       `json:"ssd"`
}
