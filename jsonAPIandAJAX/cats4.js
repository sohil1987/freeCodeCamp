(function () {
  'use strict';
  /* code here */

  function iniciar () {
    getMessage.addEventListener('click', cambiarTexto);
  }

  function cambiarTexto () {
    console.log('COMIENZA');
    // GET ejemplo
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'https://cors-test.appspot.com/test', true);
    // xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.onreadystatechange = function () {
      if (this.status == 200 && this.readyState == 4) {
        console.log('OK');
        console.log('response: ' + this.responseText);
      }
    };
    xhr.send();

    // PUT ejemplo
    var req = new XMLHttpRequest();

    req.open('PUT', 'https://maptestapi.apispark.net/v1/maps/4', true);
    req.onreadystatechange = function () {
      if (req.readyState === 4) {
        console.log(req.responseText);
      }
    };
    req.setRequestHeader('Content-type', 'application/json');
    req.send('{ "id": "4", "name": "Meteorites", "type": "d3js", (…)');

  /*
      // GET ejemplo con API sin activar CORS
      var req = new XMLHttpRequest()

      req.open('GET', 'https://www.freecodecamp.com/json/cats.json', true)
      req.onreadystatechange = function () {
        if (req.readyState === 4) {
          console.log(req.responseText)
        }
      }
      req.withCredentials = true
      req.setRequestHeader('Content-type', 'application/json')
      req.send()
      */
  }

  addEventListener('load', iniciar);
}());
