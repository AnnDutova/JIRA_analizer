import { Options } from 'highcharts';

export const resolveStateChartOptions: Options = {
  legend: {
    enabled: true,
  },
  credits: {
    enabled: false,
  },
  title: {
    text: 'Resolve tasks',
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

