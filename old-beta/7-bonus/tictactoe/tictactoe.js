/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Init TicTacToe');
    X.addEventListener('click', startGame);
    O.addEventListener('click', startGame);
  }

  let game = 'off';
  let player = '';
  let cpu = '';
  let board = [];
  let boardFree = [];

  function startGame (ev) {
    const side = ev.target.id; // event = pressed number box
    game = 'on';
    if (side === 'X') {
      player = 'X';
      cpu = 'O';
    }
    if (side === 'O') {
      player = 'O';
      cpu = 'X';
    }
    board = ['', '', '', '', '', '', '', '', ''];
    boardFree = [0, 1, 2, 3, 4, 5, 6, 7, 8];
    document.getElementsByClassName('board')[0].removeAttribute('hidden');
    // document.getElementsByClassName('control')[0].setAttribute('hidden', true)
    document.getElementsByClassName('control')[0].style.display = 'none';
    createBoardClicks();
  }

  function createBoardClicks () {
    const click = document.getElementsByClassName('action2');
    for (let i = 0; i < click.length; i++) {
      click.item(i).addEventListener('click', gameFlow);
    // console.log(click.item(i).id)
    }
  }

  function gameFlow (ev) {
    const box = ev.target.id; // event = pressed number box
    board[box] = player;
    document.getElementById(box).removeEventListener('click', gameFlow);
    document.getElementById(box).innerText = player;
    adjustBoards();
    if (win(player)) {
      gameEnds(player);
    } else if (boardFree.length === 0) {
      gameEnds('tie');
    } else {
      cpuMoves();
      if (win(cpu)) {
        gameEnds(cpu);
      }
    }
  // console.log('Board', board)
  // console.log('BoarFree', boardFree)
  }

  function gameEnds (user) {
    for (let i = 0; i < boardFree.length; i++) { // disable boxs on board
      document.getElementById(boardFree[i]).removeEventListener(
        'click', gameFlow);
    }
    document.getElementsByClassName('result')[0].removeAttribute('hidden');
    if (user === 'tie') {
      document.getElementById('gameResult').innerText = 'TIE !';
    }
    if (user === cpu) {
      document.getElementById('gameResult').innerText = 'You LOSE !';
    }
    if (user === player) {
      document.getElementById('gameResult').innerText = 'You WIN !';
    }
    document.getElementById('continue').addEventListener('click', reload);
  }

  function reload () {
    location.reload();
  }

  function win (user) {
    if (board[0] === user) {
      if (board[1] === user) {
        if (board[2] === user) return true;
      }
      if (board[3] === user) {
        if (board[6] === user) return true;
      }
      if (board[4] === user) {
        if (board[8] === user) return true;
      }
    }
    if (board[1] === user && board[4] === user && board[7] === user) return true;
    if (board[2] === user) {
      if (board[4] === user) {
        if (board[6] === user) return true;
      }
      if (board[5] === user) {
        if (board[8] === user) return true;
      }
    }
    if (board[3] === user && board[4] === user && board[5] === user) return true;
    if (board[6] === user && board[7] === user && board[8] === user) return true;
  // console.log(user, board)
  }

  function cpuMoves () {
    const cpuBox = boardFree[Math.floor((Math.random() * boardFree.length))];
    // console.log('Are', boardFree.length, 'boces and CPU takes', cpuBox)
    board[cpuBox] = cpu;
    document.getElementById(cpuBox).removeEventListener('click', gameFlow);
    document.getElementById(cpuBox).innerText = cpu;
    adjustBoards();
  }

  function adjustBoards () {
    let aux = [];
    for (let i = 0; i < boardFree.length; i++) {
      if (board[boardFree[i]] === '') {
        aux.push(boardFree[i]);
      }
    }
    boardFree = aux.slice();
  }

  return {
    init: init
  };
}());

window.addEventListener('load', app.init);
