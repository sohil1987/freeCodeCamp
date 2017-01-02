'use strict';
/* jshint node: true */

var secret = require('./../../secret.json');
var db = require('./../models/database.js');

var controller = {
  test: function (req, res) {
    db.testConnection(req, res, connStatus);
  }

};

module.exports = controller;

function enviarLista (res) {
  res.send('<pre>' + JSON.stringify(res.rows, null, 2) + '</pre>');
}

function connStatus (res) {
  console.log('connStatus = >', res.status);
  if (res.status) {
    res.json('OK');
  } else if (!res.status) {
    res.json('MAL');
  }
}
