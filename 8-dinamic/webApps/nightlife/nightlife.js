var app = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  function init () {
    console.log('Inicio');
  }

  return {
    inicio: init
  };
}());

window.addEventListener('load', app.inicio);
