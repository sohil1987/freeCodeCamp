(function () {
  'use strict';
  /* code here */

  var type = 'work';
  var status = 'off';
  var timer;
  var sound = new Audio('./images/shot.ogg');

  function init () {
    console.log('Init');
    var click = document.getElementsByClassName('action');
    for (var i = 0; i < click.length; i++) {
      click.item(i).addEventListener('click', clickAction);
    // console.log(click.item(i).id)
    }
  }

  function updateClock () {
    var now = document.getElementById('clock').innerText.split(':');
    var mins = parseInt(now[0]);
    var secs = parseInt(now[1]);
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
      var aux;
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
        document.getElementById('clock').style.color = 'green';
        document.getElementById('type').innerText = 'WORKING !';
      }
    }
  }

  function clock () {
    if (status === 'off') {
      status = 'on';
      timer = setInterval(updateClock, 1000);
    } else {
      status = 'off';
      clearInterval(timer);
    }
    console.log('CLOCK', status);
  }

  function changeTimer (id, inc) {
    var time = parseInt(document.getElementById(id).innerText);
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
  }

  function clickAction (ev) {
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
        case 'clock':
          clock();
          break;
      }
    } else if (status === 'on' && ev.target.id === 'clock') {
      console.log('STATUs ON');
      clock();
    }
  }

  addEventListener('load', init);
}());
