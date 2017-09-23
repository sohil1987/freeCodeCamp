/*jshint node: true */
'use strict';

function truncateString (str, num) {
  if (str.length <= num) {
    return str;
  }
  let result = str.slice(0, num) + '...';
  return result;
}

console.log(truncateString('A-tisket a-tasket A green and yellow basket', 8));
