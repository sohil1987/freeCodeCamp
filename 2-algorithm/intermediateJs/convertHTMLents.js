/*jslint node: true */
'use strict';

function convertHTML (str) {
  // &colon;&rpar
  str = str.split('');
  for (var i = 0; i < str.length; i++) {
    str[i] = convert(str[i]);
  }
  return str.join('');
}

function convert (str) {
  if (str === '<') {return '&lt;';}
  if (str === '>') {return '&gt;';}
  if (str === '&') {return '&amp;';}
  if (str === '"') {return '&quot;';}
  if (str === "'") {return '&apos;';}
  return str;
}

console.log(convertHTML('Dolce & Gabbana'));
