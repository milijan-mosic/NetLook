package metrics

import (
	disk "github.com/shirou/gopsutil/disk"

	models "github.com/milijan-mosic/net-look/agent/models"
	utils "github.com/milijan-mosic/net-look/agent/utils"
)

func GetSSDUsage() models.SSD {
	diskStat, _ := disk.Usage("/")

	diskUsage := models.SSD{
		Total: utils.FloatToString(diskStat.UsedPercent),
		Used:  utils.FloatToString(float64(diskStat.Total) / utils.BytesToGigabytes()),
		Usage: utils.FloatToString(float64(diskStat.Used) / utils.BytesToGigabytes()),
	}

	return diskUsage
}
