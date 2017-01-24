(function () {
  'use strict';
  /* code here */

  // var channels = ['OgamingSC2' , 'freecodecamp', 'storbeck', 'terakilobyte', 'habathcx', 'RobotCaleb', 'thomasballinger', 'noobs2ninjas', 'beohoff', 'brunofin', 'comster404', 'test_channel', 'cretetion', 'sheevergaming', 'TR7K', 'ESL_SC2']

  var channels = ['OgamingSC2', 'ESL_SC2'];

  function init () {
    btnAll.addEventListener('click', listChannels);
    btnOnline.addEventListener('click', listChannels);
    btnOffline.addEventListener('click', listChannels);
  }

  function listChannels (ev) {
    var type = ev.target.id;
    for (var i = 0; i < channels.length; i++) {
      var xhr = new XMLHttpRequest();
      var url = 'https://api.twitch.tv/kraken/streams/';
      url = url + channels[i];
      xhr.open('GET', url, true);
      xhr.send();
      xhr.onload = function () {
        if (xhr.readyState === 4) {
          if (xhr.status === 200) {
            console.log(JSON.parse(xhr.responseText).stream.channel.display_name + ' OK');
          // printChannel(xhr.responseText, type)
          } else {
            console.error(xhr.statusText);
          }
        }
      };
    }
  }

  function printChannel (data, type) {
    data = JSON.parse(data);
    var name = data.stream.channel.display_name;
    var game = data.stream.game;
    var logo = data.stream.channel.logo;
    var url = data.stream.channel.url;
    var status = data.stream.channel.status;
    // var html = ''
    var html = document.getElementById('results').innerHTML;
    html += '<div class="oneResult row text-xs-center">';
    html += '<div class="col-xs-3">';
    html += '<img class="logo" alt="logo" src="' + logo + '"/>';
    html += '</div>';
    html += '<div class="col-xs-3">';
    html += '<a href="' + url + '">' + name + '</a>';
    html += '</div>';
    html += '<div class="col-xs-6">';
    html += game + ' --> ' + status;
    html += '</div>';
    html += '</div>';
    html += '';
    document.getElementById('results').innerHTML = html;
  }

  addEventListener('load', init);
}());

// Esto es con JSON pero no se porque no termina de ir bien
