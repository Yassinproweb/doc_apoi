import { Chart } from "chart.js/auto";

Chart.defaults.font.family = 'Familjen Grotesk';

const ctx = document.getElementById("graph").getContext('2d');
new Chart(ctx, {
  type: 'doughnut',
  data: {
    labels: [
      'Online',
      'Offline'
    ],
    datasets: [{
      data: [146, 85],
      backgroundColor: ['rgb(56, 156, 223, 0.55)', 'rgb(56, 156, 223, 0.75)'],
      borderWidth: 2,
      borderColor: '#fefeff'
    }]
  },
  options: {
    responsive: true,
    cutout: '45%',
    plugins: {
      legend: {
        position: 'bottom',
        font: {
          size: 16
        }
      },
      title: {
        display: true,
        text: 'Monthly Appointments',
        color: '#1A1B23',
        font: {
          family: 'Mozilla Headline',
          size: 18,
          bold: 700
        }
      }
    }
  }
});
