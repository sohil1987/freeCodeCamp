/*jshint node: true */
'use strict';

function pairElement (str) {
  let sol = [];
  for (let i = 0; i < str.length;i++) {
    switch (str[i]) {
      case 'A':
        sol.push(['A', 'T']);
        break;
      case 'T':
        sol.push(['T', 'A']);
        break;
      case 'C':
        sol.push(['C', 'G']);
        break;
      case 'G':
        sol.push(['G', 'C']);
    }
  }
  return sol;
}

console.log(pairElement('GCG'));
