package metrics

import (
	cpu "github.com/shirou/gopsutil/cpu"

	models "github.com/milijan-mosic/net-look/agent/models"
	utils "github.com/milijan-mosic/net-look/agent/utils"
)

func GetCPUsUsage() []models.CPU {
	allCpus, _ := cpu.Percent(0, true)

	var cpuReport []models.CPU
	for i, percent := range allCpus {
		item := models.CPU{
			Number: utils.IntegerToString(i + 1),
			Usage:  utils.FloatToString(percent),
		}
		cpuReport = append(cpuReport, item)
	}

	return cpuReport
}
