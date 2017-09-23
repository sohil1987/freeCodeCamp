/*jshint node: true */
'use strict';

function sumPrimes (num) {
  let total = 0;
  let primes = [];
  for (let i = 2; i <= num; i++) {
    if (isPrime(i)) {
      primes.push(i);
      total += i;
    }
  }
  return total;
}

function isPrime (num) {
  for (let i = 2; i < num; i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

console.log(sumPrimes(10));
