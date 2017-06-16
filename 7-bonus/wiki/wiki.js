/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Init Wikipedia Viewer');
    createClicks();
  }

  function clickAction (e) {
    const action = e.target.id;
    switch (action) {
      case 'search':
        getSearch();
        break;
      case 'random':
        goRandom();
        break;
      default:
        console.log('Error, not action defined');
    }
  }

  function getSearch () {
    const limit = 10;
    const searching = document.getElementById('searchBox').value;
    const url = 'https://en.wikipedia.org/w/api.php?action=query&prop=extracts&exlimit=max&format=json&exsentences=1&exintro=&explaintext=&generator=search&gsrlimit=' + limit + '&gsrsearch=' + searching + '&callback=app.writeRes';
    const script = document.createElement('script');
    script.src = url;
    document.body.appendChild(script);
  }

  function writeRes (data) {
    const wikilink = 'http://en.wikipedia.org/?curid=';
    let results = data.query.pages;
    let pages = Object.keys(results);
    let html = '';
    pages.forEach(function (page) {
      let title = results[page].title;
      let text = results[page].extract;
      let pagelink = wikilink + results[page].pageid;
      html += '<div class="article">';
      html += '<div class="head">';
      html += '<a href= "' + pagelink + '" target="_blank"> ' + title + '</a>';
      html += '</div>';
      html += '<div class="text">';
      html += '<span>' + text + '<span>';
      html += '</div>';
      html += '</div>';
      if (html !== undefined) {
        document.getElementsByClassName('articles')[0].innerHTML = html;
      }
    });
  }

  function goRandom () {
    window.open('https://en.wikipedia.org/wiki/Special:Random');
  }

  function createClicks () {
    const clickType = document.getElementsByClassName('action');
    for (let i = 0; i < clickType.length; i++) {
      clickType.item(i).addEventListener('click', clickAction);
    // console.log(clickType.item(i).innerText)
    }
    searchBox.addEventListener('keydown', function (e) {
      const searching = document.getElementById('searchBox').value;
      if (e.keyCode === 13 && searching !== '') {
        getSearch();
      }
    });
  }

  return {
    init: init,
    writeRes: writeRes
  };
}());

window.addEventListener('load', app.init);
