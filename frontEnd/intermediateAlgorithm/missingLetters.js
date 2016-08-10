/*jslint node: true */
'use strict';

var test = 'abcdefghijklmnopqrstuvwxyz';

function fearNotLetter (str) {
  var begin = test.indexOf(str[0]);
  console.log(test[begin]);
  for (var i = 0; i < str.length; i++) {
    if (str[i] !== test[begin + i]) {
      console.log('BINGO --> ', str[i], ' !== ', test[begin + i]);
      return test[begin + i];
    }
  }
  return undefined;
}

console.log('d --> ', fearNotLetter('abce'));
console.log('i --> ', fearNotLetter('abcdefghjklmno'));
console.log('undefined --> ', fearNotLetter('bcd'));
console.log('undefined --> ', fearNotLetter('yz'));
