function makeFriendlyDates (arr) {
  var beginDate = new Date(arr[0]);
  var beginDay = beginDate.getDate();
  var beginMonth = beginDate.getMonth() + 1;
  var beginYear = beginDate.getFullYear();
  var endDate = new Date(arr[1]);
  var endDay = endDate.getDate();
  var endMonth = endDate.getMonth() + 1;
  var endYear = endDate.getFullYear();

  var beginTime = beginDate.getTime();
  var endTime = endDate.getTime();
  var mSecYear = 1000 * 60 * 60 * 24 * 365;

  var currentYear = new Date().getFullYear();

  var result = [];
  var a, b = '';

  if (beginTime === endTime) {
    a = getMonth(beginMonth) + ' ' + getDay(beginDay) + ', ' + beginYear;
  } else if (beginYear === endYear && beginMonth === endMonth) {
    a = getMonth(beginMonth) + ' ' + getDay(beginDay);
    b = getDay(endDay);
  } else if (((beginYear === endYear || beginYear === endYear - 1) || beginYear === endYear) && beginMonth != endMonth) {
    a = getMonth(beginMonth) + ' ' + getDay(beginDay) + ', ' + beginYear;
    b = getMonth(endMonth) + ' ' + getDay(endDay);
    if (beginYear === currentYear) {
      a = getMonth(beginMonth) + ' ' + getDay(beginDay);
    }
  } else {
    console.log('PLUS');
    a = getMonth(beginMonth) + ' ' + getDay(beginDay) + ', ' + beginYear;
    b = getMonth(endMonth) + ' ' + getDay(endDay) + ', ' + endYear;
    if ((endTime - beginTime) / mSecYear < 1) {
      b = getMonth(endMonth) + ' ' + getDay(endDay);
    }
  }

  if (a) result.push(a);
  if (b) result.push(b);
  return result;
}

function getMonth (num) {
  var month = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
  return month[num - 1];
}

function getDay (num) {
  switch (num) {
    case 1:
    case 21:
    case 31:
      return num + 'st';
    case 2:
    case 22:
      return num + 'nd';
    case 3:
    case 23:
      return num + 'rd';
    default:
      return num + 'th';
  }
}

// console.log('RES -->', makeFriendlyDates(['2016-07-01', '2016-07-04']))
// console.log('RES -->', makeFriendlyDates(['2016-12-01', '2017-02-03']))
// console.log('RES -->', makeFriendlyDates(['2016-12-01', '2018-02-03']))
// console.log('RES -->', makeFriendlyDates(['2017-03-01', '2017-05-05']))
// console.log('RES -->', makeFriendlyDates(['2018-01-13', '2018-01-13']))
console.log('RES -->', makeFriendlyDates(['2022-09-05', '2023-09-04']));
// console.log('RES -->', makeFriendlyDates(['2022-09-05', '2023-09-05']))
