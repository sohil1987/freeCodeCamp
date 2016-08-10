/*jslint node: true */
'use strict';

function spinalCase (str) {
  // "It's such a fine line between stupid, and clever."
  // --David St. Hubbins

  return str.replace(/(?!^)([A-Z])/g, ' $1')
    .replace(/[_\s]+(?=[a-zA-Z])/g, '-').toLowerCase();
}

console.log(spinalCase('This Is Spinal Tap'));
console.log(spinalCase('thisIsSpinalTap'));
console.log(spinalCase('The_Andy_Griffith_Show'));
