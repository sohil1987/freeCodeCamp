/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  let type = 'work';
  let status = 'off';
  let timer;
  let sound = new Audio('./../../assets/sounds/clock/wakeup.mp3');

  function init () {
    console.log('Init Pomodoro Clock');
    let click = document.getElementsByClassName('action');
    for (let i = 0; i < click.length; i++) {
      click.item(i).addEventListener('click', clickAction);
    // console.log(click.item(i).id)
    }
  }

  function updateClock () {
    let now = document.getElementById('clock').innerText.split(':');
    let mins = parseInt(now[0]);
    let secs = parseInt(now[1]);
    if (secs === 0) {
      mins--;
      secs = 59;
    } else if (secs <= 10) {
      secs--;
      secs = '0' + secs;
    } else {
      secs--;
    }
    if (mins < 10) {
      mins = '0' + mins;
    }
    // console.log(mins, secs)
    now = mins + ':' + secs;
    document.getElementById('clock').innerText = now;
    if (now === '00:00') {
      sound.play();
      console.log('END');
      let aux;
      if (type === 'work') {
        type = 'break';
        aux = document.getElementById('breakTime').innerText;
        if (aux < 10) {
          aux = '0' + aux;
        }
        aux = aux + ':00';
        document.getElementById('clock').innerText = aux;
        document.getElementById('clock').style.color = 'red';
        document.getElementById('type').innerText = 'RESTING !';
      } else {
        type = 'work';
        aux = document.getElementById('workTime').innerText;
        if (aux < 10) {
          aux = '0' + aux;
        }
        aux = aux + ':00';
        document.getElementById('clock').innerText = aux;
        document.getElementById('clock').style.color = '#333333';
        document.getElementById('type').innerText = 'WORKING !';
      }
    }
  }

  function maintainSize () {
    let now = document.getElementById('clock').innerText.split(':');
    let mins = parseInt(now[0]);
    let secs = parseInt(now[1]);
    if (mins < 10) {
      mins = '0' + mins;
    }
    if (secs === 0) {
      secs = '00';
    }
    // console.log(mins, secs)
    now = mins + ':' + secs;
    document.getElementById('clock').innerText = now;
  }

  function clock () {
    if (status === 'off') {
      status = 'on';
      document.getElementById('type').innerText = 'WORKING !';
      timer = setInterval(updateClock, 1000);
    } else {
      status = 'off';
      clearInterval(timer);
    }
    console.log('CLOCK', status);
  }

  function changeTimer (id, inc) {
    let time = parseInt(document.getElementById(id).innerText);
    if (id === 'breakTime') {
      if (time < 30 && inc > 0) time++;
      if (time > 1 && inc < 0) time--;
    }
    if (id === 'workTime') {
      if (time > 1 && inc < 0) time--;
      if (time < 99 && inc > 0) time++;
      // if (time < 10) time = '0' + time
      document.getElementById('clock').innerText = time + ':00';
    }
    document.getElementById(id).innerText = time;
    maintainSize();
  }

  function clickAction (ev) {
    console.log('Pressed > ', ev.target.id);
    if (ev.target.id === 'reload') {
      location.reload();
    }
    if (status === 'off') {
      switch (ev.target.id) {
        case 'breakTimeMinus':
          changeTimer('breakTime', -1);
          break;
        case 'breakTimePlus':
          changeTimer('breakTime', +1);
          break;
        case 'workTimeMinus':
          changeTimer('workTime', -1);
          break;
        case 'workTimePlus':
          changeTimer('workTime', +1);
          break;
        case 'start':
          clock();
          break;
      }
    } else if (status === 'on' && ev.target.id === 'start') {
      clock();
    }
  }

  return {
    init: init
  };
}());

addEventListener('load', app.init);
