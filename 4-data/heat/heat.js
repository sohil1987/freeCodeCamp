(function () {
  'use strict';
  /* code here */

  var urlData =
  'https://raw.githubusercontent.com/FreeCodeCamp/ProjectReferenceData/master/global-temperature.json';
  var urlData2 = ('./data.json');
  var dataset = [];

  function init () {
    console.log('Inicio');
    getData();
  }

  function goOn () {
    var baseT = dataset.baseTemperature;
    dataset = dataset.monthlyVariance;
    // console.log(dataset.length)

    var margin = {
      top: 50,
      right: 50,
      bottom: 50,
      left: 50
    };
    var w = 960 - margin.left - margin.right;
    var h = 600 - margin.top - margin.bottom;

    // console.log(minPos, ' --> ', maxPos)
    // console.log(minTime, ' --> ', maxTime)

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
    xhr.open('GET', urlData2); // add false to synchronous request
    xhr.send();
  }

  window.addEventListener('load', init);
}());
