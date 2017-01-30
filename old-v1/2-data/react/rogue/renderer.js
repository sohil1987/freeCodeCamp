var rendererPlay = (function () {
  'use strict';
  // console.log('Drawing Map...')
  var canvas;
  var ctx;
  var cols = 13;
  var rows = 11;
  var ppp = 40;
  var center = dungeon.center;
  var posX = cols - Math.floor(cols / 2);
  var posY = rows - Math.floor((rows / 2));
  // var hero = new Image()
  // hero.src = ('./images/knight.png')

  function setMyCanvas () {
    canvas = document.getElementById('myCanvas');
    ctx = canvas.getContext('2d');
    canvas.width = cols * ppp;
    canvas.height = rows * ppp;
  }

  function drawMyCanvasCenteredOnHero (map) {
    // map[1][1] = 5
    // var kk = (ppp / 50) * 0.07
    // console.log('Dungeon center at', center.x, center.y)
    var x, y;
    for (x = center.x - posX; x < center.x + posX; x++) {
      // console.log('COLUMNA', x)
      if (x >= 0 && x < dungeon.cols) {
        for (y = center.y - posY; y < center.y + posY; y++) {
          if (map[x][y] === undefined || x > dungeon.cols - 1 || y >
            dungeon.rows - 1) {
            ctx.fillStyle = 'black'; // no exists
          } else {
            var tile = map[x][y];
            // console.log(x, y, tile)
            if (tile === 0) {
              ctx.fillStyle = 'teal'; // earth
            } else if (tile === 1) {
              ctx.fillStyle = 'darkgrey'; // wall
            } else if (tile === 2) {
              ctx.fillStyle = 'bisque'; // walkable floor
            } else if (tile === 3) {
              ctx.fillStyle = 'coral'; // foe
            } else if (tile === 4) {
              ctx.fillStyle = 'skyblue'; // weapon
            } else if (tile === 5) {
              ctx.fillStyle = 'darkseagreen'; // potions
            } else if (tile === 6) {
              ctx.fillStyle = 'magenta'; // hero
            } else if (tile === 7) {
              ctx.fillStyle = 'gold'; // FINAL BOSS
            }
          }
          var i = x - center.x - posX + cols;
          var j = y - center.y - posY + rows;
          ctx.fillRect((i * ppp) + 1, (j * ppp) + 1, ppp - 1, ppp - 1);
        // console.log((i * ppp) + 1, (j * ppp) + 1, ppp - 1, ppp - 1)
        /*if (tile === 5) {
          ctx.drawImage(hero, (x * ppp) + 6, (y * ppp), hero.width * kk,
            hero.height * kk)
        } */
        }
      } else {
        for (y = center.y - posY; y < center.y + posY; y++) {
          ctx.fillStyle = 'black'; // no exists
          var i2 = x - center.x - posX + cols;
          var j2 = y - center.y - posY + rows;
          ctx.fillRect((i2 * ppp) + 1, (j2 * ppp) + 1, ppp - 1, ppp - 1);
        // console.log((i * ppp) + 1, (j * ppp) + 1, ppp - 1, ppp - 1)
        }
      }
    }
  }
  return {
    setMyCanvas: setMyCanvas,
    drawMyCanvasCenteredOnHero: drawMyCanvasCenteredOnHero
  };
}());

var rendererBig = (function () {
  'use strict';
  // console.log('Drawing Map...')
  var canvas;
  var ctx;
  var cols = dungeon.cols;
  var rows = dungeon.rows;
  var ppp = dungeon.ppp;

  function setMyCanvas () {
    canvas = document.getElementById('myCanvas');
    ctx = canvas.getContext('2d');
    canvas.width = cols * ppp;
    canvas.height = rows * ppp;
  }

  function drawMyCanvas (map) {
    for (var x = 0; x < cols; x++) {
      for (var y = 0; y < rows; y++) {
        var tile = map[x][y];
        if (tile === 0) {
          ctx.fillStyle = 'teal'; // earth
        } else if (tile === 1) {
          ctx.fillStyle = 'darkgrey'; // wall
        } else if (tile === 2) {
          ctx.fillStyle = 'bisque'; // walkable floor
        } else if (tile === 3) {
          ctx.fillStyle = 'coral'; // foe
        } else if (tile === 4) {
          ctx.fillStyle = 'skyblue'; // weapon
        } else if (tile === 5) {
          ctx.fillStyle = 'darkseagreen'; // potions
        } else if (tile === 6) {
          ctx.fillStyle = 'magenta'; // hero
        } else if (tile === 7) {
          ctx.fillStyle = 'gold'; // FINAL BOSS
        }
        ctx.fillRect((x * ppp) + 1, (y * ppp) + 1, ppp - 1, ppp - 1);
      }
    }
  }
  return {
    setMyCanvas: setMyCanvas,
    drawMyCanvas: drawMyCanvas
  };
}());
