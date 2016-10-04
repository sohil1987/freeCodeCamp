(function () {
  'use strict';
  /* code here */

  var urlData =
  'https://raw.githubusercontent.com/FreeCodeCamp/ProjectReferenceData/master/GDP-data.json';
  var urlData2 = ('./data.json');
  var data = [];
  var dataset = [];

  function init () {
    console.log('Inicio');
    getData();
  }

  function goOn () {
    // console.log(data)
    //  console.log(dataset)
    document.getElementsByClassName('notes')[0].innerText = dataset.description;
    data = [5, 10, 15, 20, 25];

    d3.select('.barChart').selectAll('div')
      .data(data)
      .enter()
      .append('div')
      .attr('class', 'bar')
      .style('height', function (d) {
        return d + 'px';
      });
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
          dataset = res;
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
