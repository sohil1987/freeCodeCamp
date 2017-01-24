/*jslint node: true */
'use strict';

function diffArray (arr1, arr2) {
  var newArr = [];
  newArr = joinArrays(arr1, newArr);
  newArr = joinArrays(arr2, newArr);
  for (var i = 0; i < newArr.length; i++) {
    if (arr1.indexOf(newArr[i]) !== -1 && arr2.indexOf(newArr[i]) !== -1) {
      newArr.splice(i, 1);
      i--;
    }
  }
  return newArr;
}

function joinArrays (arr, newArr) {
  for (var i = 0; i < arr.length; i++) {
    // if (!newArr.includes(arr[i])) {    // ES7
    if ( (newArr.indexOf(arr[i]) === -1)) {
      newArr.push(arr[i]);
    }
  }
  return newArr;
}

console.log(diffArray([1, 2, 3, 5], [1, 2, 3, 4, 5]));
console.log(diffArray(['diorite', 'andesite', 'grass', 'dirt', 'pink wool', 'dead shrub'], ['diorite', 'andesite', 'grass', 'dirt', 'dead shrub']));
