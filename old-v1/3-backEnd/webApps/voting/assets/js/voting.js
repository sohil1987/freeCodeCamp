var guest = (function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Hi From Client JS in guest', go);
    if (go) {
      app.drawCharts();
    } else {
      alert('You already voted on that poll');
      // window.location = '/voting/guest/'
      window.location.assign('/voting/guest/');
    }
  }

  return {
    init: init
  };
}());

var logged = (function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Hi From Client JS in logged', go);
    var user = app.getURLParameter('user');
    if (go) {
      console.log('LETS GO');
      app.drawCharts();
    } else {
      alert(user + ' has already voted on that poll');
      // window.location = '/voting/logged/?user=' + user
      window.location.assign('/voting/logged/?user=' + user);
    }
    document.getElementById('whichUser').innerHTML = user;
    document.getElementById('linkNewPoll').href += '?user=' + user;
  }

  function addOption (ev) {
    var poll = ev[0].value; // voting OptionID from first option
    if (ev.value === '0') { // add option choosed
      var newOption = prompt('Enter new option');
      if (newOption !== undefined && newOption !== '') {
        var url = '/voting/newColumn/?poll=' + poll;
        url += '&newOption=' + newOption;
        window.location.assign(url);
        console.log(url);
      }
    }
  }

  return {
    init: init,
    addOpt: addOption
  };
}());

var login = (function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Hi From Client JS in login', go);
    if (go === 'INVALID PASSWORD') {
      alert('INVALID PASSSWORD');
    }
  }

  return {
    init: init
  };
}());

var newPoll = (function () {
  'use strict';
  /* code here */

  function init () {
    console.log('Hi From Client JS in newPoll', go);
    var user = app.getURLParameter('user');
    document.getElementById('whichUser').innerHTML = user;
  }

  return {
    init: init
  };
}());

var app = (function () {
  'use strict';
  /* code here */

  function drawCharts () {
    var backgroundColor = [];
    for (let chart = 0; chart < go.length; chart++) {
      var data = {
        labels: [],
        datasets: [{
          data: [],
          backgroundColor: []
        }]
      };
      for (let option = 0; option < go[chart].Options.length;option++) {
        data.labels.push(go[chart].Options[option].Option);
        data.datasets[0].data.push(go[chart].Options[option].NumVotes);
        data.datasets[0].backgroundColor.push(getRandomColor());
      }
      var aux = document.getElementsByClassName('card' + go[chart].PollID);
      var ctx;
      for (let i = 0; i < aux.length; i++) {
        ctx = aux[i];
        var myPieChart = new Chart(ctx, {
          type: 'pie',
          data: data,
          options: {
            legend: {
              labels: {
                fontSize: 12,
                padding: 3,
                boxWidth: 10
              }
            }
          }
        });
      }
    }
  }

  function getRandomColor () {
    var letters = '0123456789ABCDEF'.split('');
    var color = '#';
    for (var i = 0; i < 6; i++) {
      color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
  }

  // https://stackoverflow.com/questions/11582512/how-to-get-url-parameters-with-javascript  
  function getURLParameter (name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search) || [null, ''])[1].replace(/\+/g, '%20')) || null;
  }

  return {
    drawCharts: drawCharts,
    getURLParameter: getURLParameter

  };
}());
