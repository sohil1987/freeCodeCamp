'use strict';
/* jshint node: true */

var secure = require('./secret.json');
var mysql = require('mysql');

var con = mysql.createConnection({
  host: secure.mysql.host,
  user: secure.mysql.user,
  password: secure.mysql.password,
  database: secure.mysql.db
});

var db = {
  testConnection: function () {
    console.log('Connecting......');
    con.connect(function (err) {
      if (err) {
        console.log('Error connecting to DB');
        return;
      }
      console.log('Connection OK');
      con.end(function () {});
    });
  },
  getList: function (callback) {
    console.log('Getting list......');
    con.query('SELECT * FROM image ORDER BY image.idimage DESC LIMIT 10', function (err, result) {
      if (err) throw err;
      // console.log(result)
      callback(result);
    });
  },
  create: function (req, res) {
    console.log('create', req.query.q);
    var data = { search: req.query.q };
    con.query('INSERT INTO image SET ?', data, function (err, result) {
      if (err) throw err;
    });
  }
};

module.exports = db;
