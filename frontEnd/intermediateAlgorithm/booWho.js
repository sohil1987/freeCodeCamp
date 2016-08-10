/*jslint node: true */
'use strict';

function booWho (bool) {
  // What is the new fad diet for ghost developers? The Boolean.
  bool = typeof (bool) === 'boolean';
  return bool;
}

console.log(booWho({ 'a': 1 }));
