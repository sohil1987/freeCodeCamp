/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const baseUrl = 'https://brusbilis.com/freecodecamp/6-backEnd/issue/v1/test1/';
  // const baseUrl = 'http://localhost:3000/issue/v1/test1/'

  function init () {
    console.log('Init issue');
    let actions = document.getElementsByClassName('actionBut');
    for (let i = 0; i < actions.length; i++) {
      actions[i].addEventListener('click', action);
    }
  // makeAjaxRequest(baseUrl, 'GET', getList)
  // makeAjaxRequest(baseUrl + 0, 'GET', getOne)
  }

  function action (e) {
    let urlData = baseUrl;
    switch (e.target.id) {
      case 'b1':
        // console.log('List all issues')
        makeAjaxRequest(urlData, 'GET', getList);
        break;
      case 'b2':
        // console.log('List an issue')
        let t2 = document.getElementById('t2').value;
        urlData += t2;
        if (t2 !== '' && t2 >= -1) {
          makeAjaxRequest(urlData, 'GET', getOne);
        }
        break;
      case 'b3':
        // console.log('Create Issue')
        let t31 = document.getElementById('t31').value;
        let t32 = document.getElementById('t32').value;
        let t33 = document.getElementById('t33').value;
        let t34 = document.getElementById('t34').value;
        if (t31 === '' || t32 === '' || t33 === '') {
          return;
        }
        urlData += '?title=' + t31;
        urlData += '&text=' + t32;
        urlData += '&createdBy=' + t33;
        if (t34 !== '') {
          urlData += '&assignedTo=' + t34;
        }
        console.log(urlData);
        makeAjaxRequest(urlData, 'POST', createOne);
        break;
      case 'b4':
        // console.log('Update an issue')
        let t41 = document.getElementById('t41').value;
        let t42 = document.getElementById('t42').value;
        let t43 = document.getElementById('t43').value;
        let t44 = document.getElementById('t44').value;
        let t45 = document.getElementById('t45').value;
        let close = document.getElementById('close').checked;
        if (t45 === '' || t45 < 0) {
          return;
        }
        if (t45 !== '') {
          urlData += '?issueID=' + t45;
        }
        if (t41 !== '') {
          urlData += '&title=' + t41;
        }
        if (t41 !== '') {
          urlData += '&text=' + t42;
        }
        if (t43 !== '') {
          urlData += '&createdBy=' + t43;
        }
        if (t44 !== '') {
          urlData += '&assignedTo=' + t44;
        }
        if (close) {
          urlData += '&close=true';
        }
        makeAjaxRequest(urlData, 'PUT', updateOne);
        break;
      case 'b5':
        // console.log('Delete an issue')
        let t5 = document.getElementById('t5').value;
        if (t5 === '' || t5 < 0) {
          return;
        }
        urlData += '?id=' + t5;
        makeAjaxRequest(urlData, 'DELETE', deleteOne);
        break;
    }
  }

  function getList (data) {
    let html = '<div class="issues">';
    for (let i = 0; i < data.issues.length; i++) {
      html +=
        `<div class="issue">
<span>{id} : ${data.issues[i].id} , {title} : ${data.issues[i].title}</span>
</div>`;
    }
    html += '</div>';
    showResponse(html, 'jsonRes1');
  }

  function getOne (data) {
    if (data.error !== undefined) {
      let html = data.error;
      showResponse(html, 'jsonRes2');
      return;
    }
    let status = 'closed';
    if (data.isOpen) {
      status = 'open';
    }
    let html = '<div class="issues">';
    html +=
      `<div class="issue">
<span> ${data.id} - ${data.title} (${status})</span><br>
<span>${data.text}</span><br>
<span><b>Created by:</b> ${data.author}<b> - Assigned to:</b> ${data.assignee}</span><br>
<span><b>Created on:</b> ${data.creationDate}<b> - Last Updated:</b> ${data.latestUpdate}</span><br>
</div>`;
    html += '</div>';
    showResponse(html, 'jsonRes2');
  }

  function createOne (data) {
    showResponse(data.info, 'jsonRes3');
  }

  function updateOne (data) {
    if (data.info !== '') {
      showResponse(data.info, 'jsonRes4');
      return;
    }
    showResponse(data.error, 'jsonRes4');
  }

  function deleteOne (data) {
    if (data.info !== '') {
      showResponse(data.info, 'jsonRes5');
      return;
    }
    showResponse(data.error, 'jsonRes5');
  }

  function getStatus (isOpen) { // creo que sobra, se hace en el servidor ??
    if (isOpen) {
      return 'open';
    }
    return 'closed';
  }

  function showResponse (data, place) {
    // data = JSON.stringify(data, null, ' ')
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
