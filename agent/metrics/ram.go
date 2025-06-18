package metrics

import (
	mem "github.com/shirou/gopsutil/mem"

	models "github.com/milijan-mosic/net-look/agent/models"
	utils "github.com/milijan-mosic/net-look/agent/utils"
)

func GetRAMUsage() models.RAM {
	vMem, _ := mem.VirtualMemory()

	ramUsage := models.RAM{
		Total: utils.FloatToString(float64(vMem.Total) / utils.BytesToGigabytes()),
		Used:  utils.FloatToString(float64(vMem.Used) / utils.BytesToGigabytes()),
		Usage: utils.FloatToString(vMem.UsedPercent),
	}

	return ramUsage
}
