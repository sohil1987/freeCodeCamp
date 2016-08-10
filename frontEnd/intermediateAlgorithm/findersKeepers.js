/*jslint node: true */
'use strict';

function findElement (arr, func) {
  var num = arr.filter(test);
  return num[0];
}

function test (num) {
  return num % 2 === 0;
}

console.log(findElement([1, 3, 5, 8, 9, 10], function (num) { return num % 2 === 0; }));
