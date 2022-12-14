import { Options } from 'highcharts';

export const complexityTaskChartOptions: Options = {
  chart: {
    type: 'column',
  },
  credits: {
    enabled: false,
  },
  title: {
    text: 'Complexity task',
  },
  yAxis: {
    visible: false,
    gridLineColor: '#fff',
  },
  legend: {
    enabled: false,
  },
  xAxis: {
    lineColor: '#fff',
    categories: []
  },

  plotOptions: {
    series: {
      borderRadius: 5,
    } as any,
  },

  series: [
    {
      type: 'column',
      color: '#506ef9',
      data: [],
    },
  ],
};
