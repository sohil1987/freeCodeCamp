(function () {
  'use strict';
  /* code here */

  var result = 0;
  var log = 0;
  var operation = '';

  function init () {
    console.log('Inicio');
    var clickType = document.getElementsByClassName('card');
    for (var i = 0; i < clickType.length; i++) {
      clickType.item(i).addEventListener('click', clickCal);
    // console.log(clickType.item(i).innerText)
    }
  }

  function clickCal (ev) {
    var action = ev.target.innerText;
    var prev = document.getElementById('display').innerText;
    if (action === '=') {
      printResult(result, log);
    } else if (action === 'AC') {
      log = 0;
      result = 0;
      operation = '';
      printResult(0, 0);
    } else if (action === 'CE') {
      result = 0;
      printResult(0, log);
    } else if (action === '.') {
      log += log;
      printResult(prev, log);
    } else if (Number.isNaN(parseInt(action))) {
      if (operation === '+') {
        result += result;
      }
      if (operation === '-') {}
      if (operation === '&#247') {}
      if (operation === '&#215') {}
      operation = action;
      log += log;
      printResult(prev, log);
    } else {
      console.log('NUMERO', action);
      log += log;
      printResult(result, log);
    }
  }

  function printResult (display, minidisplay) {
    document.getElementById('display').innerText = display;
    document.getElementById('minidisplay').innerText = minidisplay;
  }

  addEventListener('load', init);
}());
