/*jshint node: true */
'use strict';

function sumFibs (num) {
  let total = 2;
  let fib = [1, 1];
  for (let i = 2; i < 100; i++) {
    let next = (fib[i - 1] + fib[i - 2]);
    if (next > num) {
      return total;
    }
    fib.push(next);
    if (next % 2 !== 0) {
      total += next;
    }
  }
}

console.log(sumFibs(4));
