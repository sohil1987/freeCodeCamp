/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const baseUrl = 'https://brusbilis.com/freecodecamp/6-backEnd/message/v1/';
  // const baseUrl = 'http://localhost:3000/message/v1/'

  function init () {
    console.log('Init message');
    let actions = document.getElementsByClassName('actionBut');
    for (let i = 0; i < actions.length; i++) {
      actions[i].addEventListener('click', action);
    }
    makeAjaxRequest(baseUrl, 'GET', getThreads);
  // makeAjaxRequest(baseUrl + 1, 'GET', getReplies)
  }

  function action (e) {
    let urlData = baseUrl;
    switch (e.target.id) {
      case 'b1':
        // console.log('b1')
        // console.log(urlData)
        makeAjaxRequest(urlData, 'GET', getThreads);
        break;
      case 'b2':
        // console.log('b2')
        let t21 = document.getElementById('t21').value;
        let t22 = document.getElementById('t22').value;
        let t23 = document.getElementById('t23').value;
        if (t21 === '' || t22 === '') {
          return;
        }
        urlData += '?title=' + t21 + '&text=' + t22;
        if (t23 !== '') {
          urlData += '&password=' + t23;
        }
        // console.log(urlData)
        makeAjaxRequest(urlData, 'POST', createThread);
        break;
      case 'b3':
        // console.log('b3')
        let t3 = document.getElementById('t3').value;
        if (t3 === '' || t3 < 1 || isNaN(t3)) {
          return;
        }
        urlData += '?board=' + t3;
        // console.log(urlData)
        makeAjaxRequest(urlData, 'PUT', reportThread);
        break;
      case 'b4':
        // console.log('b4')
        let t41 = document.getElementById('t41').value;
        let t42 = document.getElementById('t42').value;
        if (t41 === '' || t41 < 1 || isNaN(t41)) {
          return;
        }
        urlData += '?board=' + t41;
        if (t42 !== '') {
          urlData += '&password=' + t42;
        }
        // console.log(urlData)
        makeAjaxRequest(urlData, 'DELETE', deleteThread);
        break;
      case 'b5':
        // console.log('b5')
        let t5 = document.getElementById('t5').value;
        if (t5 === '' || t5 < 0 || isNaN(t5)) {
          return;
        }
        urlData += t5;
        // console.log(urlData)
        makeAjaxRequest(urlData, 'GET', getReplies);
        break;
      case 'b6':
        // console.log('b6')
        let t61 = document.getElementById('t61').value;
        let t62 = document.getElementById('t62').value;
        let t63 = document.getElementById('t63').value;
        if (t61 === '' || t61 < 0 || isNaN(t61) || t62 === '') {
          return;
        }
        urlData += t61 + '?text=' + t62;
        if (t63 !== '') {
          urlData += '&password=' + t63;
        }
        // console.log(urlData)
        makeAjaxRequest(urlData, 'POST', createReply);
        break;
      case 'b7':
        // console.log('b7')
        let t7 = document.getElementById('t7').value;
        if (t7 === '' || t7 < 1 || isNaN(t7)) {
          return;
        }
        urlData += t7;
        // console.log(urlData)
        makeAjaxRequest(urlData, 'PUT', reportReply);
        break;
      case 'b8':
        let t81 = document.getElementById('t81').value;
        let t82 = document.getElementById('t82').value;
        if (t81 === '' || t81 < 1 || isNaN(t81)) {
          return;
        }
        urlData += t81;
        if (t82 !== '') {
          urlData += '?password=' + t82;
        }
        // console.log(urlData)
        makeAjaxRequest(urlData, 'DELETE', deleteReply);
        // console.log('b8')
        break;
    }
  }

  function getThreads (data) {
    let html = '<div class="threads">';
    for (let i = 0; i < data.length; i++) {
      html +=
        `<div class="thread">
<span>{board} : ${data[i].board} , {title} : ${data[i].title}</span>
</div>`;
    }
    html += '</div>';
    showResponse(html, 'jsonRes1');
  }
  function createThread (data) {
    showResponse(data, 'jsonRes2');
  }
  function reportThread (data) {
    showResponse(data, 'jsonRes3');
  }
  function deleteThread (data) {
    showResponse(data, 'jsonRes4');
  }

  function getReplies (data) {
    if (data.error !== undefined) {
      let html = data.error;
      showResponse(html, 'jsonRes5');
      return;
    }
    let status = 'Not reported';
    if (data.isReported) {
      status = 'Reported';
    }
    let html = '<div class="thread">';
    html +=
      `<div class="replies">
<span><b>Board</b> :  ${data.board} - ${data.title} (${status})</span><br>
<span>${data.text}</span><br>
<span><b>Created on:</b> ${data.creationDate}<b> - Last Updated:</b> ${data.lastPost}</span>
</div>`;
    for (let i = 0; i < data.replies.length; i++) {
      let status = 'Not reported';
      if (data.replies[i].isReported) {
        status = 'reported';
      }
      html += `<div class="reply">
      <span>Reply id : ${data.replies[i].id} - ${status}</span><br>
      <span>${data.replies[i].text}</span>
      </div>`;
    }
    html += '</div>';
    showResponse(html, 'jsonRes5');
  }

  function createReply (data) {
    showResponse(data, 'jsonRes6');
  }
  function reportReply (data) {
    showResponse(data, 'jsonRes7');
  }
  function deleteReply (data) {
    showResponse(data, 'jsonRes8');
  }

  function showResponse (data, place) {
    if (data.info !== '' && data.info !== undefined) {
      document.getElementById(place).innerHTML = data.info;
      return;
    }
    if (data.error !== '' && data.error !== undefined) {
      document.getElementById(place).innerHTML = data.error;
      return;
    }
    document.getElementById(place).innerHTML = data;
  }

  function makeAjaxRequest (url, action, callback) {
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4) { // 4 = "DONE"
        if (xhr.status === 200) { // 200 ="OK"
          const data = JSON.parse(xhr.responseText);
          callback(data);
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

  return {
    inicio: init
  };
}());

window.addEventListener('load', app.inicio);
