(function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Inicio');
    var clickType = document.getElementsByClassName('card');
    for (var i = 0; i < clickType.length; i++) {
      clickType.item(i).addEventListener('click', clickCal);
    // console.log(clickType.item(i).innerText)
    }
  }

  function clickCal (ev) {
    console.log(ev.target.innerText);
    var action = ev.target.innerText;
    document.getElementById('display').innerText = action;
  }

  addEventListener('load', init);
}());
