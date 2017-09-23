/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  const urlIP = 'https://freegeoip.net/json/?callback=app.getLoc';
  const key = '&APPID=a6ccf57f755e7e5bbc51b4c2c1a052d4';
  let res = {
    city: '',
    lat: '',
    lon: '',
    temp: '',
    weather: '',
    icon: ''
  };

  function init () {
    console.log('Init Local Weather');
    let script = document.createElement('script');
    script.src = urlIP;
    document.body.appendChild(script);
    document.getElementsByClassName('temp')[0].addEventListener('click', changeType);
  }

  function getLoc (data) {
    console.log('LOC', data);
    res.city = data.city + ', ' + data.country_name;
    res.lat = data.latitude;
    res.lon = data.longitude;
    document.getElementById('city').innerText = res.city;
    document.getElementById('weather').innerText = 'Disable https';
    document.getElementById('temp').innerText = 'for this page';
    document.getElementById('type').innerText = '';
    document.getElementsByClassName('data')[0].style.color = '#375EAB';
    getWeatherJson();
  }

  function getWeatherJsonP () {
    const urlWeather =
    'http://api.openweathermap.org/data/2.5/weather?lat=' + res.lat + '&lon=' +
      res.lon + key + '&callback=app.writeWeather';
    let script = document.createElement('script');
    script.src = urlWeather;
    document.body.appendChild(script);
  }

  function getWeatherJson () {
    const urlWeather =
    'http://api.openweathermap.org/data/2.5/weather?lat=' + res.lat + '&lon=' +
      res.lon + key;
    getAjaxData(urlWeather, writeWeather);
  }

  function writeWeather (data) {
    res.temp = (data.main.temp - 273.16).toFixed(1);
    res.weather = data.weather[0].description;
    res.icon = 'http://openweathermap.org/img/w/' + data.weather[0].icon + '.png';
    res.icon = '<img src="' + res.icon + '">';
    document.getElementById('city').innerText = res.city;
    document.getElementById('temp').innerText = res.temp;
    if (res.icon === undefined) {
      document.getElementById('weather').innerText = res.weather;
    } else {
      document.getElementById('weather').innerHTML = res.icon;
    }
    document.getElementsByClassName('data')[0].style.color = '#fafafa';
  }

  function changeType () {
    var actual = document.getElementById('type').innerHTML;
    if (actual === 'ºC') {
      actual = 'ºK';
      res.temp = (res.temp * (9 / 5) + 32).toFixed(1);
    } else if (actual === 'ºK') {
      actual = 'ºC';
      res.temp = ((res.temp - 32) * (5 / 9)).toFixed(1);
    }
    document.getElementById('type').innerText = actual;
    document.getElementById('temp').innerText = res.temp;
  }

  function getAjaxData (urlData, callback) {
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4) { // 4 = "DONE"
        if (xhr.status === 200) { // 200 ="OK"
          console.log(JSON.parse(xhr.responseText));
          callback(JSON.parse(xhr.responseText));
        } else {
          console.log('Error: ' + xhr.status);
        }
      }
    };
    xhr.open('GET', urlData); // add false to synchronous request
    xhr.send();
  }

  return {
    init: init,
    getLoc: getLoc,
    writeWeather: writeWeather
  };
}());

window.addEventListener('load', app.init);
