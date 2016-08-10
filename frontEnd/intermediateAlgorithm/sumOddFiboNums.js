/*jslint node: true */
'use strict';

function sumFibs (limit) {
  var total = 1;
  var serie = [0, 1];
  var i = 2;
  while (i <= limit) {
    var sum = serie[i - 1] + serie[i - 2];
    if (sum <= limit) {
      serie.push(sum);
      if (sum % 2 !== 0) {
        total += sum;
      }
    }
    i++;
  }
  return total;
}

console.log(sumFibs(5));
console.log(sumFibs(1000));
console.log(sumFibs(4000000));
console.log(sumFibs(4));
console.log(sumFibs(75024));
console.log(sumFibs(75025));
