/*jshint node: true */
'use strict';

function smallestCommons (arr) {
  arr = arr.sort(function (a, b) {
    return a - b;
  });
  let min = arr[0];
  let max = arr[1];
  const bad = worstSol(min, max);
  console.log('Lets GO -- ', min, max, bad);
  for (let i = min; i < bad; i++) {
    if (iFoundIt(i, min, max)) {
      return i;
    }
  }
  return bad;
}

function iFoundIt (num, min, max) {
  for (let i = min; i <= max; i++) {
    if (num % i !== 0) {
      return false;
    }
  }
  return true;
}

function worstSol (min, max) {
  let sol = 1;
  for (let i = min; i <= max; i++) {
    sol *= i;
  }
  return sol;
}

console.log(smallestCommons([23, 18]));
