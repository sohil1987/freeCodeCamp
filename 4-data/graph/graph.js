(function () {
  'use strict';
  /* code here */
  // v3
  // https://github.com/alignedleft/d3-book
  // http://bl.ocks.org/sathomas/11550728

  var urlData =
  'https://raw.githubusercontent.com/DealPete/forceDirected/master/countries.json';
  var dataset = [];

  function init () {
    console.log('Inicio');
    getData();
  }

  function goOn () {
    var margin = {
      top: 80,
      right: 50,
      bottom: 20,
      left: 50
    };
    var w = 960 - margin.left - margin.right;
    var h = 800 - margin.top - margin.bottom;

    var svg = d3.select('.forceDirectedSvg')
      .attr('width', w + margin.left + margin.right)
      .attr('height', h + margin.top + margin.bottom);

    var force = d3.layout.force()
      .nodes(dataset.nodes)
      .links(dataset.links)
      .size([w, h])
      .linkDistance([50])
      .charge([-100])
      .start();

    var borders = svg.selectAll('borders')
      .data(dataset.links)
      .enter()
      .append('line')
      .attr('class', 'boundary');

    var countries = svg.select('.countries')
      .data(dataset.nodes)
      .enter()
      .append('img')
      .attr('class', function (d) {
        // console.log(d.code)
        return 'flag flag-' + d.code;
      })
      .call(force.drag)
      .on('mouseover', function (d) {
        // d3.select(this)
        //  .style('fill', 'white')
        d3.select('#tooltip')
          .style('left', (d3.event.pageX - 30) + 'px')
          .style('top', (d3.event.pageY - 50) + 'px')
          .select('#country')
          .text(d.country);
        d3.select('#tooltip').classed('hidden', false);
      })

      .on('mouseout', function (d) {
        d3.select('#tooltip').classed('hidden', true);
      });

    force.on('tick', function () {
      borders
        .attr('x1', function (d) {
          return d.source.x + margin.left;
        })
        .attr('y1', function (d) {
          return d.source.y + margin.top;
        })
        .attr('x2', function (d) {
          return d.target.x + margin.left;
        })
        .attr('y2', function (d) {
          return d.target.y + margin.top;
        });

      countries
        .style('left', function (d) {
          return d.x + margin.left + 8 + 'px';
        })
        .style('top', function (d) {
          return d.y + margin.top - 6 + 'px';
        });
    });

    console.log('End');
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
