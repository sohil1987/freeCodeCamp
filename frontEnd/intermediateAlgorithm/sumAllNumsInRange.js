/*jslint node: true */
'use strict';

function sumAll (arr) {
  var max = Math.max(...arr);
  var min = Math.min(...arr);
  for (var i = 0; i <= max - min; i++) {
    arr[i] = min + i;
  }
  var sum = arr.reduce(function (a, b) {
    return a + b;
  });
  return sum;
}

console.log(sumAll([1, 4]));
console.log(sumAll([10, 5]));
