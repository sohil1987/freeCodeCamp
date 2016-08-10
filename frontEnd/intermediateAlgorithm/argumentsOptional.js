/*jslint node: true */
'use strict';

function addTogether () {
  var howMany = arguments.length;
  var result = 0;
  for (var i = 0; i < howMany; i++) {
    if (Number.isInteger(arguments[i])) {
      // console.log("I'm an integer")
      result += arguments[i];
    } else {
      return undefined;
    }
  }
  if (howMany < 2) {
    var val = arguments[0];
    var sumVals = function (num) {
      if (Number.isInteger(num)) {
        return (num + val);
      } else {
        return undefined;
      }
    };
    return sumVals;
  } else {
    return result;
  }
}

console.log(addTogether(2, 3));
console.log(addTogether('http://bit.ly/IqT6zt'));
console.log(addTogether(2)(3));
