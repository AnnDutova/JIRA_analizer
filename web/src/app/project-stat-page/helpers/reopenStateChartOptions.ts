import { Options } from 'highcharts';

export const reopenStateChartOptions: Options = {
  legend: {
    enabled: true,
  },
  credits: {
    enabled: false,
  },
  title: {
    text: 'Reopen tasks',
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

