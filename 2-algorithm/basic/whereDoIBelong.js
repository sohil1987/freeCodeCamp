/*jshint node: true */
'use strict';

function getIndexToIns (arr, num) {
  let index = 0;
  arr.sort(function (a, b) {
    return a - b;
  });
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] >= num) {
      return i;
    }
  }
  return arr.length;
}

// console.log(getIndexToIns([40, 60], 50))
console.log(getIndexToIns([2, 5, 10], 15));
