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
  res.sendFile(publicPath + '/timestamp.html');
});

app.get('/:timeString', function (req, res) {
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

console.log('var');

app.listen(3001, function () {
  console.log('App started on port 3001');
});

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
