/*jshint node: true */
'use strict';

function diffArray (arr1, arr2) {
  let newArr = [];
  let base = arr1.concat(arr2);
  base.forEach(function (element) {
    if (arr1.includes(element) && arr2.includes(element)) {
      // console.log('Not valid')
    } else {
      newArr.push(element);
    }
  });
  return newArr;
}

console.log(diffArray([1, 2, 3, 5], [1, 2, 3, 4, 5]));
