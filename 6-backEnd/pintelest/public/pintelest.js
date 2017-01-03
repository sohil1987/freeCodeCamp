var app = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  function init () {
    console.log('Inicio PINTELEST');
    if (user)   console.log('USUARIO', user.username);
    if (active) console.log('ACTIVE', active);
    document.getElementById('tb').addEventListener('click', handleTwitter);
  }

  function handleTwitter () {
    if (user) {
      window.location = 'http:/logout';
    } else {
      window.location = 'http:/login/twitter';
    }
  }

  return {
    inicio: init
  };
}());

window.addEventListener('load', app.inicio);
