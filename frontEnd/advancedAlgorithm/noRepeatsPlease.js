/*jslint node: true */
'use strict';

// https://en.wikipedia.org/wiki/Heap%27s_algorithm

function permAlone (str) {
  str = str.split('');
  var size = str.length;
  var permutations = [];

  generate(size, permutations, str);

  var repeated = /(\w)\1+/g;
  var result = permutations.filter(function (string) {
    return !string.match(repeated);
  });
  return result.length;
}

function generate (size, permutations, str) {
  if (size == 1) {
    permutations.push(str.join(''));
  } else {
    for (var i = 0; i < size; ++i) {
      generate(size - 1, permutations, str);
      if (size % 2 === 0) {
        swap(i, size - 1, str);
      } else {
        swap(0, size - 1, str);
      }
    }
  // generate(size - 1, permutations, str) // Why i have to uncomment this ??
  }
  return permutations;
}

function swap (a, b, str) {
  var aux = str[a];
  str[a] = str[b];
  str[b] = aux;
}

console.log(permAlone('aab'));
console.log(permAlone('aaa'));
console.log(permAlone('aabb'));
console.log(permAlone('abcdefa'));
console.log(permAlone('abfdefa'));
console.log(permAlone('zzzzzzzz'));
console.log(permAlone('a'));
console.log(permAlone('aaab'));
console.log(permAlone('aaabb'));
