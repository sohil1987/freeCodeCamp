'use strict';
/* jshint node: true */

var express = require('express');
var helmet = require('helmet');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var bodyParser = require('body-parser');

var secret = require('./secret.json');
var publicPath = path.join(__dirname, '_public');

var indexPintelest = require('./pintelest/routes/index');
var apiV1Pintelest = require('./pintelest/routes/apiV1');
var indexNightlife = require('./nightlife/routes/index');
var apiV1Nightlife = require('./nightlife/routes/apiV1');

var passport = require('passport');
var Strategy = require('passport-twitter').Strategy;
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var session = require('express-session');

passport.use(new Strategy({
  consumerKey: secret.twitter.consumerKey,
  consumerSecret: secret.twitter.consumerSecret,
  callbackURL: 'http://localhost:3000/pintelest' // /login/twitter/return'
}, function (token, tokenSecret, profile, cb) {
  return cb(null, profile);
}));

passport.serializeUser(function (user, cb) {
  cb(null, user);
});

passport.deserializeUser(function (obj, cb) {
  cb(null, obj);
});

var app = express();

// view engine setup
// app.set('views', path.join(__dirname, 'views'))
// app.set('view engine', 'ejs')

app.use(logger('dev'));
app.use(helmet());
// uncomment after placing your favicon in /public
app.use(favicon(path.join(publicPath, './logo/brusbb.png')));
app.use(cookieParser());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
  extended: true
}));
app.use(session({
  secret: 'keyboard cat',
  resave: true,
  saveUninitialized: true
}));
app.use(express.static(publicPath));

// Initialize Passport and restore authentication state, if any, from the
// session.
app.use(passport.initialize());
app.use(passport.session());

// several apps running on the same app
// app.use(express.static(__dirname + '/pintelest'))
// app.use(express.static(__dirname + '/nightlife'))

// routes

app.use('/pintelest/api/v1', apiV1Pintelest);
app.use('/pintelest', indexPintelest);
app.use('/pintelest/login/twitter', indexPintelest);
app.use('/pintelest/profile', indexPintelest);
app.use('/nightlife/api/v1', apiV1Nightlife);
app.use('/nightlife', indexNightlife);
app.use('/', function (req, res) {
  res.sendFile(publicPath + '/index.html');
});

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
