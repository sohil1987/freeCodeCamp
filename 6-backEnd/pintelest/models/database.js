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
  },
  getListaAllPics: function (req, res, callback) {
    var sql = 'SELECT idPic, pics.idTwitter, users.nameTwitter AS username, users.logoTwitter AS logoAuthor, text, link, likes ';
    sql += 'FROM pics ';
    sql += 'JOIN users ON pics.idTwitter = users.idTwitter ';
    var inserts = [];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err, rows) {
      if (err) {
        throw err;
      } else {
        res.pics = rows;
        // console.log('RESULTADOS --> ', rows.length)
        // console.log(rows)
        callback();
      }
    });
  },
  getListaUserPics: function (req, res, callback) {
    var sql = 'SELECT idPic, users.nameTwitter AS username, users.logoTwitter AS logoAuthor, text, link, likes ';
    sql += 'FROM pics ';
    sql += 'JOIN users ON pics.idTwitter = users.idTwitter ';
    sql += 'WHERE pics.idTwitter = ?';
    var inserts = [req.params.id];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err, rows) {
      if (err) {
        throw err;
      } else {
        res.pics = rows;
        res.pics.id = req.params.id;
        // console.log('RESULTADOS --> ', rows.length)
        // console.log(rows)
        callback();
      }
    });
  },
  postAddVote: function (req, res, callback) {
    var sql = 'SELECT COUNT (*) as times ';
    sql += 'FROM likes ';
    sql += 'WHERE idPic = ? AND idTwitter = ?';
    var inserts = [req.params.picId, req.params.userId];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err, rows) {
      if (err) {
        throw err;
      } else {
        res.rows = rows;
        // console.log('RESULTADOS DE AÑADIR VOTO --> ', rows[0].times)
        // console.log(rows)
        if (rows[0].times === 0) {
          db.insertAddVote(req, res, callback);
        } else if (rows[0].times === 1) {
          db.insertDeleteVote(req, res, callback);
        } else {
          callback(res);
        }
      }
    });
  },
  insertAddVote: function (req, res, callback) {
    var sql = 'INSERT INTO likes VALUES (?, ?)';
    var inserts = [req.params.picId, req.params.userId];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err) {
      if (err) {
        throw err;
      } else {
        // console.log('Registro GUARDADO ...........')
        db.addLikeToPic(req, res, callback(res));
      }
    });
  },
  insertDeleteVote: function (req, res, callback) {
    var sql = 'DELETE FROM likes WHERE idPic = ? AND idTwitter = ?';
    var inserts = [req.params.picId, req.params.userId];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err) {
      if (err) {
        throw err;
      } else {
        // console.log('Registro GUARDADO ...........')
        db.subLikeToPic(req, res, callback(res));
      }
    });
  },
  addLikeToPic: function (req, res, callback) {
    var sql = 'UPDATE pics SET likes = likes +1 WHERE idPic = ?';
    var inserts = [req.params.picId];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err) {
      if (err) {
        throw err;
      } else {
        // console.log('ADD +1 LIKE ...........')
        return;
      }
    });
  },
  subLikeToPic: function (req, res, callback) {
    var sql = 'UPDATE pics SET likes = likes -1 WHERE idPic = ?';
    var inserts = [req.params.picId];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err) {
      if (err) {
        throw err;
      } else {
        // console.log('DEC -1 LIKE ...........')
        return;
      }
    });
  },
  addNewPic: function (req, res, callback) {
    // console.log(req.body)
    // console.log('USER', req.user)
    var sql = 'INSERT INTO pics (idTwitter, text, link) ';
    sql += 'VALUES(?, ?, ?)';
    var inserts = [req.user.id, req.body.textPic, req.body.linkPic];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err) {
      if (err) {
        throw err;
      } else {
        console.log('NEW PIC SAVED ...........');
        callback();
      }
    });
  },
  insertNewUser: function (req, res, callback) {
    var sql = 'INSERT INTO users(idTwitter, nameTwitter,logoTwitter) ';
    sql += 'SELECT ?, ?, ? FROM dual ';
    sql += 'WHERE NOT EXISTS ';
    sql += '(SELECT idTwitter FROM users WHERE idTwitter = ?) ';
    sql += 'LIMIT 1';
    var inserts = [req.user.id, req.user.username, req.user.photos[0].value, req.user.id];
    sql = mysql.format(sql, inserts);
    // console.log(sql)
    con.query(sql, function (err) {
      if (err) {
        throw err;
      } else {
        console.log('NEW USER SAVED ...........');
        db.addNewPic(req, res, callback);
      }
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
