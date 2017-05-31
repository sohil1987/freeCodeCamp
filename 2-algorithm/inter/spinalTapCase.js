/*jshint node: true */
'use strict';

function spinalCase (str) {
  str = str.replace(/ |_/g, '-').split('');
  str.forEach(function (element, pos) {
    if (/[A-Z]/.test(element)) {
      str[pos] = element.toLowerCase();
      if (pos > 0 && str[pos - 1] !== '-') {
        str.splice(pos , 0, '-');
      }
    }
  });
  return str.join('');
}

console.log(spinalCase('ThisIs Spinal Tap'));
