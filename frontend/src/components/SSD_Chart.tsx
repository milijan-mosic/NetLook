import React from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
  ResponsiveContainer,
  Legend,
} from "recharts";
import { type ChartData } from "../types";

export const SSDChart: React.FC<ChartData> = ({ data }) => {
  return (
    <div className="p-4 rounded-2xl shadow-md">
      <h2 className="text-xl font-semibold mb-4">SSD Usage</h2>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={data}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis
            dataKey="timestamp"
            tickFormatter={(ts) => new Date(ts * 1000).toLocaleTimeString()}
          />
          <YAxis domain={[0, 100]} unit="%" />
          <Tooltip
            labelFormatter={(ts) => new Date(ts * 1000).toLocaleString()}
          />
          <Legend />
          <Line
            type="monotone"
            dataKey="usage"
            name="SSD Usage"
            stroke="#ff7300"
            strokeWidth={2}
            dot={false}
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
};
