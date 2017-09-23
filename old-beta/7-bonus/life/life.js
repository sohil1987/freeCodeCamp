/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const activeBtns = [0, 4, 6];
  const listBtns = [0, 1, 2, 3, 4, 5, 6, 7, 8];
  let status = [];
  let future = [];
  const canvas = document.getElementById('myCanvas');
  const context = canvas.getContext('2d');
  let turn = 0;
  let cols = 0;
  let rows = 0;
  let ppl = 10;
  let living = 0;
  let fps = 0;
  let onoff;
  let running = true;
  let cancel = true;

  function init () {
    console.log('Init Game of Life');
    addControlEvents();
    setControls();
    changeSpeed(activeBtns[1]);
    run();
  }

  function paintCanvasAndCreateStatus (state) {
    setMyCanvas();
    populateStatus(state);
    drawBorderMyCanvas();
    drawCirclesMyCanvas();
    addEventsForAddCells();
  }

  function prepareNextTurn () {
    let neighbors;
    for (let i = 0; i < cols; i++) {
      for (let j = 0; j < rows; j++) {
        neighbors = getNeighbors(i, j);
        if (status[i][j] === 0) {
          if (neighbors === 3) {
            future[i][j] = 1;
            living++;
          }
        } else if (status[i][j] === 1) {
          if (neighbors < 2) {
            future[i][j] = 0;
            living--;
          } else if (neighbors > 3) {
            future[i][j] = 0;
            living--;
          }
        } else {
          console.log('STATUS ERROR');
        }
      }
    }
  }

  function getNeighbors (x, y) {
    let neighbors = 0;
    let i, j;
    if (x > 0 && x < cols - 1 && y > 0 && y < rows - 1) { // INSIDE
      for (i = x - 1; i < x + 2; i++) {
        for (j = y - 1; j < y + 2; j++) {
          posibleNeighbor();
        }
      }
    } else if (x === 0 && y > 0 && y < rows - 1) { // WEST NO CORNERS
      for (i = x; i < x + 2; i++) {
        for (j = y - 1; j < y + 2; j++) {
          posibleNeighbor();
        }
      }
    } else if (x === cols - 1 && y > 0 && y < rows - 1) { // EAST NO CORNERS
      for (i = x - 1; i < x + 1; i++) {
        for (j = y - 1; j < y + 2; j++) {
          posibleNeighbor();
        }
      }
    } else if (y === 0 && x > 0 && x < cols - 1) { // NORTH NO CORNERS
      for (i = x - 1; i < x + 2; i++) {
        for (j = y; j < y + 2; j++) {
          posibleNeighbor();
        }
      }
    } else if (y === rows - 1 && x > 0 && x < cols - 1) { // SOUTH NO CORNERS
      for (i = x - 1; i < x + 2; i++) {
        for (j = y - 1; j < y + 1; j++) {
          posibleNeighbor();
        }
      }
    } else if (x === 0 && y === 0) { // NW CORNER
      for (i = x; i < x + 2; i++) {
        for (j = y; j < y + 2; j++) {
          posibleNeighbor();
        }
      }
    } else if (x === cols - 1 && y === 0) { // NE CORNER
      for (i = x - 1; i < x + 1; i++) {
        for (j = y; j < y + 2; j++) {
          posibleNeighbor();
        }
      }
    } else if (x === cols - 1 && y === rows - 1) { // SE CORNER
      for (i = x - 1; i < x + 1; i++) {
        for (j = y - 1; j < y + 1; j++) {
          posibleNeighbor();
        }
      }
    } else if (x === 0 && y === rows - 1) { // SW CORNER
      for (i = x; i < x + 2; i++) {
        for (j = y - 1; j < y + 1; j++) {
          posibleNeighbor();
        }
      }
    }

    function posibleNeighbor () {
      if (i !== x || j !== y) {
        if (status[i][j] === 1) neighbors++;
      }
    }
    return neighbors;
  // console.log(x, y, 'has', neighbors, 'neighbors')
  }

  function update () {
    turn++;
    prepareNextTurn();
    status = future.slice();
    if (living <= 0) {
      cancel = true;
    }
  }

  function render () {
    document.getElementById('score').innerText = turn;
    document.getElementById('living').innerText = living;
    drawCirclesMyCanvas();
  }

  function gameLoop () {
    setTimeout(function () {
      if (running) {
        update();
        render();
      }
      onoff = requestAnimationFrame(gameLoop);
      console.log('gameLoop sigue ... ', onoff);
      if (cancel) {
        cancelAnimationFrame(onoff);
        setScoreToZero();
        setLivingCellsToZero();
      }
    }, 1000 / fps);
  }

  function run () {
    running = true;
    document.getElementById('1').innerText = 'Pause';
    paintCanvasAndCreateStatus('random');
    setScoreToZero();
    if (cancel) {
      gameLoop();
    }
    cancel = false;
  }

  function pause () {
    if (running) {
      document.getElementById('1').innerText = 'Resume';
    } else {
      document.getElementById('1').innerText = 'Pause';
    }
    running = !running;
  }

  function clear () {
    cancel = true;
    paintCanvasAndCreateStatus('cls');
  }

  function changeSpeed (speed) {
    switch (speed) {
      case 3:
        fps = 0.5;
        break;
      case 4:
        fps = 1;
        break;
      case 5:
        fps = 3;
    }
  }

  function changeSize () {
    clear();
  }

  function click (e) {
    setControls(e);
    switch (parseInt(e.target.id)) {
      case 0:
        run();
        break;
      case 1:
        pause();
        break;
      case 2:
        clear();
        break;
      case 3:
      case 4:
      case 5:
        changeSpeed(parseInt(e.target.id));
        break;
      case 6:
      case 7:
      case 8:
        changeSize();
        break;
      default:
        console.log(e.target.id, ' -- Not recognized event');
    }
  }

  function setMyCanvas () {
    const aux = document.getElementById(activeBtns[2]).innerHTML.split('x');
    cols = parseInt(aux[0]);
    rows = parseInt(aux[1]);
    canvas.width = cols * ppl;
    canvas.height = rows * ppl;
  }

  function drawBorderMyCanvas () {
    context.rect(0, 0, cols * ppl, rows * ppl);
    context.lineWidth = 1;
    context.strokeStyle = '#333333';
    context.stroke();
  }

  function drawCirclesMyCanvas () {
    for (let i = 0; i < cols; i++) {
      for (let j = 0; j < rows; j++) {
        context.beginPath();
        context.arc(i * ppl + ppl / 2, j * ppl + ppl / 2, ppl / 2 - 1, 0, 2 *
          Math.PI, false);
        if (status[i][j] === 0) {
          context.fillStyle = 'white';
          context.strokeStyle = '#cccccc';
        } else if (status[i][j] === 1) {
          context.fillStyle = 'coral';
          context.strokeStyle = 'coral';
        } else {
          console.log('STATUS ERROR');
        }
        context.fill();
        context.lineWidth = 1;
        // context.strokeStyle = '#cccccc'
        context.stroke();
      }
    }
  }

  function populateStatus (filler) {
    setLivingCellsToZero();
    let aux;
    for (let i = 0; i < cols; i++) {
      status[i] = [];
      for (let j = 0; j < rows; j++) {
        if (filler === 'random') {
          // aux = Math.floor(Math.random() * 2) // 50% 0 - 50% 1
          aux = Math.floor(Math.random() * 4);
          if (aux > 0) {
            aux = 0;
          } else if (aux === 0) {
            aux = 1;
            living++;
          }
        } else if (filler === 'cls') {
          aux = 0;
        }
        status[i][j] = aux;
      }
    }
    future = status.slice();
    document.getElementById('living').innerText = living;
  // console.log(future[15][15])
  }

  function setScoreToZero () {
    turn = 0;
    document.getElementById('score').innerText = turn;
  }

  function setLivingCellsToZero () {
    living = 0;
    document.getElementById('living').innerText = living;
  }

  function addControlEvents () {
    const aux = document.getElementsByClassName('btn');
    for (let i = 0; i < aux.length; i++) {
      aux[i].addEventListener('click', click);
    }
  }

  function addEventsForAddCells () {
    document.getElementById('myCanvas').addEventListener('click', clickCanvas);
  }

  function clickCanvas (e) {
    if (!running) {
      const rect = canvas.getBoundingClientRect();
      const i = Math.ceil((e.clientX - rect.left) / ppl) - 1;
      const j = Math.ceil((e.clientY - rect.top) / ppl) - 1;
      // console.log(status[i][j])
      if (status[i][j] === 1) {
        status[i][j] = 0;
        living--;
        document.getElementById('living').innerText = living;
      } else {
        status[i][j] = 1;
        living++;
        document.getElementById('living').innerText = living;
      }
      drawCirclesMyCanvas();
    }
  }

  function setControls (e) {
    if (e) {
      const pos = Math.floor(parseInt(e.target.id) / 3);
      activeBtns[pos] = parseInt(e.target.id);
    }
    for (let i = 0; i < listBtns.length; i++) {
      if (activeBtns.indexOf(i) !== -1) {
        document.getElementById(i).style.backgroundColor = 'burlywood';
      } else {
        document.getElementById(i).style.backgroundColor = '#fafafa';
      }
    }
  // console.log(activeBtns)
  }

  return {
    init: init
  };
}());

window.addEventListener('load', app.init);
