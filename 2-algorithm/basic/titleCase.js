/*jshint node: true */
'use strict';

function titleCase (str) {
  str = str.toLowerCase().split(' ');
  let result = [];
  str.forEach(function (word, pos) {
    result.push(capitalizeWord(word));
  });
  return result.join(' ');
}

function capitalizeWord (word) {
  word = word.split('');
  word[0] = word[0].toUpperCase();
  word = word.join('');
  return word;
}

console.log(titleCase("I'm a little tea pot"));
