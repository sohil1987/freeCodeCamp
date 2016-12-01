'use strict';
/* jshint node: true */

var express = require('express');
var logger = require('morgan');
var path = require('path');
var multer = require('multer');
var request = require('request');
var validUrl = require('valid-url');

var secret = require('./secret.json');
var db = require('./database.js');

var app = express();

app.use(logger('dev'));

var storage = multer.memoryStorage();
var upload = multer({ storage: storage });
var type = upload.single('inputFile');

var publicPath = path.resolve(__dirname, 'public');
app.use(express.static(publicPath));

// TimeStamp

app.get('/timestamp', function (req, res) {
  res.sendFile(publicPath + '/timestamp/timestamp.html');
});

app.get('/timestamp/:timeString', function (req, res) {
  var param = req.params.timeString;
  var unixTime = 0;
  var natural = '';
  var isNumber = /^\d+$/;
  if (isNumber.test(param)) {
    unixTime = Math.floor(parseInt(param));
    natural = unixToNatural(unixTime);
  } else {
    var aux = String(param).split(' ');
    if (aux.length === 3) {
      var day = aux[1];
      if (day[day.length - 1] === ',') {
        day = aux[1].slice(0, aux[1].length - 1); // remove "," after day
      }
      var month = aux[0].toLowerCase();
      var year = aux[2];
      var need1 = months.indexOf(month) !== -1;
      var need2 = isNumber.test(day) && day > 0 && day < 32;
      var need3 = isNumber.test(year) && year > 0;
      if (need1 && need2 && need3) {
        natural = month + ' ' + day + ', ' + year;
        aux[1] = aux[1].slice(0, aux[1].length - 1);
        unixTime = naturalToUnix(aux);
      } else {
        unixTime = null;
        natural = null;
      }
    } else {
      unixTime = null;
      natural = null;
    }
  }
  res.json({
    unix: unixTime,
    natural: natural
  });
});

// Parser

app.get('/parser', function (req, res) {
  res.sendFile(publicPath + '/parser/parser.html');
});

app.get('/parser/v1/whoami', function (req, res) {
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

// Url

app.get('/url/new/:newUrl*', function (req, res, next) {
  req.params.newUrl = req.url.slice(9);
  if (validUrl.isUri(req.params.newUrl)) {
    db.create(req, res, creationNew);
  } else {
    res.sendFile(publicPath + '/url/url.html');
  }
});

app.get('/url/:idurl', function (req, res, next) {
  if (!isNaN(req.params.idurl)) {
    db.read(req, res, redirect);
  } else {
    console.log(req.params.idurl, 'Not Number = Not valid idurl url');
    // res.sendFile(publicPath + '/url/url.html') // or go next
    next();
  }
});

app.get('/url/*', function (req, res) {
  res.sendFile(publicPath + '/url/url.html');
});

// Image

app.get('/image/api/search', function (req, res, next) {
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
    db.createImage(req, res);
    resPlus.send('<pre>' + JSON.stringify(images, null, 2) + '</pre>');
  });
});

app.get('/image/api/recent', function (req, res) {
  db.getListImage(function (result) {
    res.send('<pre>' + JSON.stringify(result, null, 2) + '</pre>');
  });
});

app.get('/image', function (req, res) {
  res.sendFile(publicPath + '/image/image.html');
});

// File

app.get('/file', function (req, res) {
  res.sendFile(publicPath + '/file/file.html');
});

app.post('/file/file', type, function (req, res) {
  if (req.file) {
    res.json({
      'File': req.file.originalname,
      'Size': req.file.size
    });
  } else {
    res.json({
      'ERROR': 'File not exists'
    });
  }
});

app.listen(secret.express.port, function () {
  console.log('App started on port', secret.express.port);
});

// TimeStamp

function unixToNatural (unixTime) {
  unixTime = unixTime * 1000; // in js milliseconds not seconds
  var date = new Date(unixTime);
  var natural = months[date.getMonth()] + ' ' + date.getDate() + ', ' + date.getFullYear();
  return natural;
}

function naturalToUnix (natural) {
  var unixTime = new Date(natural[0] + ', ' + natural[1] + ', ' + natural[2]).getTime();
  return unixTime / 1000; // in js milliseconds not seconds
}

var months = ['january', 'february', 'march', 'april', 'may', 'june', 'july',
  'august', 'september', 'october', 'november', 'december'
];

// Url

function creationNew (req, res) {
  console.log('CREATION NEW', res.original, res.idurl);
  res.json({
    'original': res.original,
    'idurl': res.idurl
  });
}

function redirect (req, res) {
  console.log('REDIRECTING ..........', res.result[0].original);
  if (res.result) {
    res.redirect(res.result[0].original);
  } else {
    res.sendFile(publicPath + '/url/url.html');
  }
}
