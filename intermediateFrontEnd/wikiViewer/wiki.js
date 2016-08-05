(function () {
  'use strict';
  /* code here */

  function init () {
    readRandom.addEventListener('click', getRandom);
    searchBt.addEventListener('click', getSearch);
    searchText.addEventListener('keydown', lookForIntro);
  }

  function getRandom () {
    console.log('Wiki Random');
    window.open('https://en.wikipedia.org/wiki/Special:Random');
  }

  function getSearch () {
    var texto = document.getElementById('searchText').value;
    console.log('Searching in Wiki --> ', texto);
    var script = document.createElement('script');
    var url1 = 'http://en.wikipedia.org/w/api.php';
    var url2 = '?action=query&list=search&format=json&srsearch=';
    var url3 = texto;
    var url4 = '&callback=writeRes';
    var url = 'https://en.wikipedia.org/w/api.php?action=query&prop=extracts&exlimit=max&format=json&exsentences=1&exintro=&explaintext=&generator=search&gsrlimit=10&gsrsearch=' + texto + url4;
    script.src = url;
    // document.body.appendChild(script)
    document.head.appendChild(script);
  }

  function lookForIntro (ev) {
    if (ev.keyCode === 13) {
      getSearch();
    }
  }

  addEventListener('load', init);
}());

function writeRes (data) {
  var results = data.query.pages;
  var pages = Object.keys(results);
  var wikilink = 'http://en.wikipedia.org/?curid=';
  var html = '';
  pages.forEach(function (page) {
    var title = results[page].title;
    var text = results[page].extract;
    var pagelink = wikilink + results[page].pageid;
    // console.log(pagelink)
    html += '<div class="result text-xs-center"><a href="' + pagelink + '" >' + title + '</a>';
    html += '<p>' + text + '</p>';
    html += '</div><br/>';
    if (html !== undefined) {
      document.getElementById('results').innerHTML = html;
      console.log('undefined');
    }
  });
}
