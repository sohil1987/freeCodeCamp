(function () {
  'use strict';
  /* code here */
  /* example https://bl.ocks.org/d3noob/bdf28027e0ce70bd132edc64f1dd7ea4 */

  var urlData =
  'https://raw.githubusercontent.com/FreeCodeCamp/ProjectReferenceData/master/GDP-data.json';
  var urlData2 = ('./data.json');
  var data = [];
  // var data = [4, 8, 15, 16, 23, 42, 80]
  var data2 = [];

  function init () {
    console.log('Inicio');
    getData();
  // goOn()
  }

  function goOn () {
    // set the dimensions and margins of the graph
    var margin = {
      top: 40,
      right: 40,
      bottom: 40,
      left: 40
    };
    var width = 960 - margin.left - margin.right;
    var height = 500 - margin.top - margin.bottom;
    var minDate = new Date(data2.data[0][0]);
    var maxDate = new Date(data2.data[data2.data.length - 1][0]);
    var barSize = Math.ceil(width / data.length);
    var dateParser = d3.timeParse('%Y-%m-%d');

    console.log(barSize);
    document.getElementsByClassName('notes')[0].innerText = data2.description;

    // console.log(data)
    console.log(minDate, maxDate);

    // set the ranges
    var x = d3.scaleBand()
      .range([0, width])
      .padding(0.1)
      .domain(data.map(function (d) {
        return d;
      }));

    var x2 = d3.scaleTime()
      .domain([minDate, maxDate])
      .range([0, width]);

    var x3 = d3.scaleTime()
      .domain(d3.extent(data2.data, function (d) {
        return dateParser(d[0]);
      }))
      .range([0, width]);

    var y = d3.scaleLinear()
      .range([height, 0])
      .domain([0, d3.max(data, function (d) {
        return d;
      })]);

    var svg = d3.select('.barChart')
      .attr('width', width + margin.left + margin.right)
      .attr('height', height + margin.top + margin.bottom)
      .append('g')
      .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')');

    // append the rectangles for the bar chart
    svg.selectAll('.bar')
      .data(data)
      .enter()
      .append('rect')
      .attr('class', 'bar')
      .attr('x', function (d, i) {
        return x(d);
      })
      .attr('width', barSize) // x.bandwidth())
      .attr('y', function (d) {
        return y(d);
      })
      .attr('height', function (d) {
        return height - y(d);
      });

    // add the x Axis
    svg.append('g')
      .attr('transform', 'translate(0,' + height + ')')
      .call(d3.axisBottom(x));

    // add the y Axis
    svg.append('g')
      .call(d3.axisLeft(y));
  }

  function getData () {
    console.log('Getting data ...');
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      var DONE = 4;
      var OK = 200;
      if (xhr.readyState === DONE) {
        if (xhr.status === OK) {
          var res = JSON.parse(xhr.responseText);
          // console.log(res.data.length)
          for (var i = 0; i < res.data.length; i++) {
            // if (i % 32 === 0)
            data.push(res.data[i][1]);
          // data.push(res.data[i])
          }
          data2 = res;
          goOn();
        } else {
          console.log('Error: ' + xhr.status);
        }
      }
    };
    xhr.open('GET', urlData2); // add false to synchronous request
    xhr.send();
  }

  window.addEventListener('load', init);
}());
