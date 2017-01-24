/*jslint node: true */
'use strict';

function myReplace (str, before, after) {
  console.log(after);
  if (before[0] === before[0].toUpperCase()) {
    after = after.charAt(0).toUpperCase() + after.slice(1);
    console.log(after);
  }
  return str.replace(before, after);
}

console.log(myReplace('He is Sleeping on the couch', 'Sleeping', 'sitting'));
// console.log(myReplace('A quick brown fox jumped over the lazy dog', 'jumped', 'leaped'))
