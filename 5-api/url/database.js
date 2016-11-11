'use strict';
/* jshint node: true */

var secure = require('./secret.json');
var mysql = require('mysql');

var baseUrl = 'https://brusbilis.com/5-api/url/';
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
  getError: function (req, res) {
    return function (error) {
      res.send(500, {error: error.message});
    };
  },
  getList: function (callback) {
    console.log('Getting list......');
    con.query('SELECT * FROM url', function (err, result) {
      if (err) throw err;
      callback(result);
    });
  },
  create: function (req, res, callback) {
    console.log('create', req.params.newUrl);
    con.query('SELECT * FROM url WHERE url.original = ?', [req.params.newUrl], function (err, result) {
      if (err) throw err;
      if (result.length > 0) { // url already in the DB 
        res.original = result[0].original;
        res.idurl = baseUrl + result[0].idurl;
        callback(req, res);
      } else { // new, we add in the DB 
        var data = { original: req.params.newUrl };
        con.query('INSERT INTO url SET ?', data, function (err, result) {
          if (err) throw err;
          res.original = data.original;
          res.idurl = baseUrl + result.insertId;
          callback(req, res);
        });
      }
    });
  },
  read: function (req, res, callback) {
    console.log('read', req.params.idurl);
    con.query('SELECT * FROM url WHERE url.idurl = ?', [req.params.idurl], function (err, result) {
      if (err) throw err;
      if (result.length > 0) {
        res.result = result;
      } else {
        res.result = undefined;
      }
      callback(req, res);
    });
  },
  update: function (req, res) {
    console.log('update');
  },
  delete: function (req, res) {
    console.log('delete');
  }
};

module.exports = db;
