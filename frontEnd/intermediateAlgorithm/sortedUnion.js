/*jslint node: true */
'use strict';

function uniteUnique (arr) {
  var res = [];
  for (var i = 0; i < arguments.length; i++) {
    // console.log(arguments[i])
    for (var j = 0; j < arguments[i].length; j++) {
      // console.log(arguments[i][j])
      if (res.indexOf(arguments[i][j]) === -1) {
        res.push(arguments[i][j]);
      }
    }
  }
  return res;
}

console.log(uniteUnique([1, 3, 2], [5, 2, 1, 4], [2, 1]));
