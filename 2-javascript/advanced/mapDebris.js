function orbitalPeriod (arr) {
  var GM = 398600.4418;
  var radius = 6367.4447;
  var result = [];

  for (var index in arr) {
    result.push({
      name: arr[index].name,
      orbitalPeriod: getOP(arr[index].avgAlt, GM, radius)
    });
  }

  return result;
}

function getOP (alt, GM, radius) {
  var op = 2 * Math.PI * Math.sqrt((Math.pow(alt + radius, 3) / GM));
  return Math.round(op);
}

console.log(orbitalPeriod([{
  name: 'sputnik',
  avgAlt: 35873.5553
}]));

console.log(orbitalPeriod([{
  name: 'iss',
  avgAlt: 413.6
}, {
  name: 'hubble',
  avgAlt: 556.7
}, {
  name: 'moon',
  avgAlt: 378632.553
}]));
