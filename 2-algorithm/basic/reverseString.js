/*jshint node: true */
'use strict';

function reverseString (str) {
  str = str.split('').reverse().join('');
  return str;
}

console.log(reverseString('hello'));
