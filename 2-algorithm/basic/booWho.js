/*jshint node: true */
'use strict';

function booWho (bool) {
  bool = typeof (bool) === 'boolean';
  return bool;
}

console.log(booWho(null));
