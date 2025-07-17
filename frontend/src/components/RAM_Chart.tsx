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

export const RAMChart: React.FC<ChartData> = ({ data }) => {
  return (
    <div className="p-4 bg-white rounded-2xl shadow-md">
      <h2 className="text-xl font-semibold mb-4">RAM Usage</h2>
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
            name="RAM Usage"
            stroke="#82ca9d"
            strokeWidth={2}
            dot={false}
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
};
