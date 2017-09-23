/* jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  // const baseURL = 'http://localhost:3000/'
  const baseURL = 'https://brusbilis.com/freecodecamp/7-bonus2/';
  // const baseWs = 'ws://localhost:3000/stock/socket/'
  const baseWs = 'wss://brusbilis.com/freecodecamp/7-bonus2/stock/socket/';
  console.log('onload');
  let sock = new WebSocket(baseWs);
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
    const code = document.getElementById('stockCode').value;
    let url = baseURL + 'stock/api/add/?code=' + code;
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
    let url = app.baseURL + 'stock/api/del/?code=' + e.target.id;
    app.getAjaxData(url, function adios (stockList) {
      console.log('deleated', stockList);
      // app.action()
      // location.reload(true)
      sock.send(JSON.stringify({message: 'go'}));
    });
  }

  function wakeUpWS () {
    let wsuri = baseWs;
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
      let key = e.charCode || e.keyCode || 0;
      if (key == 13) {
        e.preventDefault();
        addStock();
      }
    };
    document.getElementById('addStockCode').addEventListener('click', addStock);
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
