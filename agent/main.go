package main

import (
	"log"
	"time"

	emitter "github.com/milijan-mosic/net-look/agent/emitter"
	clock "github.com/milijan-mosic/net-look/agent/metrics"
	cpu "github.com/milijan-mosic/net-look/agent/metrics"
	ram "github.com/milijan-mosic/net-look/agent/metrics"
	ssd "github.com/milijan-mosic/net-look/agent/metrics"
	models "github.com/milijan-mosic/net-look/agent/models"
)

var tick = 1 * time.Second

func main() {
	log.Println("Booting up...")

	for {
		cpuUsage := cpu.GetCPUsUsage()
		ramUsage := ram.GetRAMUsage()
		diskUsage := ssd.GetSSDUsage()
		timeData := clock.GetTimeData()

		data := models.Package{
			AgentName: "Test",
			Cpu:       cpuUsage,
			Ram:       ramUsage,
			Ssd:       diskUsage,
			Time:      timeData,
		}

		packagedData := emitter.PackJson(data)
		emitter.Emit(packagedData, "http://master:10000/api/1.0/receive", true)

		time.Sleep(tick)
	}
}
