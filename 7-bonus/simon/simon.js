(function simon () {
  'use strict';
  /* code here */

  var strict = 'off';
  var start = 'off';
  var pos = ['green', 'red', 'yellow', 'blue'];
  var colors = ['green', '#b20000', '#cccc00', 'blue'];
  var colors2 = ['chartreuse', '#ff0000', '#ffff00', 'cyan'];
  var cpu = []; // ['green', '#b20000', '#cccc00', 'blue']
  var cpuMoves = []; // [0, 1, 2, 3]
  var playerMoves = [];
  var counter = 0;
  var round = 0;
  var audio1 = new Audio('./sounds/simon1.mp3');
  var audio2 = new Audio('./sounds/simon2.mp3');
  var audio3 = new Audio('./sounds/simon3.mp3');
  var audio4 = new Audio('./sounds/simon4.mp3');
  var audio = [audio1, audio2, audio3, audio4];

  function init () {
    console.log('Init');
    document.getElementById('strict').addEventListener('click', changeStatus);
    document.getElementById('start').addEventListener('click', changeStatus);
    document.getElementById('restart').addEventListener('click', reload);
  }

  function startRound (restart) {
    if (!restart) {
      addMovement();
    }
    animate(function (valor) {
      upBoard();
    });
  }

  function userResponse (ev) {
    turnOn(pos.indexOf(ev.target.id), 100);
    playerMoves.push(pos.indexOf(ev.target.id));
    document.getElementById('counter').textContent = playerMoves.length + 1;
    document.getElementById('round').textContent = cpuMoves.length;
    var success = playerMoves.every(function (element, index) {
      return element === cpuMoves[index];
    });
    console.log(success);
    if (!success) {
      playerMoves = [];
      downBoard();
      if (strict === 'off') {
        startRound(true);
      }
      if (strict === 'on') {
        location.reload();
      }
    }
    if (success && playerMoves.length === cpuMoves.length) {
      if (playerMoves.length >= 20) {
        alert('YOU WIN!');
        location.reload();
      }
      playerMoves = [];
      downBoard();
      document.getElementById('next').removeAttribute('hidden');
      document.getElementById('score').setAttribute('hidden', 'true');
      setTimeout(function () {
        document.getElementById('next').setAttribute('hidden', 'true');
        document.getElementById('score').removeAttribute('hidden');
        setTimeout(function () {
          startRound(false);
        }, 1000);
      }, 3000);
    }
  }

  function addMovement () {
    var rnd = colors[Math.floor((Math.random() * colors.length))];
    cpu.push(rnd);
    cpuMoves.push(colors.indexOf(rnd));
  }

  function animate (cb) {
    var i = 0;
    var interval = setInterval(function () {
      document.getElementById('counter').textContent = i + 1;
      document.getElementById('round').textContent = cpu.length;
      turnOn(colors.indexOf(cpu[i]), 500);
      i++;
      if (i >= cpu.length) {
        clearInterval(interval);
        // console.log('ANIMATION ENDS')
        cb('ANIMATION ENDS');
      }
    }, 1000);
  }

  function turnOn (p, time) {
    audio[p].play();
    document.getElementById(pos[p]).style.backgroundColor = colors2[p];
    setTimeout(function () {
      document.getElementById(pos[p]).style.backgroundColor = colors[p];
    }, time);
  }

  function upBoard () {
    document.getElementById('counter').textContent = playerMoves.length + 1;
    document.getElementById('round').textContent = cpuMoves.length;
    document.getElementsByClassName('board')[0].style.pointerEvents =
      'auto';
    var click = document.getElementsByClassName('action');
    for (var i = 0; i < click.length; i++) {
      click.item(i).addEventListener('click', userResponse);
    // console.log(click.item(i).id)
    }
  }

  function downBoard () {
    document.getElementsByClassName('board')[0].style.pointerEvents =
      'none';
  }

  function changeStatus (ev) {
    var target = ev.target.id;
    var x = document.getElementById(target);
    if (target === 'strict') {
      if (strict === 'off') {
        strict = 'on';
        x.style.borderLeft = '5px solid green';
        x.style.borderRight = '5px solid green';
      } else {
        strict = 'off';
        x.style.borderLeft = '5px solid red';
        x.style.borderRight = '5px solid red';
      }
    }
    if (target === 'start') {
      if (start === 'off') {
        start = 'on';
        x.style.borderLeft = '5px solid green';
        x.style.borderRight = '5px solid green';
        startRound(false);
      } else {
        start = 'off';
        x.style.borderLeft = '5px solid red';
        x.style.borderRight = '5px solid red';
        downBoard();
      }
    }
  }

  function reload () {
    location.reload();
  }

  addEventListener('load', init);
}());
