/*jslint node: true */
'use strict';

function checkCashRegister (price, cash, cid) {
  var change = [
    ['ONE HUNDRED', 0],
    ['TWENTY', 0],
    ['TEN', 0],
    ['FIVE', 0],
    ['ONE', 0],
    ['QUARTER', 0],
    ['DIME', 0],
    ['NICKEL', 0],
    ['PENNY', 0]
  ];
  var remaining = cash - price;
  cid = cid.reverse();
  var types = [100, 20, 10, 5, 1, 0.25, 0.10, 0.05, 0.01];
  var available = parseFloat(cid.reduce(getSum, 0)).toFixed(2);
  // console.log(available, 'EN CAJA para deuda de', remaining)
  if (available < remaining) {
    return 'Insufficient Funds';
  } else if (available == remaining) {
    return 'Closed';
  }
  var stillInRegister = cid.slice();
  var units = 0;
  var i = 0;
  while (remaining > 0 && i < types.length) {
    console.log('VUELTA ', i);
    units = cid[i][1] / types[i];
    if (types[i] <= remaining) {
      if (units * types[i] > remaining) {
        units = Math.floor(remaining / types[i]);
      }
      change[i][1] = units;
      remaining = remaining - units * types[i];
      remaining = parseFloat(remaining).toFixed(2);
    }
    stillInRegister.shift();
    var caja = stillInRegister.reduce(getSum, 0);
    // caja = parseFloat(caja).toFixed(2)
    remaining = parseFloat(remaining).toFixed(2);
    if (caja < remaining) {
      return 'Insufficient Funds';
    }
    i++;
  }
  return repair(change, types);
}

function repair (change, types) {
  change = change.filter(function (val, i, res) {
    val[1] = val[1] * types[i];
    return val[1] !== 0;
  });
  return change;
}

function getSum (total, num) {
  return total + num[1];
}
