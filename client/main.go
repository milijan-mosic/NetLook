package main

import (
	"fmt"
	"time"

	clock "github.com/milijan-mosic/net-look/client/metrics"
	cpu "github.com/milijan-mosic/net-look/client/metrics"
	ram "github.com/milijan-mosic/net-look/client/metrics"
	ssd "github.com/milijan-mosic/net-look/client/metrics"

	models "github.com/milijan-mosic/net-look/client/models"

	emitter "github.com/milijan-mosic/net-look/client/emitter"
)

var tick = 1 * time.Second

func main() {
	fmt.Println("Booting up...")

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
		// fmt.Println(data) // DEBUG

		packagedData := emitter.PackJson(data)
		emitter.Emit(packagedData, "http://localhost:11000/receive", false)

		time.Sleep(tick)
	}
}
