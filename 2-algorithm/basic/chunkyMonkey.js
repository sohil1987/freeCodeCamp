/*jshint node: true */
'use strict';

function chunkArrayInGroups (arr, size) {
  let result = [];
  let partial = [];
  arr.forEach(function (ele, pos) {
    partial.push(ele);
    if ((pos + 1) % size === 0) {
      result.push(partial);
      partial = [];
    } else if (pos === arr.length - 1) {
      result.push(partial);
    }
  });
  return result;
}

console.log(chunkArrayInGroups(['a', 'b', 'c', 'd'], 2));
// console.log(chunkArrayInGroups([0, 1, 2, 3, 4, 5], 4))
// console.log(chunkArrayInGroups([0, 1, 2, 3, 4, 5, 6, 7, 8], 4))
