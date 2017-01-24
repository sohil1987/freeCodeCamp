(function () {
  'use strict';
  /* code here */

  function init () {
    loadQuote();
    newQuote.addEventListener('click', loadQuote);
    tweet.addEventListener('click', retweet);
  }

  function loadQuote () {
    var script = document.createElement('script');
    script.src =
      'http://api.forismatic.com/api/1.0/?method=getQuote&format=jsonp&lang=en&jsonp=writeRes';
    // document.body.appendChild(script)
    document.head.appendChild(script);
  }

  function retweet () {
    var text = document.getElementById('text').innerHTML; // .split(' ').join('')
    var author = document.getElementById('author').innerHTML;
    var pegote = encodeURIComponent('"' + text + '" ' + author);
    var url = ('https://twitter.com/intent/tweet?text=' + pegote);
    console.log(url);
    window.open(url);
  }

  addEventListener('load', init);
}());

function writeRes (data) {
  'use strict';
  document.getElementById('text').innerHTML = data.quoteText;
  document.getElementById('author').innerHTML = data.quoteAuthor;
  console.log(data);
}
/*

http://api.forismatic.com/api/1.0/?method=getQuote&format=jsonp&lang=en&jsonp=writeRes
*/
