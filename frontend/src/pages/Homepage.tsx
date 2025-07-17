import { useEffect, useState } from "react";
import { CPUChart } from "../components/CPU_Chart";
import { RAMChart } from "../components/RAM_Chart";
import { SSDChart } from "../components/SSD_Chart";
import { type ApiResponse } from "../types";

interface MetricPoint {
  timestamp: number;
  usage: number;
}

function Homepage() {
  const url: string = "/api/1.0/metrics";

  // Use the MetricPoint type for your state
  const [cpuData, setCpuData] = useState<MetricPoint[]>([]);
  const [ramData, setRamData] = useState<MetricPoint[]>([]);
  const [ssdData, setSsdData] = useState<MetricPoint[]>([]);

  useEffect(() => {
    const fetchData = () => {
      fetch(url)
        .then((res) => res.json())
        .then((json: ApiResponse) => {
          const agent = json.data.agents[0];

          const cpuData: MetricPoint[] = agent.cpus.map((c) => ({
            timestamp: c.timestamp,
            usage: parseFloat(c.usage),
          }));

          const ramData: MetricPoint[] = agent.ram.map((r) => ({
            timestamp: r.timestamp,
            usage: parseFloat(r.usage),
          }));

          const ssdData: MetricPoint[] = agent.ssd.map((s) => ({
            timestamp: s.timestamp,
            usage: parseFloat(s.usage),
          }));

          setCpuData(cpuData);
          setRamData(ramData);
          setSsdData(ssdData);
        })
        .catch((err) => console.error("Failed to fetch or parse JSON", err));
    };

    fetchData();

    const intervalId = setInterval(fetchData, 1000);
    return () => clearInterval(intervalId);
  }, []);

  return (
    <div className="flex flex-col bg-black text-white">
      <div className="flex flex-col justify-center items-center">
        <h1 className="text-3xl font-bold my-8">NetLook B)</h1>

        <div className="w-full p-4 space-y-6">
          <CPUChart data={cpuData} />
          <RAMChart data={ramData} />
          <SSDChart data={ssdData} />
        </div>
      </div>
    </div>
  );
}

export default Homepage;
