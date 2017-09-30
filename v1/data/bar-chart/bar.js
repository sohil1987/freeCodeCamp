/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const url = 'https://raw.githubusercontent.com/freeCodeCamp/ProjectReferenceData/master/GDP-data.json';
  const urlFile = './bar.json';

  function init () {
    console.log('Init Bar Chart');
    getAjaxData(url, doIt);
  // getAjaxData(urlFile, doIt)
  }

  function doIt (dataset) {
    const data = dataset.data;
    document.getElementsByClassName('notes')[0].innerText = dataset.description;
    const margin = {
      top: 30,
      right: 10,
      bottom: 30,
      left: 50
    };
    const w = 885 - margin.left - margin.right;
    const h = 600 - margin.top - margin.bottom;
    const barPadding = 0;
    const quarters = dataset.data.length;
    const monthName = ['January', 'February', 'March', 'April', 'May', 'June',
      'July', 'August', 'September', 'October', 'November', 'December'
    ];

    const min = d3.min(data, function (d) {
      return d[1];
    });
    const max = d3.max(data, function (d) {
      return d[1];
    });
    const minDate = new Date(data[0][0]);
    const maxDate = new Date(data[quarters - 1][0]);

    const yScale = d3.scaleLinear()
      .domain([0, max])
      .range([h, 0]);

    const yAxis = d3.axisLeft(yScale);
    const xScale = d3.scaleTime()
      .domain([minDate, maxDate])
      .range([0, w]);

    const xAxis = d3.axisBottom(xScale);

    // console.log(min, ' --> ', max)
    // console.log(yScale(min), ' --> ', yScale(max))
    // console.log(minDate, ' --> ', maxDate)

    const svg = d3.select('.barChart')
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
        let date = new Date(d[0]);
        const year = date.getFullYear();
        const month = monthName[date.getMonth()];
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
  }

  function getAjaxData (urlData, callback) {
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4) { // 4 = "DONE"
        if (xhr.status === 200) { // 200 ="OK"
          callback(JSON.parse(xhr.responseText));
        } else {
          console.log('Error: ' + xhr.status);
        }
      }
    };
    xhr.open('GET', urlData); // add false to synchronous request
    xhr.send();
  }

  return {
    init: init
  };
}());

window.addEventListener('load', app.init);
