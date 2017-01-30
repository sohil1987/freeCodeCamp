'use strict';
/* jshint node: true */

var express = require('express');
var logger = require('morgan');
var path = require('path');
var multer = require('multer');

var app = express();

app.use(logger('dev'));

var storage = multer.memoryStorage();
var upload = multer({ storage: storage });
var type = upload.single('inputFile');

var publicPath = path.resolve(__dirname, 'public');
app.use(express.static(publicPath));

app.get('*', function (req, res) {
  res.sendFile(publicPath + '/file.html');
});

app.post('/', type, function (req, res) {
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

app.listen(3005, function () {
  console.log('App started on port 3005');
});
