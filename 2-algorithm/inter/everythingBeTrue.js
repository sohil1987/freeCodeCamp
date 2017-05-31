/*jshint node: true */
'use strict';

function truthCheck (collection, pre) {
  // Is everyone being true?
  let cont = 0;
  for (let i in collection) {
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

console.log(truthCheck([{'user': 'Tinky-Winky', 'sex': 'male'}, {'user': 'Dipsy', 'sex': 'male'}, {'user': 'Laa-Laa', 'sex': 'female'}, {'user': 'Po', 'sex': 'female'}], 'sex'));
