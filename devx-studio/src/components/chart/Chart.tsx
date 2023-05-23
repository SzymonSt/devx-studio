import React, { useRef, useEffect } from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
  } from 'chart.js';
  import { Chart } from 'react-chartjs-2';

type LineChartProps = {
    verticalId: string;
}

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
);

const LineChart: React.FC<LineChartProps> = ({verticalId}) => {
    const options = {
      responsive: true,
      plugins: {
        legend: {
          position: 'top' as const,
        },
        title: {
          display: true,
          text: verticalId,
        },
      },
    };
    const data = {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
      datasets: [
          {
              type: 'line' as const,
              label: 'Dataset 1',
              data: [0, 10, 5, 2, 20, 30, 45],
              borderColor: 'rgb(255, 99, 132)',
          },
          {
              type: 'scatter' as const,
              label: 'Dataset 2',
              data: [, , , 0, , , 0],
              backgroundColor: 'rgb(99, 211, 132, 1)', 
              style: 'circle' as const,
          }
      ]
    };
    return(
            <Chart type='line' data={data} options={options} />
    );
}

export default LineChart;
