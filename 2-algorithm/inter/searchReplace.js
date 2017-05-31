/*jshint node: true */
'use strict';

function myReplace (str, before, after) {
  if (before[0] === before[0].toUpperCase()) {
    after = after.replace(after[0], after[0].toUpperCase());
  }
  str = str.replace(before, after);
  return str;
}

console.log(myReplace('A quick brown fox jumped over the lazy dog', 'jumped', 'leaped'));
