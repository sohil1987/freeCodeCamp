'use strict';
/* jshint node: true */

var express = require('express');
var logger = require('morgan');
var path = require('path');

var app = express();

app.use(logger('dev'));

var publicPath = path.resolve(__dirname, 'public');
app.use(express.static(publicPath));

app.get('/', function (req, res) {
  res.sendFile(publicPath + '/parser.html');
});

app.get('/v1/whoami', function (req, res) {
  var ip = req.headers['x-forwarded-for'] || req.connection.remoteAddress;
  if (ip.substr(0, 7) == '::ffff:') {
    ip = ip.substr(7);
  }
  var lang = (JSON.stringify(req.headers['accept-language'])).slice(1, 6);
  var ops = JSON.stringify(req.headers['user-agent']).split('(');
  ops = ops[1].toString().split(';');
  res.json({
    'IP Address': ip,
    'Language': lang,
    'Operative System': ops[0] + ops[1]
  });
});

app.listen(3002, function () {
  console.log('App started on port 3002');
});
