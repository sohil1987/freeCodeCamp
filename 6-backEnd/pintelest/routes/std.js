'use strict';
/* jshint node: true */

var express = require('express');
var router = express.Router();
// var path = require('path')

var db = require('./../models/database.js');

var passport = require('passport');
var passConf = require('./../config/passport.js');

var baseUrl = ''; // 'https://brusbilis.com/freecodecamp/6-backEnd/pintelest'
var pics;

// router.use(express.static(__dirname + './../public'))

router.get('/', function (req, res, next) {
  // console.log('RUTA ', __dirname)
  db.getListaAllPics(req, res, function () {
    res.render('pages/home', {
      user: req.user, active: 'home', pics: res.pics
    });
  });
});

router.get('/user/:id', function (req, res, next) {
  // var name = req.params.name
  db.getListaUserPics(req, res, function () {
    // console.log('PICS .....', res.pics)
    res.render('pages/user', {
      user: req.user, active: 'user', pics: res.pics
    });
  });
});

router.get('/pics', function (req, res) {
  res.redirect(baseUrl + '/');
});

router.get('/myPics', passConf.checkAuthentication, function (req, res) {
  db.getListaMyPics(req, res, function () {
    // console.log('PICS .....', res.pics)
    res.render('pages/myPics', {
      user: req.user, active: 'myPics', pics: res.pics
    });
  });
});

router.get('/myPics/delete/:id', passConf.checkAuthentication,
  function (req, res) {
    // console.log('myPics...', req.params.id)
    db.deletePic(req, res, function () {
      res.redirect(baseUrl + '/myPics');
    });
  });

router.get('/addPic', passConf.checkAuthentication, function (req, res) {
  res.render('pages/addPic',
    { user: req.user, active: 'addPic',pics: undefined });
});

router.post('/addPic', passConf.checkAuthentication, function (req, res) {
  var body = req.body;
  // console.log('req.body', req.body)
  db.insertNewUser(req, res, function () {
    // res.redirect('/user/' + req.body.idTwitter)
    res.redirect(baseUrl + 'myPics');
  });
});

router.get('/profile', passConf.checkAuthentication, function (req, res) {
  res.render('pages/profile',
    { user: req.user, active: 'profile',  pics: undefined});
});

router.get('/vote/:userId/:picId', passConf.checkAuthentication,
  function (req, res) {
    db.postAddVote(req, res, function () {
      res.redirect(baseUrl + '/');
    });
  }
);

router.get('/logout', function (req, res) {
  req.session.destroy(function (err) {
    res.redirect(baseUrl + '/');
  });
});

router.get('/login/twitter',
  passport.authenticate('twitter'));

router.get('/login/twitter/return',
  passport.authenticate('twitter', { failureRedirect: '/login' }),
  function (req, res) {
    // res.redirect('/myPics')
    res.redirect(baseUrl + '/profile');
  });

module.exports = router;
