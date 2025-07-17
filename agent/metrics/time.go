package metrics

import (
	"time"

	models "github.com/milijan-mosic/net-look/agent/models"
)

func GetTimeData() models.TimeStamp {
	currentTime := time.Now()
	dateTime := currentTime.Format("02.01.2006 15:04:05")
	unixTimestamp := currentTime.Unix()

	timeStamp := models.TimeStamp{
		Date:      dateTime,
		Timestamp: unixTimestamp,
	}

	return timeStamp
}
