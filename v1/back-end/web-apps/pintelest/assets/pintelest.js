/*jshint node: true*/

const app = (function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Init PINTELEST');
  }

  return {
    inicio: init
  };
}());

window.addEventListener('load', app.inicio);
