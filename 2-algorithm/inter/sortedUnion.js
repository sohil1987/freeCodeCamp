/*jshint node: true */
'use strict';

function uniteUnique () {
  let sol = [];
  let arr = [];
  for (let i = 0; i < arguments.length; i++) {
    for (let j = 0; j < arguments[i].length; j++) {
      arr.push(arguments[i][j]);
    }
  }
  for (let i = 0; i < arr.length; i++) {
    console.log(arr[i], sol.includes(arr[i]));
    if (!sol.includes(arr[i])) {
      sol.push(arr[i]);
    }
  }
  return sol;
}

console.log(uniteUnique([1, 3, 2], [5, 2, 1, 4], [2, 1]));
