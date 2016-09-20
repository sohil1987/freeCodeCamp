(function () {
  'use strict';
  /* code here */

  function iniciar () {
    getMessage.addEventListener('click', cambiarTexto);
  }

  function cambiarTexto () {
    console.log('BOTON PRETADO');
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'https://www.freecodecamp.com/json/cats.json');
    xhr.send();
    xhr.addEventListener('load', function () {
      if (xhr.status === 200) {
        var ro = JSON.parse(xhr.responseText);
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
        console.log(xhr.responseText);
      }
    });
  }

  addEventListener('load', iniciar);
}());

// SIN CORS

/*
To allow client-side JavaScript applications to fetch information from MediaWiki APIs, add the following header to API responses, allowing the response to be read by an application running on a different domain:

Access-Control-Allow-Origin: *

*/
