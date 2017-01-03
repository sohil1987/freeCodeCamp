'use strict';
/* jshint node: true */

var secret = require('./../secret.json');
var mysql = require('mysql');

var db = {
  testConnection: function (req, res, callback) {
    // console.log('Connecting ......')
    con.connect(function (err) {
      if (err) {
        console.log('Error connecting to DB => ', err);
        res.status = false;
      } else {
        console.log('Connection OK');
        res.status = true;
      }
      con.end(function () {});
      callback(res);
    });
  }
};

var con = mysql.createConnection({
  host: secret.mysql.host,
  user: secret.mysql.user,
  password: secret.mysql.password,
  database: secret.mysql.db
});

module.exports = db;
