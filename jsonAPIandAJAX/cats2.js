(function () {
  'use strict';
  /* code here */
  // https://www.freecodecamp.com/json/cats.json

  function iniciar () {
    getMessage.addEventListener('click', makeCorsRequest);
  }

  function getTitle (texto) {
    var ro = JSON.parse(texto);
    var nc = '';
    for (var i = 0; i < ro.length; i++) {
      nc += '<div class="gatos">';
      nc += 'ID = ' + ro[i].id + '<br/>';
      nc += '<img src="' + ro[i].imageLink + '"';
      nc += ' alt="' + ro[i].altText + '"/> <br/>';
      nc += 'Nombres = ' + ro[i].codeNames + '<br/>';
      nc += '</div> <br/><br/><br/>';
    }
    document.getElementsByClassName('message')[0].innerHTML = nc;
    console.log(texto);
    return texto;
  }

  // Create the XHR object.
  function createCORSRequest (method, url) {
    var xhr = new XMLHttpRequest();
    if ('withCredentials' in xhr) {
      // XHR for Chrome/Firefox/Opera/Safari.
      xhr.open(method, url, true);
    } else if (typeof XDomainRequest != 'undefined') {
      // XDomainRequest for IE.
      xhr = new XDomainRequest();
      xhr.open(method, url);
    } else {
      // CORS not supported.
      xhr = null;
    }
    return xhr;
  }

  // Make the actual CORS request.
  function makeCorsRequest () {
    var url = 'https://cors-test.appspot.com/test';
    // var url = 'http://updates.html5rocks.com'
    var xhr = createCORSRequest('GET', url);
    // xhr.setRequestHeader('X-Custom-Header', 'value')

    if (!xhr) {
      alert('CORS not supported');
      return;
    }
    // Response handlers.
    xhr.onload = function () {
      var text = xhr.responseText;
      var title = getTitle(text);
      alert('Response from CORS request to ' + url + ': ' + title);
    };
    xhr.onerror = function () {
      alert('Woops, there was an error making the request.');
    };
    xhr.send();
  }

  addEventListener('load', iniciar);
}());
