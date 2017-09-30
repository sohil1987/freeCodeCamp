/*jshint node: true */
'use strict';

function factorialize (num) {
  if (num === 0) return 1;
  var total = 1;
  for (let i = 1; i <= num; i++) {
    total *= i;
  }
  return total;
}

console.log(factorialize(5));
