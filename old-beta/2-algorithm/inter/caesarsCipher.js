/*jshint node: true */
'use strict';

function rot13 (str) {
  let sol = '';
  for (let i = 0; i < str.length; i++) {
    if (/[A-Z]/.test(str[i])) { // if is a letter between A-Z
      console.log(str[i], str.charCodeAt(i));
      if (str.charCodeAt(i) > 77) {
        sol += String.fromCharCode(str.charCodeAt(i) - 13);
      } else {
        sol += String.fromCharCode(str.charCodeAt(i) + 13);
      }
    } else {
      sol += str[i];
    }
  }
  return sol;
}

console.log(rot13('SERR PBQR PNZC'));
