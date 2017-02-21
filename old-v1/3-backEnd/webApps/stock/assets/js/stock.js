var app = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  // var baseURL = 'http://localhost:3006/'
  var baseURL = 'https://brusbilis.com/freecodecamp/old-v1/webapps/';
  // var baseWs = 'ws://localhost:3006/stock/socket/'
  var baseWs = 'wss://brusbilis.com/freecodecamp/old-v1/webapps/stock/socket/';
  console.log('onload');
  var sock = new WebSocket(baseWs);
  sock.onopen = function () {
    console.log('connected to ' + baseWs);
  };
  sock.onclose = function (e) {
    console.log('connection closed (' + e.code + ')');
  };
  sock.onmessage = function (e) {
    console.log('message received: ' + e.data);
    location.reload(true);
  // app.action()
  };

  function init () {
    console.log('Inicio APP');
    // wakeUpWS()
    app.action();
  }

  function action () {
    getAjaxData(baseURL + 'stock/api/get/', function initialStatus (stockList) {
      addEvents();
      datos.names = stockList;
      datos.readData();
      views.createBoxes();
    });
  }

  function addStock () {
    var code = document.getElementById('stockCode').value;
    var url = baseURL + 'stock/api/add/?code=' + code;
    // console.log('ADDING ...', url)
    getAjaxData(url, function hola (stockList) {
      console.log('added', stockList);
      // app.action()
      // location.reload(true)
      sock.send(JSON.stringify({message: 'go'}));
    });
  }

  function delStock (e) {
    console.log('BORRAR ..', e.target.id);
    var url = app.baseURL + 'stock/api/del/?code=' + e.target.id;
    app.getAjaxData(url, function adios (stockList) {
      console.log('deleated', stockList);
      // app.action()
      // location.reload(true)
      sock.send(JSON.stringify({message: 'go'}));
    });
  }

  function wakeUpWS () {
    var wsuri = baseWs;
    console.log('onload');
    sock = new WebSocket(wsuri);
    sock.onopen = function () {
      console.log('connected to ' + wsuri);
    };
    sock.onclose = function (e) {
      console.log('connection closed (' + e.code + ')');
    };
    sock.onmessage = function (e) {
      console.log('message received: ' + e.data);
      location.reload(true);
    // app.action()
    };
  }

  function addEvents () {
    // prevent enter submit , redirect to addStock instead
    // console.log('ADD EVENTS')
    document.getElementsByClassName('addStock')[0].onkeypress = function (e) {
      var key = e.charCode || e.keyCode || 0;
      if (key == 13) {
        e.preventDefault();
        addStock();
      }
    };
    document.getElementById('addStockCode').addEventListener('click', addStock);
  }

  function getAjaxData (urlData, callback) {
    var xhr = new XMLHttpRequest();
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
    inicio: init,
    baseURL: baseURL,
    getAjaxData: getAjaxData,
    action: action,
    delStock: delStock
  };
}());

window.addEventListener('load', app.inicio);

/*

AAPL
MMM
GM
MSFT
ABT
ABBV
ATVI
ATI
ADBE
ADP

*/
