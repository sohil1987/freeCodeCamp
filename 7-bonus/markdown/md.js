(function () {
  'use strict';
  /* code here */
  //  https://cdnjs.com/libraries/marked

  function init () {
    console.log('Init');
    window.addEventListener('resize', adjustSize);
    document.getElementById('content').addEventListener('input', previewText);
    adjustSize();
    initialText();
  }

  function initialText () {
    var text =
    "It's very easy to make some words **bold** and other words *italic* with Markdown.\n\nYou can even [link to my Freecodecamp projects!](https://brusbilis.com/freecodecamp)\n";

    document.getElementById('content').placeholder = text;
    document.getElementById('preview').innerHTML = marked(text);
  }

  function previewText (ev) {
    document.getElementById('preview').innerHTML = marked(ev.target.value);
  }

  function adjustSize () {
    var h = window.innerHeight;
    document.getElementById('content').style.height = h - 120 + 'px';
  }

  window.addEventListener('load', init);
}());
