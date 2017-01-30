/*jslint node: true */
'use strict';

function truthCheck (collection, pre) {
  // Is everyone being true?
  var cont = 0;
  for (var i in collection) {
    // console.log(collection[i].pre)
    if (collection[i].hasOwnProperty(pre) && Boolean(collection[i][pre])) {
      cont++;
    }
  }
  if (cont === collection.length) {
    return true;
  } else {
    return false;
  }
}

console.log(truthCheck([{
  'name': 'Pete',
  'onBoat': true
}, {
  'name': 'Repeat',
  'onBoat': true,
  'alias': 'Repete'
}, {
  'name': 'FastFoward',
  'onBoat': true
}], 'onBoat'));
