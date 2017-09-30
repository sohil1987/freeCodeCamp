/*jshint node: true */
'use strict';

function bouncer (arr) {
  var result = [];
  result = arr.filter(removeFalsy);
  return result;
}

function removeFalsy (data) {
  if (data === false || data === null || data === 0 || data === '' || data ===
    undefined) {
    return;
  }
  return data;
}

console.log(bouncer([7, 'ate', '', false, 9]));
