/*jshint node: true */
'use strict';

function destroyer (arr) {
  let help = [];
  for (let i = 1; i < arguments.length; i++) {
    help.push(arguments[i]);
  }
  for (let i = arr.length - 1; i >= 0; i--) {
    // console.log(i, arr[i])
    if (help.includes(arr[i])) {
      // arr.splice(arr.indexOf(arr[i]), 1)
      arr.splice(i, 1);
    }
  }
  return arr;
}

console.log(destroyer([1, 2, 3, 1, 2, 3], 2, 3));
