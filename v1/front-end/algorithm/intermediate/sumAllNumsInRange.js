/*jshint node: true */
'use strict';

function sumAll (arr) {
  let total = 0;
  let min = parseInt(arr[0]);
  let max = parseInt(arr[1]);
  if (min > max) {
    let aux = min;
    min = max;
    max = aux;
  }
  for (let i = min; i <= max; i++) {
    total += i;
  }
  return total;
}

console.log(sumAll([1, 4]));
