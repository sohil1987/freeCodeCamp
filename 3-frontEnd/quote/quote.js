const app = (function () {
  'use strict';
  /* code here */

  const url = 'https://api.forismatic.com/api/1.0/?method=getQuote&format=jsonp&lang=en&jsonp=app.writeRes';

  function init () {
    console.log('Init Random Quote Machine');
    let t = document.getElementById('tweet-quote');
    let n = document.getElementById('new-quote');
    t.addEventListener('click', tweetQuote);
    n.addEventListener('click', newQuote);
    newQuote();
  }

  function tweetQuote () {
    let text = document.getElementById('text').textContent;
    let author = document.getElementById('author').textContent;
    let aux = encodeURIComponent(text + ' ' + author);
    let url = ('https://twitter.com/intent/tweet?text=' + aux);
    window.open(url);
  }

  function newQuote () {
    let script = document.createElement('script');
    script.src = url;
    document.body.appendChild(script);
  }

  function writeRes (data) {
    // console.log(data)
    document.getElementById('text').textContent = '"' + data.quoteText + '"';
    document.getElementById('author').textContent = '- ' + data.quoteAuthor + ' -';
  }

  return {
    init: init,
    writeRes: writeRes
  };
}());

addEventListener('load', app.init);
