/*jslint node: true */
'use strict';

function pairwise (arr, arg) {
  var result = [];

  for (var i = 0; i < arr.length; i++) {
    for (var j = 0; j < arr.length; j++) {
      if (i !== j && arr[i] + arr[j] === arg) {
        if (result.indexOf(i) === -1 && result.indexOf(j) === -1) {
          result.push(i, j);
        }
      }
    }
  }

  if (arr.length === 0) {
    return 0;
  }
  result = result.reduce(function (a, b) {
    return a + b;
  });
  return result;
}

console.log(pairwise([1, 4, 2, 3, 0, 5], 7));
console.log(pairwise([1, 3, 2, 4], 4));
console.log(pairwise([1, 1, 1], 2));
console.log(pairwise([0, 0, 0, 0, 1, 1], 1));
console.log(pairwise([], 100));
