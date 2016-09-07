(function () {
  'use strict';
  /* code here */
  // https://www.freecodecamp.com/json/cats.json
  // https://cors-test.appspot.com/test

  function iniciar () {
    // getMessage.addEventListener('click', makeCorsRequest)
    getCORS('https://www.freecodecamp.com/json/cats.json', function (request) {
      var response = request.currentTarget.response || request.target.responseText;
      console.log(response);
    });
  }

  function getTitle (request) {
    var texto;
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

  function getCORS (url, success) {
    var xhr = new XMLHttpRequest();

    xhr.open('GET', url);
    xhr.setRequestHeader('Content-Type', 'application/json');

    xhr.onload = success;
    xhr.send();
    return xhr;
  }

  addEventListener('load', iniciar);
}());
