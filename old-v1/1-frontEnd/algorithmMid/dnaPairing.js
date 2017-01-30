/*jslint node: true */
'use strict';

function pairElement (str) {
  str = str.split('');
  var res = [];
  var aux = [];
  for (var i = 0; i < str.length; i++) {
    aux[0] = str[i];
    aux[1] = pair(str[i]);
    res.push(aux);
    aux = [];
  }
  return res;
}

function pair (str) {
  if (str === 'C') {return 'G';}
  if (str === 'G') {return 'C';}
  if (str === 'A') {return 'T';}
  if (str === 'T') {return 'A';}
}

console.log(pairElement('ATCGA'));
