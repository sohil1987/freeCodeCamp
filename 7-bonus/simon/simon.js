(function () {
  'use strict';
  /* code here */

  var strict = 'off';
  var start = 'off';
  var options = ['green', 'red', 'yellow', 'blue'];
  var cpu = [];
  var player = [];

  function init () {
    console.log('Init');
    document.getElementById('strict').addEventListener('click', changeStatus);
    document.getElementById('start').addEventListener('click', changeStatus);
  }

  function startGame () {
    // upBoard()
    var fin = false;
    while (!fin) {
      addMovement();
      animate();
      fin = userResponse();
      fin = true;
    }
  }

  function addMovement () {
    var cpuMove = options[Math.floor((Math.random() * options.length))];
    cpu.push(cpuMove);
    console.log(cpu);
  }

  function animate () {
    for (var i = 0; i < cpu.length; i++) {
      setTimeout(console.log(cpu[i]), 1000);
    // console.log(cpu[i])
    }
  }

  function animate2 (sequence) {
    var i = 0;
    var interval = setInterval(function () {
      encenderLuz(sequence[i]);
      i++;
      if (i >= sequence.length) {
        clearInterval(interval);
      }
    }, 600);
  }

  function userResponse () {
    // console.log()
  }

  function aunNo (ev) {
    console.log(ev.target.id);
  }

  function upBoard () {
    document.getElementsByClassName('board')[0].style.pointerEvents = 'all';
    var click = document.getElementsByClassName('action');
    for (var i = 0; i < click.length; i++) {
      click.item(i).addEventListener('click', aunNo);
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
        startGame();
      } else {
        start = 'off';
        x.style.borderLeft = '5px solid red';
        x.style.borderRight = '5px solid red';
        downBoard();
      }
    }
  }

  addEventListener('load', init);
}());

/*
start
round
- random number (1-4) to sequence
- animate sequence to the player
- user interaction (while not bad response and clicks < sequence lebgth wait player input)

more rounds until fails
*/
