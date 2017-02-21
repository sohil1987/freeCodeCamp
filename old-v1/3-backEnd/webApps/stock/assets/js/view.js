var views = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  function init () {
    console.log('Inicio VIEWS');
  }

  function createBoxes () {
    if (datos.names.length === 0 || datos.names[0] === '') {
      return;
    }
    document.getElementById('list').innerHTML = '';
    document.getElementById('list').innerHTML = getHTML();
    var aux = document.getElementsByClassName('infoSpan');
    for (var i = 0; i < aux.length; i++) {
      document.getElementById(datos.names[i]).addEventListener('click', app.delStock);
    }
  }

  function getHTML () {
    var res = '';
    for (var i = 0; i < datos.names.length; i++) {
      res += '<div class="stock">';
      res += '  <div class="title">';
      res += '    <div class="titleSpan">';
      res += datos.names[i];
      res += '    </div>';
      res += '    <div id="' + datos.names[i] + '" class="infoSpan" > ' + 'X' + '</div > ';
      res += '  </div>';
      res += '</div>';
    }
    return res;
  }

  return {
    inicio: init,
    createBoxes: createBoxes
  };
}());
