/*jshint node: true */
'use strict';

function largestOfFour (arr) {
  let result = [];
  arr.forEach(function (element) {
    let max = element[0];
    element.forEach(function (value) {
      if (value > max) {
        max = value;
      }
    });
    result.push(max);
  });
  return result;
}

console.log(largestOfFour([[4, 5, 1, 3], [13, 27, 18, 26], [32, 35, 37, 39], [1000, 1001, 857, 1]]));
