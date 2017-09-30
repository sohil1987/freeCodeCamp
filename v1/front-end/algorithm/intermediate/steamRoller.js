/*jshint node: true */
'use strict';

function steamrollArray (arr) {
  let result = [];
  let flatten = function (arg) {
    if (Array.isArray(arg)) {
      arg.forEach(flatten);
    } else {
      result.push(arg);
    }
  };
  arr.forEach(flatten);
  return result;
}

console.log(steamrollArray([1, [2], [3, [[4]]]]));
