/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const baseUrl = 'https://brusbilis.com/freecodecamp/6-backEnd/library/v1/books/';
  // const baseUrl = 'http://localhost:3000/library/v1/books/'

  function init () {
    console.log('Init metric');
    let actions = document.getElementsByClassName('actionBut');
    for (let i = 0; i < actions.length; i++) {
      actions[i].addEventListener('click', action);
    }
  }

  function action (e) {
    let urlData = baseUrl;
    switch (e.target.id) {
      case 'b1':
        // console.log(urlData)
        makeAjaxRequest(urlData, 'GET', showResponse1);
        break;
      case 'b2':
        let input2 = document.getElementById('t2').value;
        if (input2 !== '') { // && input >= 0) {
          urlData += input2;
          // console.log(urlData)
          makeAjaxRequest(urlData, 'GET', showResponse2);
        }
        break;
      case 'b3':
        let input3 = document.getElementById('t3').value;
        if (input3 !== '') { // && input >= 0) {
          urlData += '?title=' + input3;
          // console.log(urlData)
          makeAjaxRequest(urlData, 'POST', none);
        }
        break;
      case 'b4':
        let input41 = document.getElementById('t41').value;
        let input42 = document.getElementById('t42').value;

        if (input41 !== '' && input42 !== '') {
          urlData += + input41 + '?comment=' + input42;
          // console.log(urlData)
          makeAjaxRequest(urlData, 'POST', none);
        }
        break;
      case 'b5':
        // console.log(urlData)
        makeAjaxRequest(urlData, 'DELETE', none);
        break;
      case 'b6':
        let input6 = document.getElementById('t6').value;
        if (input6 > 0) { // && input >= 0) {
          urlData += input6;
          // console.log(urlData)
          makeAjaxRequest(urlData, 'DELETE', none);
        }
        break;
    }
  }

  function makeAjaxRequest (url, action, callback) {
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4) { // 4 = "DONE"
        if (xhr.status === 200) { // 200 ="OK"
          if (action === 'GET') {
            const data = JSON.parse(xhr.responseText);
            callback(data);
          } else {
            callback();
          }
        } else {
          console.log('Error: ' + xhr.status);
        }
      }
    };
    xhr.open(action, url);
    if (action === 'GET') {
      xhr.send();
    } else if (action !== 'GET') {
      xhr.setRequestHeader('Content-Type',
        'application/x-www-form-urlencoded; charset=UTF-8');
      // xhr.send(params)
      xhr.send();
    }
  }

  function none () {
    // console.log('none')
  }
  function showResponse1 (data) {
    showResponse(data, 'jsonRes1');
  }
  function showResponse2 (data) {
    showResponse(data, 'jsonRes2');
  }

  function showResponse (data, place) {
    data = JSON.stringify(data, null, ' ');
    document.getElementById(place).innerHTML = data;
  }

  return {
    inicio: init
  };
}());

window.addEventListener('load', app.inicio);
