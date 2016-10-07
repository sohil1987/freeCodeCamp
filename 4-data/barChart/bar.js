(function () {
  'use strict';
  /* code here */

  var urlData =
  'https://raw.githubusercontent.com/FreeCodeCamp/ProjectReferenceData/master/GDP-data.json';
  var dataset = [];

  function init () {
    console.log('Inicio');
    getData();
  }

  function goOn () {
    // console.log(data)
    //  console.log(dataset)
    document.getElementsByClassName('notes')[0].innerText = dataset.description;

    var margin = {
      top: 30,
      right: 10,
      bottom: 30,
      left: 50
    };
    var w = 885 - margin.left - margin.right;
    var h = 600 - margin.top - margin.bottom;
    var barPadding = 0;
    var quarters = dataset.data.length;
    var data = dataset.data;
    var monthName = ['January', 'February', 'March', 'April', 'May', 'June',
      'July', 'August', 'September', 'October', 'November', 'December'
    ];

    var min = d3.min(data, function (d) {
      return d[1];
    });
    var max = d3.max(data, function (d) {
      return d[1];
    });
    var minDate = new Date(data[0][0]);
    var maxDate = new Date(data[quarters - 1][0]);

    var yScale = d3.scaleLinear()
      .domain([0, max])
      .range([h, 0]);

    var yAxis = d3.axisLeft(yScale);

    var xScale = d3.scaleTime()
      .domain([minDate, maxDate])
      .range([0, w]);

    var xAxis = d3.axisBottom(xScale);

    // console.log(min, ' --> ', max)
    // console.log(yScale(min), ' --> ', yScale(max))
    // console.log(minDate, ' --> ', maxDate)

    var svg = d3.select('.barChartSvg')
      .attr('width', w + margin.left + margin.right)
      .attr('height', h + margin.top + margin.bottom)
      .attr('class', 'barChartSvg')
      // .style('background', '#fafafa')
      .append('g')
      .attr('transform', 'translate(' + margin.left + ', ' + margin.top +
        ')');

    svg.selectAll('.bar')
      .data(data)
      .enter()
      .append('rect')
      .attr('class', 'bar')
      .attr('x', function (d, i) {
        return i * w / quarters;
      })
      .attr('y',
        function (d) {
          return yScale(d[1]);
        })
      .attr('width', w / quarters - barPadding)
      .attr('height', function (d) {
        return h - yScale(d[1]);
      })
      .attr('fill', 'teal')

      .on('mouseover', function (d) {
        var date = new Date(d[0]);
        var year = date.getFullYear();
        var month = monthName[date.getMonth()];
        date = year + ' - ' + month;
        // console.log(date)
        d3.select(this)
          .style('fill', 'skyblue');
        d3.select('#tooltip')
          .select('#value')
          .text(d[1]);
        d3.select('#tooltip')
          .style('left', d3.event.pageX - 10 + 'px')
          .style('top', d3.event.pageY - 40 + 'px')
          .select('#date')
          .text(date);
        d3.select('#tooltip').classed('hidden', false);
      })

      .on('mouseout', function (d) {
        d3.select(this)
          .style('fill', 'teal');
        d3.select('#tooltip').classed('hidden', true);
      });

    svg.append('g')
      .attr('class', 'axis')
      .call(yAxis);

    svg.append('g')
      .attr('transform', 'translate(0,' + h + ')')
      .attr('class', 'axis')
      .call(xAxis);

    console.log('END');
  }

  function getData () {
    console.log('Getting data ...');
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
