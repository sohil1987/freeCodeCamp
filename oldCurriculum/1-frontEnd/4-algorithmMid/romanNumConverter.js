/*jslint node: true */
'use strict';

function convertToRoman (num) {
  var result = '';
  var ones, tens, hunds, thous;
  var hundsA = ['C', 'CC', 'CCC', 'CD', 'D', 'DC', 'DCC', 'DCCC', 'CM'];
  var tensA = ['X', 'XX', 'XXX', 'XL', 'L', 'LX', 'LXX', 'LXXX', 'XC'];
  var onesA = ['I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX'];
  thous = Math.floor(num / 1000);
  hunds = Math.floor(num % 1000 / 100);
  tens = Math.floor(num % 100 / 10);
  ones = Math.floor(num % 10);
  console.log(thous, hunds, tens, ones);
  for (var i = 0; i < thous;i++) {
    result += 'M';
  }
  if (hunds) {
    result += hundsA[hunds - 1];
  }
  if (tens) {
    result += tensA[tens - 1];
  }
  if (ones) {
    result += onesA[ones - 1];
  }
  return result;
}

console.log(convertToRoman(36));
console.log(convertToRoman(68));
console.log(convertToRoman(97));
console.log(convertToRoman(1000));
console.log(convertToRoman(3999));
