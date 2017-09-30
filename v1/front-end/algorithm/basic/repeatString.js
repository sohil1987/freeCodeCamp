/*jshint node: true */
'use strict';

function repeatStringNumTimes (str, num) {
  let result = str;
  if (num < 0) {
    return '';
  }
  for (let i = 1; i < num; i++) {
    result += str;
  }
  return result;
}

console.log(repeatStringNumTimes('abc', 3));
