/*jshint node: true */

var dungeon = {
  cols: 50,
  rows: 40,
  ppp: 10,
  map: [],
  center: {},
  maxRooms: 20,
  rooms: [],
  minSizeRoom: 5,
  maxSizeRoom: 12,
  minLenCorridor: 4,
  maxLenCorridor: 15,
  enemyDensity: 20, // less is higher
  randomWall: {},
  nextFeature: {},
  newRoom: {},
  foes: [],
  createDungeon: function () {
    'use strict';
    this.createSize();
    this.createRooms();
    this.populate();
    this.getFoes();
  },
  createSize: function () {
    'use strict';
    this.map = help.initializeMultiArray(this.cols, this.rows, 0);
  },
  createRooms: function () {
    'use strict';
    this.createSingleRoomInCenter();
    this.putWalls(0);

    // for (var i = 0; i < this.maxRooms; i++) {
    while (this.rooms.length < this.maxRooms) {
      this.pickRandomWallFromAnyRoom();
      this.pickRandomFeature();
      this.convertFeatureToRoom();
      if (this.checkIsRoomForFeature()) {
        this.addFeature();
        this.putAllWalls();
      } else {
        this.map[this.randomWall.x][this.randomWall.y] = 1;
      }
    }
  },
  createSingleRoomInCenter: function () {
    'use strict';
    var width = help.getRandomNumber(this.minSizeRoom, this.maxSizeRoom);
    var height = help.getRandomNumber(this.minSizeRoom, this.maxSizeRoom);
    var x = Math.floor((this.cols - width) / 2);
    var y = Math.floor((this.rows - height) / 2);
    for (var i = x - 1; i < x + width - 1; i++) {
      for (var j = y - 1; j < y + height - 1; j++) {
        this.map[i][j] = 2;
      }
    }
    this.rooms.push({});
    this.rooms[0].x = x - 1;
    this.rooms[0].y = y - 1;
    this.rooms[0].width = width;
    this.rooms[0].height = height;
    this.rooms[0].foes = Math.floor(this.rooms[0].width * this.rooms[0].height /
      this.enemyDensity) + Math.floor(Math.random() * 2);
  // console.log(this.rooms[0])
  },
  pickRandomWallFromAnyRoom: function () {
    'use strict';
    var x, y;
    var found = false;
    var cont = 0;
    while (!found && cont < 500) {
      x = help.getRandomNumber(1, this.cols - 1);
      y = help.getRandomNumber(1, this.rows - 1);
      if (this.map[x][y] === 1) {
        if (this.isNotCorner(x, y)) {
          // console.log('BOTIN', x, y)
          this.map[x][y] = 2;
          found = true;
        }
      }
      cont++;
    }
    this.randomWall.x = x;
    this.randomWall.y = y;
  },
  pickRandomFeature: function () {
    'use strict';
    var type = help.getRandomNumber(0, 100);
    this.nextFeature.x = this.randomWall.x;
    this.nextFeature.y = this.randomWall.y;
    this.nextFeature.dir = this.getFeatureDirection(this.randomWall.x,
      this.randomWall.y);
    switch (true) {
      case ( type < 50):
        // console.log(type, 'corridor')
        if (this.nextFeature.dir === 1 || this.nextFeature.dir === 3) {
          this.nextFeature.width = 1;
          this.nextFeature.height = help.getRandomNumber(this.minLenCorridor,
            this.maxLenCorridor);
        } else if (this.nextFeature.dir === 2 || this.nextFeature.dir ===
          4) {
          this.nextFeature.width = help.getRandomNumber(this.minLenCorridor,
            this.maxLenCorridor);
          this.nextFeature.height = 1;
        }
        break;
      case ( type >= 50):
        // console.log(type, 'room')
        this.nextFeature.width = help.getRandomNumber(this.minSizeRoom, this
          .maxSizeRoom);
        this.nextFeature.height = help.getRandomNumber(this.minSizeRoom, this
          .maxSizeRoom);
        break;
    }
  },
  convertFeatureToRoom: function () {
    'use strict';
    // console.log('New Feature', this.nextFeature)
    if (this.nextFeature.dir === 1) {
      this.newRoom.x = this.nextFeature.x;
      this.newRoom.y = this.nextFeature.y - this.nextFeature.height;
    } else if (this.nextFeature.dir === 2) {
      this.newRoom.x = this.nextFeature.x + 1;
      this.newRoom.y = this.nextFeature.y;
    } else if (this.nextFeature.dir === 3) {
      this.newRoom.x = this.nextFeature.x;
      this.newRoom.y = this.nextFeature.y + 1;
    } else if (this.nextFeature.dir === 4) {
      this.newRoom.x = this.nextFeature.x - this.nextFeature.width;
      this.newRoom.y = this.nextFeature.y;
    }
    this.newRoom.width = this.nextFeature.width;
    this.newRoom.height = this.nextFeature.height;
    this.newRoom.foes = Math.floor(this.nextFeature.width * this.nextFeature
        .height / this.enemyDensity) + Math.floor(Math.random() * 2);
  },
  checkIsRoomForFeature: function () {
    'use strict';
    // console.log('New Room', this.newRoom)
    var isRoom = true;
    for (var i = this.newRoom.x; i < this.newRoom.x + this.newRoom.width; i++) {
      for (var j = this.newRoom.y; j < this.newRoom.y + this.newRoom.height; j++) {
        if (i < 0 || i >= this.cols || j < 0 || j >= this.rows) { // Protection
          isRoom = false;
        } else if (this.map[i][j] === 2) {
          isRoom = false;
        }
      }
    }
    // console.log(isRoom)
    return isRoom;
  },
  addFeature: function () {
    'use strict';
    // console.log('GO ON !')
    for (var i = this.newRoom.x; i < this.newRoom.x + this.newRoom.width; i++) {
      for (var j = this.newRoom.y; j < this.newRoom.y + this.newRoom.height; j++) {
        this.map[i][j] = 2;
      }
    }
    this.rooms.push(this.newRoom);
    this.newRoom = {};
  // console.log(this.rooms[this.rooms.length - 1].width)
  },
  populate: function () {
    'use strict';
    // console.log('Rooms', this.rooms.length)
    var x, y;
    for (var i = 0; i < this.rooms.length; i++) {
      var room = this.rooms[i];
      var units = room.foes;
      // console.log('Room', i, this.rooms[i])
      while (units > 0) { // FOES
        x = help.getRandomNumber(room.x, room.x + room.width - 1);
        y = help.getRandomNumber(room.y, room.y + room.height - 1);
        if (this.map[x][y] === 2) {
          this.map[x][y] = 3;
          units--;
        }
      }
      units = room.foes;
      while (units > 1) { // WEAPONS
        x = help.getRandomNumber(room.x, room.x + room.width - 1);
        y = help.getRandomNumber(room.y, room.y + room.height - 1);
        if (this.map[x][y] === 2) {
          this.map[x][y] = 4;
          units--;
        }
      }
      units = room.foes;
      while (units > 1) { // POTIONS
        x = help.getRandomNumber(room.x, room.x + room.width - 1);
        y = help.getRandomNumber(room.y, room.y + room.height - 1);
        if (this.map[x][y] === 2) {
          this.map[x][y] = 5;
          units--;
        }
      }
    }
    // HERO{
    var hero = false;
    while (!hero) {
      x = help.getRandomNumber(1, this.cols - 1);
      y = help.getRandomNumber(1, this.rows - 1);
      if (this.map[x][y] === 2) {
        this.map[x][y] = 6;
        // console.log('HERO', x, y)
        this.center.x = x;
        this.center.y = y;
        hero = true;
      }
    }
    // BOSS
    var boss = false;
    while (!boss) {
      x = help.getRandomNumber(1, this.cols - 1);
      y = help.getRandomNumber(1, this.rows - 1);
      if (this.map[x][y] === 2) {
        this.map[x][y] = 7;
        // console.log('BOSS', x, y)
        boss = true;
      }
    }
  },

  getFoes: function () {
    var cont = 0;
    for (var x = 0; x < dungeon.cols; x++) {
      for (var y = 0; y < dungeon.rows; y++) {
        if (dungeon.map[x][y] === 3) {
          var myFoe = Object.create(app.foe);
          myFoe.x = x;
          myFoe.y = y;
          this.foes.push(myFoe);
        }
      }
    }
  // console.log(this.foes)
  },

  putWalls: function (roomNumber) {
    'use strict';
    var x = this.rooms[roomNumber].x;
    var y = this.rooms[roomNumber].y;
    var width = this.rooms[roomNumber].width;
    var height = this.rooms[roomNumber].height;
    // console.log(x, y, roomWidth, roomHeight)
    for (var i = x - 1; i < x + 1 + width; i++) {
      for (var j = y - 1; j < y + 1 + height; j++) {
        if (this.map[i][j] === 0) {
          // console.log(i, j)
          this.map[i][j] = 1;
        }
      }
    }
  },
  putAllWalls: function () {
    'use strict';
    for (var x = 0; x < this.cols; x++) {
      for (var y = 0; y < this.rows; y++) {
        // console.log(x, y)
        if (this.map[x][y] === 0) {
          if (y > 0 && this.map[x][y - 1] === 2) { // NORTH
            this.map[x][y] = 1;
          } else if (x < this.cols - 1 && this.map[x + 1][y] === 2) { // EAST
            this.map[x][y] = 1;
          } else if (y < this.rows - 1 && this.map[x][y + 1] === 2) { // SOUTH
            this.map[x][y] = 1;
          } else if (x > 0 && this.map[x - 1][y] === 2) { // WEST
            this.map[x][y] = 1;
          }
        }
      }
    }
  },
  isNotCorner: function (x, y) {
    'use strict';
    // console.log(x, y)
    if (x < 1 || x >= this.cols - 1 || y < 1 || y >= this.rows - 1) { // Protec
      return false;
    } else if (this.map[x - 1][y] === 2) {
      return true;
    } else if (this.map[x + 1][y] === 2) {
      return true;
    } else if (this.map[x][y - 1] === 2) {
      return true;
    } else if (this.map[x][y + 1] === 2) {
      return true;
    } else return false;
  },
  getFeatureDirection: function (x, y) {
    'use strict';
    var aux = help.getRandomNumber(1, 4);
    var cont = 0;
    var ok = false;
    while (!ok && cont < 10) {
      if (x < 0 || x >= this.cols || y < 0 || y >= this.rows) { // Protection
        // console.log()
      } else if (this.map[x][y + 1] === 2 && aux === 1) { // NORTH
        ok = true;
      } else if (this.map[x - 1][y] === 2 && aux === 2) { // EAST
        ok = true;
      } else if (this.map[x][y - 1] === 2 && aux === 3) { // SOUTH
        ok = true;
      } else if (this.map[x + 1][y] === 2 && aux === 4) { // WEST
        ok = true;
      } else {
        cont++;
        aux = help.getRandomNumber(1, 4);
      }
    }
    return aux;
  }
};
