/*jshint node: true */
'use strict';

function whatIsInAName (collection, source) {
  let arr = [];
  // Only change code below this line
  let wanted = Object.keys(source);
  for (let i = 0; i < collection.length;i++) {
    let insert = true;
    for (let j = 0; j < wanted.length; j++) {
      if (collection[i].hasOwnProperty(wanted[j])) {
        if (collection[i][wanted[j]] !== source[wanted[j]]) {
          insert = false;
        }
      }
      if (!collection[i].hasOwnProperty(wanted[j])) {
        insert = false;
      }
    }
    if (insert) {
      arr.push(collection[i]);
    }
  }
  // Only change code above this line
  return arr;
}

whatIsInAName([{ first: 'Romeo', last: 'Montague' }, { first: 'Mercutio', last: null }, { first: 'Tybalt', last: 'Capulet' }], { last: 'Capulet' });
