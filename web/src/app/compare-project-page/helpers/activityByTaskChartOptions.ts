import { Options } from 'highcharts';

export const activityByTaskChartOptions: Options = {
  legend: {
    enabled: true,
  },
  credits: {
    enabled: false,
  },
  title: {
    text: 'Activity by task',
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

