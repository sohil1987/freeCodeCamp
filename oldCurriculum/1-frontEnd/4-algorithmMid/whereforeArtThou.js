/*jslint node: true */
'use strict';

function whatIsInAName (collection, source) {
  var arr = [];
  // Only change code below this line
  var wanted = Object.keys(source);
  console.log(wanted);
  for (var i = 0; i < collection.length;i++) {
    console.log('Estudiando --> ', collection[i]);
    var insert = true;
    for (var j = 0; j < wanted.length; j++) {
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

// console.log(whatIsInAName([{ first: 'Romeo', last: 'Montague' }, { first: 'Mercutio', last: null }, { first: 'Tybalt', last: 'Capulet' }], { last: 'Capulet' }))
console.log(whatIsInAName([{ 'a': 1, 'b': 2 }, { 'a': 1 }, { 'a': 1, 'b': 2, 'c': 2 }], { 'a': 1, 'c': 2 }));
