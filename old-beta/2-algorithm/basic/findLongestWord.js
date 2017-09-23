/*jshint node: true */
'use strict';

function findLongestWord (str) {
  str = str.split(' ');
  let longest = 0;
  str.forEach(function (word) {
    if (word.length > longest) {
      longest = word.length;
    }
  });
  return longest;
}

console.log(findLongestWord('The quick brown fox jumped over the lazy dog'));
