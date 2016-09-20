/*jslint node: true */
'use strict';

function steamrollArray (arr) {
  var result = [];

  var flatten = function (arg) {
    if (Array.isArray(arg)) {
      arg.forEach(flatten);
    } else {
      result.push(arg);
    }
  };
  arr.forEach(flatten);
  return result;
}

steamrollArray([1, [2],
  [3, [
    [4]
  ]]
]);
