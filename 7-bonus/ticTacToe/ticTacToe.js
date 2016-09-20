(function () {
  'use strict';
  /* code here */

  var game = 'off';
  var player = '';
  var board = [];

  function init () {
    console.log('Init');
    var click = document.getElementsByClassName('action');
    for (var i = 0; i < click.length; i++) {
      click.item(i).addEventListener('click', clickAction);
    // console.log(click.item(i).id)
    }
  }

  function clickAction (ev) {
    var event = ev.target.id;
    if (game === 'on') {
      if (event !== 'X' && event !== 'O') {
        console.log('Partida en marcha, jugador elige', event);
        gameFlow(event);
      }
    } else if (event === 'X' || event === 'O') {
      startGame(event);
    }
  }

  function gameFlow (event) {
    console.log(event);
  }

  function startGame (side) {
    game = 'on';
    if (side === 'X') player = 'X';
    if (side === 'O') player = 'O';
    console.log(side);
    board = ['', '', '', '', '', '', '', '', ''];
    console.log(board);
    gameFlow(bur);
  }

  addEventListener('load', init);
}());
