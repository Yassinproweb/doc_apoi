import { Chart } from "chart.js/auto";

const ctx = document.getElementById("graph").getContext('2d');
new Chart(ctx, {
  type: 'doughnut',
  data: {
    labels: [
      'Offline',
      'Online'
    ],
    datasets: [{
      data: [85, 146],
      backgroundColor: ['rgb(116, 216, 253)', 'rgb(76, 176, 243)'],
      borderWidth: 2,
      borderColor: '#fefeff'
    }]
  },
  options: {
    responsive: true,
    cutout: '45%',
    plugins: {
      legend: {
        position: 'bottom'
      },
      title: {
        display: true,
        text: 'Monthly Appointments'
      }
    }
  }
});
