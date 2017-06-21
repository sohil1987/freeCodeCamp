/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const urlData = 'https://fcctop100.herokuapp.com/api/fccusers/top/recent';
  const urlData2 = 'https://fcctop100.herokuapp.com/api/fccusers/top/alltime';
  let recentData = [];
  let allData = [];
  const arrow = './../../assets/images/arrowDown.png';

  function init () {
    console.log('Init Camper LeaderBoard');
    getData(urlData);
  }

  function activeOrderByClicks () {
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
    let orderedByLast = '';
    let orderedByAll = '';
    if (dataSet === recentData) {
      orderedByLast = `<img class="logo" src="${arrow}">`;
      orderedByAll = '';
    } else if (dataSet === allData) {
      orderedByAll = `<img class="logo" src="${arrow}">`;
      orderedByLast = '';
    }
    let res = `      
    <table class="table">
        <thead class="thead">
          <tr>
            <th class="">#</th>
            <th class="name">Camper Name</th>
            <th id="theadLast" class="date">
              <div>Last Month</div>${orderedByLast}
            </th>
            <th id="theadAll" class="date">
              <div>All Time</div>${orderedByAll}
            </th>
          </tr>
        </thead>
        <tbody>`;

    for (let i = 0; i < dataSet.length; i++) {
      const linkName = '<a href="https://freecodecamp.com/' + dataSet[i].username + '">' + dataSet[i].username + '</a>';
      res +=
        `<tr>
          <th class="rowNumber" scope="row">${i + 1}
          </th>
          <td class="camper">
            <img class="logo" src="${dataSet[i].img}">
            &nbsp;<div>${linkName}</div>
          </td>
          <td class="points">${dataSet[i].recent}</td>
          <td class="points">${dataSet[i].alltime}</td>
        </tr>`;
    }
    res += `
        </tbody>
        </table>`;
    document.getElementById('dataTable').innerHTML = res;
    // theadLast.style.class = ''
    if (dataSet === recentData) {
      theadAll.classList.remove('date');
      theadLast.className = ('date');
    } else if (dataSet === allData) {
      theadLast.classList.remove('date');
      theadAll.className = ('date');
    }
    // each population events listener disappears so we renove them
    activeOrderByClicks();
  }

  function getData (url) {
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      const DONE = 4;
      const OK = 200;
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

  return {
    init: init
  };
}());

window.addEventListener('load', app.init);
