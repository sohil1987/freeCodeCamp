var guest = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  function init () {
    console.log('Hi From Client JS in guest', go);
  }

  return {
    inicio: init
  };
}());

var login = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  function init () {
    console.log('Hi From Client JS in login', go);
    if (go === 'INVALID PASSWORD') {
      alert('INVALID PASSSWORD');
    }
  }

  return {
    inicio: init
  };
}());

var logged = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  function init () {
    console.log('Hi From Client JS in logged', go);
    var user = app.getURLParameter('user');
    var search = document.getElementsByClassName('searchBox')[0].placeholder;
    console.log('User', user, ' - Search ', search);
    var aux = document.getElementsByClassName('voteLinks');
    for (var i = 0; i < aux.length; i++) {
      aux[i].href += '&user=' + user + '&search=' + search;
    }
  }

  return {
    inicio: init
  };
}());

var app = (function () {
  'use strict';
  /* jshint node: true */
  /* code here */

  // https://stackoverflow.com/questions/11582512/how-to-get-url-parameters-with-javascript  
  function getURLParameter (name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search) || [null, ''])[1].replace(/\+/g, '%20')) || null;
  }

  function getBaseUrl () {
    return ('/freecodecamp/v1/apps/');
  }

  return {
    getURLParameter: getURLParameter,
    getBaseUrl: getBaseUrl
  };
}());
