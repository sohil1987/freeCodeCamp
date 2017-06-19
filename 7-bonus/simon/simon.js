/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  let strict = 'off';
  let start = 'off';
  const pos = ['green', 'red', 'yellow', 'blue'];
  const colors = ['green', '#b20000', '#cccc00', 'blue'];
  const colors2 = ['chartreuse', '#ff0000', '#ffff00', 'cyan'];
  let cpu = []; // ['green', '#b20000', '#cccc00', 'blue']
  let cpuMoves = []; // [0, 1, 2, 3]
  let playerMoves = [];
  let counter = 0;
  let round = 0;
  const audio1 = new Audio('./../../assets/sounds/simon/simon1.mp3');
  const audio2 = new Audio('./../../assets/sounds/simon/simon2.mp3');
  const audio3 = new Audio('./../../assets/sounds/simon/simon3.mp3');
  const audio4 = new Audio('./../../assets/sounds/simon/simon4.mp3');
  const audio = [audio1, audio2, audio3, audio4];

  function init () {
    console.log('Init Simon Game');
    document.getElementById('strict').addEventListener('click', changeStatus);
    document.getElementById('start').addEventListener('click', changeStatus);
    document.getElementById('restart').addEventListener('click', reload);
  }

  return {
    init: init
  };
}());

window.addEventListener('load', app.init);
