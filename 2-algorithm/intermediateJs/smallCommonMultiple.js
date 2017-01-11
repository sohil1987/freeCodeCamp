/*jslint node: true */
'use strict';

function smallestCommons (arr) {
  arr = arr.sort(function (a, b) {
    return a - b;
  });
  var test = [];
  var min = arr[0];
  var max = arr[1];
  var index = max * 2;
  var found = false;
  for (var i = min; i <= max; i++) {
    test.push(i);
  }
  while (!found) { // } && index < 100) {
    if (isOk(index, test) && isOk(index, test)) {
      found = true;
    }
    index++;
  }
  return index - 1;
}

function isOk (num, test) {
  for (var i = 0; i < test.length; i++) {
    // console.log('looking ', num, ' divisible by ', test[i])
    if (num % test[i] !== 0) {
      return false;
    }
  }
  return true;
}

console.log('RES --> ', smallestCommons([1, 5]));
console.log('RES --> ', smallestCommons([5, 1]));
console.log('RES --> ', smallestCommons([1, 13]));
console.log('RES --> ', smallestCommons([23, 18]));
