var app = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  function init () {
    console.log('Inicio PINTELEST');
    if (active) console.log('ACTIVE', active);
    document.getElementById('tb').addEventListener('click', handleTwitter);
    if (user) {
      console.log('USUARIO', user.username);
      createVotingEvents();
    }
    checkImages();
  }

  function createVotingEvents () {
    if (pics) {
      console.log('PICS', pics.length);
      for (let i = 0; i < pics.length; i++) {
        // console.log(pics[i])
        let el = 'vote' + pics[i].idPic;
        document.getElementById(el).addEventListener('click', checkIfVoted);
      }
    }
  }

  function handleTwitter () {
    if (user) {
      window.location = '/logout';
    } else {
      window.location = '/login/twitter';
    }
  }

  function checkIfVoted (e) {
    let picId = e.target.id.slice(4);
    let userId = user.id;
    // console.log('User... ', userId, ' --- Pic... ', picId)
    window.location = '/vote/' + userId + '/' + picId;
  }

  function checkImages () {
    var array = document.getElementsByClassName('cardImage');
    for (let i = 0; i < array.length; i++) {
      urlPicExists(array[i]);
    }
  }

  function urlPicExists (pic) {
    // console.log('testing ...', pic.src)
    var imageData = new Image();
    imageData.onload = function () {};
    imageData.onerror = function () {
      pic.src = './../images/photoNot.png';
    };
    imageData.src = pic.src;
  }

  return {
    inicio: init
  };
}());

window.addEventListener('load', app.inicio);
