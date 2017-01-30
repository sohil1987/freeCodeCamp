(function () {
  'use strict';
  /* code here */

  /* quantize and invertExtent https://github.com/d3/d3-scale */

  var urlData =
  'https://raw.githubusercontent.com/FreeCodeCamp/ProjectReferenceData/master/global-temperature.json';
  var dataset = [];

  function init () {
    console.log('Inicio');
    getData();
    console.log('End');
  }

  function goOn () {
    var baseT = dataset.baseTemperature;
    dataset = dataset.monthlyVariance;
    // console.log(dataset.length)
    // There are 263 years

    var months = ['January', 'February', 'March', 'April', 'May', 'June',
      'July', 'August', 'September', 'October', 'November', 'December'
    ];

    var colors = ['#004c00', '#008000', '#329932', '#66b266', '#99cc99',
      '#cce5cc', '#ffdae0', '#ff9999', '#ff6666', '#ff3232', '#ff0000',
      '#990000'
    ];

    var margin = {
      top: 10,
      right: 41,
      bottom: 80,
      left: 130
    };
    var w = 960 - margin.left - margin.right;
    var h = 450 - margin.top - margin.bottom;

    var minYear = new Date(d3.min(dataset, function (d) {
      return d.year;
    }), 0);
    var maxYear = new Date(d3.max(dataset, function (d) {
      return d.year;
    }), 0);
    var parseYear = d3.timeFormat('%Y');
    var numYears = parseYear(maxYear) - parseYear(minYear) + 1;

    var minVariance = d3.min(dataset, function (d) {
      return d.variance;
    });
    var maxVariance = d3.max(dataset, function (d) {
      return d.variance;
    });

    // console.log(baseT + minVariance, ' --> ', maxVariance + baseT)
    // console.log(numYears)
    // log(parseYear(minYear), ' --> ', parseYear(maxYear))

    var yScale = d3.scaleLinear()
      .domain([0, months.length])
      .range([0, h]);

    var yAxis = d3.axisLeft(yScale)
      .tickFormat(function (d, i) {
        return months[i];
      });

    var xScale = d3.scaleTime()
      .domain([minYear, maxYear])
      .range([0, w]);

    var xAxis = d3.axisBottom(xScale)
      .ticks(10);

    var colorScale = d3.scaleQuantize()
      .domain([minVariance + baseT, maxVariance + baseT])
      .range(colors);

    var svg = d3.select('.heatMap')
      .attr('width', w + margin.left + margin.right)
      .attr('height', h + margin.top + margin.bottom)
      .attr('class', 'heatMap')
      // .style('background', '#fafafa')
      .append('g')
      .attr('transform', 'translate(' + margin.left + ', ' + margin.top +
        ')');

    svg.selectAll('.bar')
      .data(dataset)
      .enter()
      .append('rect')
      .attr('x', function (d) {
        return (d.year - parseYear(minYear)) * (w / numYears);
      })
      .attr('y', function (d) {
        // return yScale(d.month - 1) // same as below
        return (d.month - 1) * (h / months.length);
      })
      .attr('width', w / numYears)
      .attr('height', h / months.length)
      .attr('fill', function (d) {
        return getColor(d);
      })
      .on('mouseover', function (d) {
        var date = months[d.month - 1] + ' - ' + d.year;
        var temp = 'Average : ' + (d.variance + baseT).toFixed(3) + ' ºC';
        var variation = 'Variance : ' + d.variance.toFixed(3) + ' ºC';
        // d3.select(this)
        //  .style('fill', 'white')
        d3.select('#tooltip')
          .style('left', (d3.event.pageX - 100) + 'px')
          .style('top', (d3.event.pageY - 110) + 'px')
          .select('#date')
          .text(date);
        d3.select('#tooltip')
          .select('#temp')
          .text(temp);
        d3.select('#tooltip')
          .select('#variation')
          .text(variation);
        d3.select('#tooltip').classed('hidden', false);
      })

      .on('mouseout', function (d) {
        d3.select('#tooltip').classed('hidden', true);
      });

    svg.append('g')
      .attr('class', 'axis')
      .call(yAxis)
      .call(function centerLabels (selection) {
        var hop = h / months.length;
        selection.selectAll('text')
          .attr('transform', 'translate(0,' + (hop / 2) + ')');
      });

    svg.append('g')
      .attr('transform', 'translate(0,' + h + ')')
      .attr('class', 'axis')
      .call(xAxis);

    // legends
    svg.selectAll('.legend')
      .data(colors)
      .enter()
      .append('rect')
      .attr('x', function (d, i) {
        var ancho = w / 25;
        return (ancho * i) + (w / 4);
      })
      .attr('y', h + 40)
      .attr('width', w / 25)
      .attr('height', 10)
      .style('fill', function (d, i) {
        return colors[i];
      });

    // legends text
    svg.selectAll('.legend')
      .data(colors)
      .enter()
      .append('text')
      .attr('class', 'legend')
      .attr('x', function (d, i) {
        var ancho = w / 25;
        return (ancho * i) + (w / 4);
      })
      .attr('y', h + 65)
      .text(function (d) {
        var r = colorScale.invertExtent(d);
        return (r[0].toFixed(1));
      });

    svg.append('g')
      .append('text')
      .attr('class', 'legend')
      .attr('x', 90)
      .attr('y', h + 55)
      .text('Color Scale (ºC)');

    function getColor (d) {
      var temp = (d.variance + baseT).toFixed(3);
      return colorScale(temp);
    }
  }

  function getData () {
    // console.log('Getting data ...')
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      var DONE = 4;
      var OK = 200;
      if (xhr.readyState === DONE) {
        if (xhr.status === OK) {
          dataset = JSON.parse(xhr.responseText);
          goOn();
        } else {
          console.log('Error: ' + xhr.status);
        }
      }
    };
    xhr.open('GET', urlData); // add false to synchronous request
    xhr.send();
  }

  window.addEventListener('load', init);
}());
