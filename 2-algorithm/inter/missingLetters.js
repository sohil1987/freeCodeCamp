/*jshint node: true */
'use strict';

const letters = 'abcdefghijklmnopqrstuvwxyz';

function fearNotLetter (str) {
  let begin = letters.indexOf(str[0]);
  let end = letters.indexOf(str[str.length - 1]);
  let j = begin;
  for (let i = 0; i < end - begin; i++) {
    if (str[i] === letters[j]) {
      j++;
    } else {
      return letters[j];
    }
  }
  return undefined;
}

// console.log(fearNotLetter('abce'))
console.log(fearNotLetter('stvwx'));
