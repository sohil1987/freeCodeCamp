/*jslint node: true */
'use strict';

function sym () {
  var args = [];
  for (var i = 0; i < arguments.length; i++) {
    args.push(arguments[i]);
  }
  return args.reduce(symmetric);
}

function symmetric (a1, a2) {
  var result = [];
  a1.forEach(function (index) {
    if (a2.indexOf(index) === -1 && result.indexOf(index) === -1) {
      result.push(index);
    }
  });
  a2.forEach(function (index) {
    if (a1.indexOf(index) === -1 && result.indexOf(index) === -1) {
      result.push(index);
    }
  });
  return result;
}

// console.log(sym([1, 2, 3], [5, 2, 1, 4]))
console.log(sym([3, 3, 3, 2, 5], [2, 1, 5, 7], [3, 4, 6, 6], [1, 2, 3]));
