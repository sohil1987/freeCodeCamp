(function () {
  'use strict';
  /* code here */

  var urlData = 'https://fcctop100.herokuapp.com/api/fccusers/top/recent';
  var urlData2 = 'https://fcctop100.herokuapp.com/api/fccusers/top/alltime';
  var recentData = [];
  var allData = [];

  function init () {
    console.log('Inicio');
    getData(urlData);
  }

  function goOn () {
    document.getElementById('theadLast').addEventListener('click',
      function () {
        populateTable(recentData);
      });
    document.getElementById('theadAll').addEventListener('click',
      function () {
        populateTable(allData);
      });
  }

  function populateTable (dataSet) {
    var orderedByLast = '';
    var orderedByAll = '';
    if (dataSet === recentData) {
      orderedByLast = '<img class="arrow" src="./images/arrowDown20x20.png">';
    } else if (dataSet === allData) {
      orderedByAll = '<img class="arrow" src="./images/arrowDown20x20.png">';
    }

    var table =
    '<table class="table table-sm table-striped table-bordered text-xs-center">';
    var theadPos =
    '<thead class="thead-inverse"><tr><th class="col-xs-1 text-xs-center">#</th>';
    var theadName = '<th class="col-xs-5 text-xs-center">Camper Name</th>';
    var theadLast =
    '<th id="theadLast" class="col-xs-3 text-xs-center">Last Month ' +
      orderedByLast + '</th>';
    var theadAll =
    '<th id="theadAll" class="col-xs-3 text-xs-center">All time ' +
      orderedByAll + '</th></tr></thead><tbody>';
    var res = table + theadPos + theadName + theadLast + theadAll;

    for (var i = 0; i < dataSet.length; i++) {
      var linkName =
      '<a href="https://freecodecamp.com/' + dataSet[i].username + '">' +
        dataSet[i].username + '</a>';
      res += '<tr><th class="col-xs-1 text-xs-center" scope="row">' + (i + 1) +
        '</th>';
      res += '<td class="col-xs-5 text-xs-left"><img class="logo" src="' +
        dataSet[i].img +
        '"> ';
      res += ' &nbsp; ' + linkName + '</td>';
      res += '<td class="col-xs-3">' + dataSet[i].recent + '</td>';
      res += '<td class="col-xs-3">' + dataSet[i].alltime + '</td></tr>';
    }
    res += '</tbody>';
    document.getElementById('dataTable').innerHTML = res;
    // each population events listener disappears so we renove them
    goOn();
  }

  function getData (url) {
    // console.log('Getting data ...')
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      var DONE = 4;
      var OK = 200;
      if (xhr.readyState === DONE) {
        if (xhr.status === OK) {
          if (url === urlData) {
            recentData = JSON.parse(xhr.responseText);
            getData(urlData2);
          }
          if (url === urlData2) {
            allData = JSON.parse(xhr.responseText);
            populateTable(recentData);
          }
        } else {
          console.log('Error: ' + xhr.status);
        }
      }
    };
    xhr.open('GET', url); // add false to synchronous request
    xhr.send();
  }

  window.addEventListener('load', init);
}());
