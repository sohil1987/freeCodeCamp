'use strict';
/* jshint node: true */

var express = require('express');
var logger = require('morgan');
var path = require('path');

var secret = require('./secret.json');
var db = require('./database.js');
var validUrl = require('valid-url');

var app = express();

app.use(logger('dev'));

var publicPath = path.resolve(__dirname, 'public');
app.use(express.static(publicPath));

app.get('/new/:newUrl*', function (req, res, next) {
  req.params.newUrl = req.url.slice(5);
  if (validUrl.isUri(req.params.newUrl)) {
    db.create(req, res, creationNew);
  } else {
    res.sendFile(publicPath + '/url.html');
  }
});

app.get('/:idurl', function (req, res, next) {
  if (!isNaN(req.params.idurl)) {
    db.read(req, res, redirect);
  } else {
    console.log(req.params.idurl, 'Not Number = Not valid idurl url');
    // res.sendFile(publicPath + '/url.html') // or go next
    next();
  }
});

app.get('*', function (req, res) {
  res.sendFile(publicPath + '/url.html');
});

app.listen(secret.express.port, function () {
  console.log('App started on port:' + secret.express.port);
});

function creationNew (req, res) {
  console.log('CREATION NEW', res.original, res.idurl);
  res.json({
    'original': res.original,
    'idurl': res.idurl
  });
}

function redirect (req, res) {
  if (res.result) {
    res.redirect(res.result[0].original);
  } else {
    res.sendFile(publicPath + '/url.html');
  }
}
