package metrics

import (
	mem "github.com/shirou/gopsutil/mem"

	models "github.com/milijan-mosic/net-look/client/models"
	utils "github.com/milijan-mosic/net-look/client/utils"
)

func GetRAMUsage() models.RAM {
	vMem, _ := mem.VirtualMemory()

	ramUsage := models.RAM{
		Total: utils.FloatToString(float64(vMem.Total) / utils.BytesToGigabytes()),
		Used:  utils.FloatToString(float64(vMem.Used) / utils.BytesToGigabytes()),
		Usage: utils.FloatToString(vMem.UsedPercent),
	}
	// DEBUG
	// fmt.Printf("RAM Usage -> %.2f%%, Total -> %.2f GB, Used -> %.2f GB\n",
	// 	vMem.UsedPercent,
	// 	float64(vMem.Total)/(1024*1024*1024),
	// 	float64(vMem.Used)/(1024*1024*1024),
	// )

	return ramUsage
}
