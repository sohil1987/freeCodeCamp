'use strict';
/* jshint node: true */

var express = require('express');
var router = express.Router();
// var path = require('path')

var passport = require('passport');
var passConf = require('./../config/passport.js');

router.use(express.static(__dirname + './../publicPintelest'));

router.get('/', function (req, res, next) {
  // console.log('RUTA ', __dirname)
  res.render('pages/home', {user: req.user,  active: 'home'});
});

router.get('/pics', function (req, res) {
  res.redirect('/');
});

router.get('/myPics', passConf.checkAuthentication, function (req, res) {
  res.render('pages/myPics' , {user: req.user,  active: 'myPics'});
});

router.get('/addPic', passConf.checkAuthentication, function (req, res) {
  res.render('pages/addPic' , {user: req.user,  active: 'addPic'});
});

router.get('/profile', passConf.checkAuthentication, function (req, res) {
  res.render('pages/profile' , {user: req.user,  active: 'profile'});
});

router.get('/logout', function (req, res) {
  req.session.destroy(function (err) {
    res.redirect('/');
  });
});

router.get('/login/twitter',
  passport.authenticate('twitter'));

router.get('/login/twitter/return',
  passport.authenticate('twitter', { failureRedirect: '/login' }),
  function (req, res) {
    res.redirect('/profile');
  });

module.exports = router;
