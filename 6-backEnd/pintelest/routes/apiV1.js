'use strict';
/* jshint node: true */

var express = require('express');
var router = express.Router();

var controllers = require('./../controllers/controllers.js');

router.use(function (req, res, next) {
  res.header('Access-Control-Allow-Origin', '*');
  res.header('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
  next();
});

router.get('/', controllers.test);

router.get('/test1', controllers.test1);

router.get('/test2', controllers.test2);

//  Golang microJson
/*
router.get('/test3', function (req, res) {
  var http = require('http')
  var url = 'http://127.0.0.1:3020'
  var callback = function (response) {
    var str = ''
    // another chunk of data has been recieved, so append it to `str`
    response.on('data', function (chunk) {
      str += chunk
    })
    // the whole response has been recieved, so we just print it out here
    response.on('end', function () {
      console.log(str)
    })
  }
  http.request(url, callback).end()
  res.end()
})
*/
module.exports = router;
