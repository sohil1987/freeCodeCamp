/*jshint node: true */
'use strict';

function dropElements (arr, func) {
  // Drop them elements.
  while (arr.length > 0 && !func(arr[0])) {
    arr.shift();
  }
  return arr;
}

console.log(dropElements([1, 2, 3], function (n) {return n < 3; }));
