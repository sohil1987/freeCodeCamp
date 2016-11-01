var app = (function () {
  'use strict';
  /* code here */
  // http://www.roguebasin.com/index.php?title=Main_Page
  // http://www.roguebasin.com/index.php?title=Dungeon-Building_Algorithm

  var weapon = [
    [0, 'dagger', '1d6+4'],
    [1, 'sword', '2d6+6'],
    [2, '2H sword', '3d6+8'],
    [3, 'magical sword', '4d6+10'],
    [10, 'foe weapon', '3d6+0']
  ];

  var hero = {
    xp: 0,
    level: 1,
    hp: 100,
    wp: 0,
    damage: function () {
      return getDamage(weapon[this.wp][2]);
    }
  };

  var foe = {
    xp: 100,
    level: 1,
    hp: 30,
    wp: 4,
    damage: function () {
      return getDamage(weapon[this.wp][2]);
    }
  };

  var foes = [];

  function init () {
    console.log('Inicio');
    dungeon.createDungeon();
    rendererBig.setMyCanvas();
    rendererBig.drawMyCanvas(dungeon.map);
    getFoes();
    addEvents();
  // console.log('DAÑO del heroe ...', hero.damage())
  // console.log('DAÑO del enemigo ...', foe.damage())
  }

  function getDamage (dice) {
    console.log(dice);
    var damage = 0;
    dice = dice.split('d6+');
    for (var i = 0; i < dice[0]; i++) {
      damage += bb.getRandomNumber(1, 6);
    }
    damage += parseInt(dice[1]);
    return damage;
  }

  function getFoes () {
    var cont = 0;
    for (var x = 0; x < dungeon.cols; x++) {
      for (var y = 0; y < dungeon.rows; y++) {
        if (dungeon.map[x][y] === 3) {
          var myFoe = Object.create(foe);
          myFoe.x = x;
          myFoe.y = y;
          foes.push(myFoe);
        }
      }
    }
    console.log(foes);
  }

  function addEvents () {
    // console.log(hero.damage)
    document.getElementById('myCanvas').addEventListener('keydown', action);
    document.getElementById('myCanvas').focus();
  }

  function action (e) {
    var key = String.fromCharCode(e.keyCode);
    console.log(key);
    if (key === '8' || key === 'W' || key === 'w' || key === 'h') {
      // console.log(key, '... N')
    } else if (key === '6' || key === 'E' || key === 'e' || key === 'f') {
      // console.log(key, '... E')
    } else if (key === '2' || key === 'S' || key === 's' || key === 'b') {
      // console.log(key, '... S')
    } else if (key === '4' || key === 'Q' || key === 'q' || key === 'd') {
      // console.log(key, '... W')
    }
  }

  return {
    init: init
  };
}());

addEventListener('load', app.init);
