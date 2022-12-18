import { Options } from 'highcharts';

export const openStateChartOptions: Options = {
  legend: {
    enabled: true,
  },
  credits: {
    enabled: false,
  },
  title: {
    text: 'Open tasks',
  },
  yAxis: {
    visible: true,
  },

  xAxis: {
    visible: true,

    categories: [],
  },



  series: [],
};

