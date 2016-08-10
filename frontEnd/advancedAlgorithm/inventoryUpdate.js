/*jslint node: true */
'use strict';

function updateInventory (current, update) {
  var stock = current.slice();
  var aux = [];
  for (var i = 0; i < current.length; i++) {
    aux[i] = current[i][1];
  }
  for (var i = 0; i < update.length; i++) {
    if (aux.indexOf(update[i][1]) !== -1) {
      // console.log(i, update[i][1])
      // console.log('POSICION', aux.indexOf(update[i][1]))
      // console.log(update[i][0])
      stock[aux.indexOf(update[i][1])][0] += update[i][0];
    } else {
      // console.log(i, update[i][1])
      stock.push([update[i][0], update[i][1]]);
    }
  }
  stock.sort(function (a, b) {
    if (a[1] === b[1]) {
      return 0;
    } else {
      return (a[1] < b[1]) ? -1 : 1;
    }
  });
  return stock;
}

console.log(
  updateInventory([
    [21, 'Bowling Ball'],
    [2, 'Dirty Sock'],
    [1, 'Hair Pin'],
    [5, 'Microphone']
  ], [
    [2, 'Hair Pin'],
    [3, 'Half-Eaten Apple'],
    [67, 'Bowling Ball'],
    [7, 'Toothpaste']
  ]));

console.log(updateInventory([], [
  [2, 'Hair Pin'],
  [3, 'Half-Eaten Apple'],
  [67, 'Bowling Ball'],
  [7, 'Toothpaste']
]));
