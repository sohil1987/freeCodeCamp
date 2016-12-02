'use strict';
/* jshint node: true */

var express = require('express');
var router = express.Router();
var path = require('path');

router.use(express.static(__dirname + './../public3'));

router.get('/', function (req, res, next) {
  console.log('RUTA ', __dirname);
  res.sendFile(path.join(__dirname + './../public3/clone.html'));
});

module.exports = router;
