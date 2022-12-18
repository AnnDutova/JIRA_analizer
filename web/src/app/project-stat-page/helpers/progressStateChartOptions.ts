import { Options } from 'highcharts';

export const progressStateChartOptions: Options = {
  legend: {
    enabled: true,
  },
  credits: {
    enabled: false,
  },
  title: {
    text: 'Progress tasks',
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

