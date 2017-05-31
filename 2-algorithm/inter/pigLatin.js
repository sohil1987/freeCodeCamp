/*jshint node: true */
'use strict';

const vowels = 'aeiou';

function translatePigLatin (str) {
  if (vowels.includes(str[0])) { // word begins with a vowel
    return str + 'way';
  }
  let aux = '';
  for (let i = 0; i < str.length; i++) {
    if (!vowels.includes(str[i])) {
      aux += str[i];
    } else {
      return str.slice(aux.length, str.length) + aux + 'ay';
    }
  }
  return str + 'ay'; // words without vowels
}

// console.log(translatePigLatin('consonant'))
// console.log(translatePigLatin('glove'))
console.log(translatePigLatin('glvsssss'));
