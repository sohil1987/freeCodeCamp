'use strict';
/* jshint node: true */

var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var bodyParser = require('body-parser');

var secret = require('./secret.json');
var publicPath = path.join(__dirname, 'public');

var index = require('./_routes/index');
var apiV1 = require('./_routes/apiV1');

var app = express();

// view engine setup
// app.set('views', path.join(__dirname, 'views'))
// app.set('view engine', 'ejs')

// uncomment after placing your favicon in /public
app.use(favicon(path.join(publicPath, 'logo/brusbb.png')));
app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
  extended: false
}));
app.use(express.static(publicPath));

// routes
app.use('/api/v1', apiV1);
app.use('/', index);

app.listen(secret.express.port, function () {
  console.log('App started on port', secret.express.port);
});

// catch 404 and forward to error handler
app.use(function (req, res, next) {
  var err = new Error('Not Found');
  err.status = 404;
  next(err);
});

// error handler
app.use(function (err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500).json(res.locals.message);
});

module.exports = app;
