/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const baseUrl = 'https://brusbilis.com/freecodecamp/6-backEnd/metric/v1/convert?input=';
  // const baseUrl = 'http://localhost:3000/metric/v1/convert?input='

  function init () {
    console.log('Init metric');
    document.getElementsByClassName('convertBut')[0].addEventListener('click', action);
  }

  function action (e) {
    let quest = document.getElementsByClassName('inputText')[0].value;
    let urlData = baseUrl + quest;
    getAjaxData(urlData, showResponse);
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

  function showResponse (data) {
    data = JSON.stringify(data, null, ' ');
    document.getElementById('jsonRes').innerHTML = data;
  }

  return {
    inicio: init
  };
}());

window.addEventListener('load', app.inicio);
