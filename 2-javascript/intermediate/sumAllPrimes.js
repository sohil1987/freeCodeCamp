/*jslint node: true */
'use strict';

function sumPrimes (limit) {
  var total = 2;
  var serie = [];
  for (var num = 3; num <= limit; num++) {
    if (isPrime(num)) {
      total += num;
    }
  }
  return total;
}

function isPrime (num) {
  var res = true;
  for (var i = 2; i < num; i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

console.log('SUMA --> ', sumPrimes(10));
console.log(sumPrimes(977));
