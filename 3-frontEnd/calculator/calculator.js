/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  let result = 0;
  let log = 0;
  let operation = '';
  const digitLimit = 17;

  function init () {
    console.log('Init Javascript Calculator');
    let clickType = document.getElementsByClassName('action');
    for (let i = 0; i < clickType.length; i++) {
      clickType.item(i).addEventListener('click', clickAction);
    // console.log(clickType.item(i).innerText)
    }
  }

  function clickAction (e) {
    let action = e.target.innerText;
    let prev = document.getElementById('result').innerText;
    if (prev.length > digitLimit) {
      tooManyDigits();
    } else if (action === '=') {
      makeOperation(action);
    } else if (action === 'AC') {
      makeAllClear();
    } else if (action === 'CE') {
      makeClear();
    } else if (action === '.') {
      if (prev[prev.length - 1] !== '.') { // avoid 2+ commas in a row
        addComma(prev);
      }
    } else if (Number.isNaN(parseFloat(action))) {
      if (!Number.isNaN(parseInt(log[log.length - 1]))) { // dont repeat ops
        if (result === 0 && action === '-') { // negative number begin with -
          addNumber(action, prev);
        } else if (action === '=') {
          makeOperation(action);
        } else {
          addNumber(action, prev);
        }
      }
    } else {
      if (action !== '0' || result !== 0) { // avoid leading zeros
        addNumber(action, prev);
      }
    }
  }

  function tooManyDigits () {
    result = 0;
    log = 0;
    operation = '';
    printResult(result, 'TOO MANY DIGITS');
  }

  function makeOperation (operation) {
    log = log.replace('\u00F7', '/');
    log = log.replace('\u00D7', '*');
    printResult(eval(log), log);
  }

  function addComma (prev) {
    result += '.';
    log += '.';
    printResult(result, log);
  }

  function addNumber (num, prev) {
    if (Number.isNaN(parseFloat(num))) {
      result = num;
    } else if (prev !== '0') {
      result = prev + num;
    } else result = num;
    if (log !== 0) {
      log += num;
    } else log = num;
    printResult(result, log);
  }

  function makeAllClear () {
    result = 0;
    log = 0;
    operation = '';
    printResult(result, log);
  }

  function makeClear () {
    result = 0;
    printResult(result, log);
  }

  function printResult (result, log) {
    if (String(result).length < digitLimit + 1) {
      document.getElementById('result').innerText = result;
      document.getElementById('log').innerText = log;
    } else {
      tooManyDigits();
    }
  }

  return {
    init: init
  };
}());

window.addEventListener('load', app.init);
