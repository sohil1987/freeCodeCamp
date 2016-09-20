(function () {
  'use strict';
  /* code here */

  function init () {
    local.addEventListener('click', getLocation);
    change.addEventListener('click', changeCF);
    getLocation();
  }

  function getLocation () {
    console.log('BEGIN');
    var script = document.createElement('script');
    var url = 'http://ipinfo.io/json?callback=writeRes';
    // var url = 'http://ip-api.com/json/?callback=writeRes'
    // var url = 'https://api.ipify.org?format=jsonp?callback=writeRes'
    script.src = url;
    // document.body.appendChild(script)
    document.head.appendChild(script);
    console.log('END');
  }

  function changeCF () {
    var actual = document.getElementById('change').innerHTML;
    var temp = document.getElementById('temp').innerHTML;
    if (actual === 'ºC') {
      actual = 'ºK';
      temp = (temp * (9 / 5) + 32).toFixed(1);
    } else if (actual === 'ºK') {
      actual = 'ºC';
      temp = ((temp - 32) * (5 / 9)).toFixed(1);
    }
    document.getElementById('change').innerHTML = actual;
    document.getElementById('temp').innerHTML = temp;
  }

  addEventListener('load', init);
}());

function writeRes (data) {
  'use strict';
  console.log('writeRes');
  var lat = data.loc.split(',')[0];
  var lon = data.loc.split(',')[1];
  console.log(lat, lon);
  getWeather(lat, lon);
}

function getWeather (lat, lon) {
  // lat = 42.4667
  // lon = -2.4500
  console.log(lat, ' | ', lon);
  var key = '&APPID=a6ccf57f755e7e5bbc51b4c2c1a052d4';
  var xhr = new XMLHttpRequest();
  var url =
  'http://api.openweathermap.org/data/2.5/weather?lat=' + lat + '&lon=' +
    lon + key + '&callback=writeTemp';
  console.log(url);
  var script = document.createElement('script');
  script.src = url;
  // document.body.appendChild(script)
  document.head.appendChild(script);
}

function writeTemp (data) {
  var city = data.name + ', ' + data.sys.country;
  var temp = (data.main.temp - 273.16).toFixed(1);
  var weather = data.weather[0].description;
  var icon = 'http://openweathermap.org/img/w/' + data.weather[0].icon + '.png';
  icon = '<img src="' + icon + '">';
  document.getElementById('city').innerHTML = city;
  document.getElementById('temp').innerHTML = temp;
  document.getElementById('weather').innerHTML = weather;
  document.getElementById('icon').innerHTML = icon;
}
