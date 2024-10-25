import { onMount } from "solid-js";
import {
  Chart,
  BarController,
  BarElement,
  CategoryScale,
  LinearScale,
  Title,
} from "chart.js";

// Register required components with Chart.js
Chart.register(BarController, BarElement, CategoryScale, LinearScale, Title);

function SimpleChart() {
  let canvasRef;

  onMount(() => {
    const ctx = canvasRef.getContext("2d");
    new Chart(ctx, {
      type: "bar",
      data: {
        labels: [
          "January",
          "February",
          "March",
          "April",
          "May",
          "June",
          "July",
        ],
        datasets: [
          {
            label: "Monthly Sales",
            data: [65, 59, 80, 81, 56, 55, 40],
            backgroundColor: "rgba(75, 192, 192, 0.2)",
            borderColor: "rgba(75, 192, 192, 1)",
            borderWidth: 1,
          },
        ],
      },
      options: {
        responsive: true,
        plugins: {
          title: {
            display: true,
            text: "Monthly Sales Data",
          },
        },
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      },
    });
  });

  return (
    <div class="flex justify-center flex-col items-center">
      <h1>Simple Bar Chart</h1>
      <canvas ref={canvasRef}></canvas>
    </div>
  );
}

export default SimpleChart;
