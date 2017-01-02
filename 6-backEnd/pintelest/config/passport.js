'use strict';
/* jshint node: true */

var passport = require('passport');
var secret = require('./../secret.json');
var Strategy = require('passport-twitter').Strategy;

passport.use(new Strategy({
  consumerKey: secret.twitter.consumerKey,
  consumerSecret: secret.twitter.consumerSecret,
  callbackURL: secret.twitter.callbackURL
}, function (token, tokenSecret, profile, cb) {
  return cb(null, profile);
}));

passport.serializeUser(function (user, cb) {
  // console.log('serializeUser')
  cb(null, user);
});

passport.deserializeUser(function (obj, cb) {
  // console.log('deserializeUser')
  cb(null, obj);
});

exports.checkAuthentication = function (req, res, next) {
  if (req.isAuthenticated()) {
    return next();
  }
  res.redirect('/');
};
