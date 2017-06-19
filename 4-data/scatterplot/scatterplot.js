/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const url = 'https://raw.githubusercontent.com/freeCodeCamp/ProjectReferenceData/master/cyclist-data.json';
  const urlFile = './scatterplot.json';

  function init () {
    console.log('Init Scatterplot Graph');
    getAjaxData(url, doIt);
  // getAjaxData(urlFile, doIt)
  }

  function doIt (dataset) {
    console.log(dataset);
    const margin = {
      top: 30,
      right: 130,
      bottom: 70,
      left: 70
    };
    const w = 960 - margin.left - margin.right;
    const h = 600 - margin.top - margin.bottom;

    const minPos = dataset[0].Place;
    const maxPos = dataset[dataset.length - 1].Place;
    const minTime = new Date(dataset[0].Seconds * 1000 - 5000);
    const maxTime = new Date(dataset[dataset.length - 1].Seconds * 1000 + 5000);

    // console.log(minPos, ' --> ', maxPos)
    // console.log(minTime, ' --> ', maxTime)

    const yScale = d3.scaleLinear()
      .domain([minPos - 1, maxPos + 1])
      .range([0, h]);

    const yAxis = d3.axisLeft(yScale);

    const xScale = d3.scaleTime()
      .domain([minTime, maxTime])
      .range([w, 0]);

    const xAxis = d3.axisBottom(xScale)
      .tickFormat(d3.timeFormat('%M:%S'));

    const svg = d3.select('.scatterplotGraph')
      .attr('width', w + margin.left + margin.right)
      .attr('height', h + margin.top + margin.bottom)
      .attr('class', 'scatterplotGraph')
      // .style('background', 'red')
      .append('g')
      .attr('transform', 'translate(' + margin.left + ', ' + margin.top + ')');

    svg.selectAll('circle')
      .data(dataset)
      .enter()
      .append('circle')
      .attr('cx', function (d) {
        return xScale(d.Seconds * 1000);
      })
      .attr('cy', function (d) {
        return yScale(d.Place);
      })
      .attr('r', 5)
      .attr('fill', function (d) {
        if (d.Doping === '') return 'green';
        if (d.Doping !== '') return 'red';
      })
      .on('mouseover', function (d) {
        // d3.select(this)
        // .style('fill', 'white')
        const value1 = d.Name + ' (' + d.Nationality + ')';
        const value2 = d.Year + ' -- ' + d.Time + ' min';
        const value3 = d.Doping;
        d3.select('#tooltip')
          .select('#value1')
          .text(value1);
        d3.select('#tooltip')
          .style('left', (0.5 * w + margin.left + margin.right) + 'px')
          .style('top', (2 / 3 * h + margin.top + margin.bottom) + 'px')
          .select('#value2')
          .text(value2);
        d3.select('#tooltip')
          .select('#value3')
          .text(value3);
        d3.select('#tooltip').classed('hidden', false);
      })

      .on('mouseout', function (d) {
        d3.select('#tooltip').classed('hidden', true);
      });

    svg.selectAll('text')
      .data(dataset)
      .enter()
      .append('text')
      .text(function (d) {
        return d.Name;
      })
      .attr('x', function (d) {
        return xScale(d.Seconds * 1000) + 7;
      })
      .attr('y', function (d) {
        return yScale(d.Place) + 4;
      })
      .attr('class', 'name');

    svg.append('g')
      .attr('class', 'axis')
      .call(yAxis);

    // yAxis label
    svg.append('text')
      .attr('transform', 'translate(' + (-35) + ', ' + (30) + ')')
      .style('text-anchor', 'middle')
      .style('font-weight', 'bold')
      .text('Place');

    svg.append('g')
      .attr('transform', 'translate(0,' + h + ')')
      .attr('class', 'axis')
      .call(xAxis);

    // xAxis label
    svg.append('text')
      .attr('transform', 'translate(' + (w / 2) + ', ' + (h + 40) + ')')
      .style('text-anchor', 'middle')
      .style('font-weight', 'bold')
      .text('Time (minutes)');

    // legends
    svg.append('g')
      .append('circle')
      .attr('cx', 0.8 * w)
      .attr('cy', 0.9 * h)
      .attr('r', 7)
      .attr('fill', 'red');
    svg.append('g')
      .attr('class', 'legend')
      .append('text')
      .attr('x', 0.8 * w + 10)
      .attr('y', 0.9 * h + 4)
      .text('Doping Allegations');

    svg.append('g')
      .append('circle')
      .attr('cx', 0.8 * w)
      .attr('cy', 0.9 * h + 20)
      .attr('r', 7)
      .attr('fill', 'green');
    svg.append('g')
      .attr('class', 'legend')
      .append('text')
      .attr('x', 0.8 * w + 10)
      .attr('y', 0.9 * h + 24)
      .text(' NoDoping Allegations');

    console.log('End');
  }

  function getColor (color) {
    console.log(color);
    if (color === '') return 'green';
    if (color !== '') return 'red';
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
