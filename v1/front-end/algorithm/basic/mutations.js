/*jshint node: true */
'use strict';

function mutation (arr) {
  let isMutation = true;
  let word = arr[0].toLowerCase();
  let test = arr[1].toLowerCase().split('');
  test.forEach(function (letter) {
    console.log(letter, ' === ', word.indexOf(letter));
    if (word.indexOf(letter) === -1) isMutation = false;
  });
  return isMutation;
}

console.log(mutation(['hello', 'hey']));
