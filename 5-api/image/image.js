'use strict';
/* jshint node: true */

var express = require('express');
var logger = require('morgan');
var path = require('path');
var request = require('request');

var secret = require('./secret.json');
var db = require('./database.js');

var app = express();

app.use(logger('dev'));
var publicPath = path.resolve(__dirname, 'public');
app.use(express.static(publicPath));

app.get('/api/search', function (req, res, next) {
  var resPlus = res; // avoid conflict with res in request(searchUrl) function
  if (!req.query.num) { req.query.num = 10; }
  if (!req.query.offset) { req.query.offset = 10; }
  if (req.query.num > 10) {req.query.num = 10; }
  if (req.query.offset > req.query.num) {req.query.offset = req.query.num;}
  var searchUrl = 'https://www.googleapis.com/customsearch/v1?q=' + req.query.q + '&cx=' + secret.api.cseID + '&num=' + req.query.num + '&key=' + secret.api.key;
  // var searchUrl = 'http://localhost:3004/test.json'
  var images = [];
  request(searchUrl, function (err, res, body) {
    console.log('STATUS CODE ==> ' /*, res.statusCode*/, resPlus.statusCode);
    console.log('ERROR ==> ', err);
    console.log(req.query.q);
    if (!err && resPlus.statusCode === 200) {
      body = JSON.parse(body);
      // console.log(body)
      var image = {};
      // console.log(req.query)
      for (var i = 0; i < req.query.offset; i++) {
        image.title = body.items[i].title;
        image.id = body.items[i].cacheId;
        image.snippet = body.items[i].snippet;
        image.displayLink = body.items[i].displayLink;
        try {
          image.thumbSrc = body.items[i].pagemap.cse_thumbnail[0].src;
        } catch(e) {
          image.thumbSrc = '';
        }
        try {
          image.imageSrc = body.items[i].pagemap.cse_image[0].src;
        } catch (e) {
          image.imageSrc = '';
        }
        images.push(image);
        image = {};
      }
    }
    db.create(req, res);
    resPlus.send('<pre>' + JSON.stringify(images, null, 2) + '</pre>');
  });
});

app.get('/api/recent', function (req, res) {
  db.getList(function (result) {
    res.send('<pre>' + JSON.stringify(result, null, 2) + '</pre>');
  });
});

app.get('*', function (req, res) {
  res.sendFile(publicPath + '/image.html');
});

app.listen(secret.express.port, function () {
  console.log('App started on port:' + secret.express.port);
});
