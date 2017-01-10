var data = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  function hacerPeticionAjax (url, accion, callback, params) {
    // console.log(accion, ' => ', url, callback, params)
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4) { // 4 = "DONE"
        if (xhr.status === 200) { // 200 ="OK"
          if (accion === 'GET') {
            // le quitamos los esteticos <pre></pre>
            var datos = JSON.parse(xhr.responseText.slice(5).slice(0, -6));
            callback(datos);
          } else {
            callback();
          }
        } else {
          console.log('Error: ' + xhr.status);
        }
      }
    };
    xhr.open(accion, url);
    if (accion === 'GET') {
      xhr.send();
    } else if (accion !== 'GET') {
      xhr.setRequestHeader('Content-Type',
        'application/x-www-form-urlencoded; charset=UTF-8');
      xhr.send(params);
    }
  }

  return {

  };
}());
