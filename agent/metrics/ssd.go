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
	// DEBUG
	// fmt.Printf("SSD Usage -> %.2f%%, Total -> %.2f GB, Used -> %.2f GB\n",
	// 	diskStat.UsedPercent,
	// 	float64(diskStat.Total)/(1024*1024*1024),
	// 	float64(diskStat.Used)/(1024*1024*1024))

	return diskUsage
}
