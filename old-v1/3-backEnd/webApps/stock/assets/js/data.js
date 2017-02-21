var datos = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  var seriesOptions = [];
  var seriesCounter = 0;
  var names = []; // ['MSFT', 'AAPL', 'GOOG']

  function init () {
    console.log('Inicio DATOS');
  }

  function createChart () {
    Highcharts.stockChart('container', {
      rangeSelector: {
        selected: 4
      },
      yAxis: {
        labels: {
          formatter: function () {
            return (this.value > 0 ? ' + ' : '') + this.value + '%';
          }
        },
        plotLines: [{
          value: 0,
          width: 2,
          color: 'silver'
        }]
      },
      plotOptions: {
        series: {
          compare: 'percent',
          showInNavigator: true
        }
      },
      tooltip: {
        pointFormat: '<span style="color:{series.color}">{series.name}</span>: <b>{point.y}</b> ({point.change}%)<br/>',
        valueDecimals: 2,
        split: true
      },
      series: seriesOptions
    });
  }

  function readData () {
    $.each(datos.names, function (i, name) {
      var url = app.baseURL + 'stock/api/get/?code=' + name.toLowerCase();
      console.log('PIDIENDO', url);
      $.getJSON(url, function (data) {
        seriesOptions[i] = {
          name: name,
          data: data
        };
        seriesCounter += 1;
        if (seriesCounter === datos.names.length) {
          createChart();
        }
      });
    });
  }

  return {
    inicio: init,
    readData: readData,
    names: names
  };
}());
