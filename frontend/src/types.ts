export interface MetricPoint {
  timestamp: number;
  usage: number;
  number?: string;
  total?: string;
  used?: string;
}

export interface ChartData {
  data: MetricPoint[];
}

export interface AgentRaw {
  cpus: {
    timestamp: number;
    usage: string;
  }[];
  ram: {
    timestamp: number;
    usage: string;
  }[];
  ssd: {
    timestamp: number;
    usage: string;
  }[];
}

export interface ApiResponse {
  data: {
    agents: AgentRaw[];
  };
}
