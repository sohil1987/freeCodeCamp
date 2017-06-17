/*jshint node: true */

const app = (function () {
  'use strict';
  /* code here */

  let result = {};
  let html = '';
  let action;
  const channels = ['OgamingSC2', 'freecodecamp', 'storbeck', 'terakilobyte', 'habathcx', 'RobotCaleb', 'thomasballinger', 'noobs2ninjas', 'beohoff', 'brunofin', 'comster404', 'test_channel', 'cretetion', 'sheevergaming', 'TR7K', 'ESL_SC2'];

  function init () {
    console.log('Init Twitch API');
    all.addEventListener('click', listChannels);
    on.addEventListener('click', listChannels);
    off.addEventListener('click', listChannels);
  }

  function listChannels (e) {
    action = e.target.id;
    html = '';
    for (let i = 0; i < channels.length; i++) {
      const url = 'https://wind-bow.gomix.me/twitch-api/streams/' + channels[i] + '?callback=app.printChannel';
      const script = document.createElement('script');
      script.src = url;
      document.body.appendChild(script);
    }
  }

  function printChannel (data) {
    // console.log('1 --> ', data)
    if (data.stream !== null && data.stream !== undefined) {
      if (action === 'all' || action === 'on') {
        getDateset(data);
      }
    } else if (data.stream === null) {
      if (action === 'off' || action === 'all') {
        let channel = data._links.channel.split('/')[data._links.channel.split('/').length - 1];
        const url = 'https://wind-bow.gomix.me/twitch-api/channels/' + channel + '?callback=app.getDateset2';
        const script = document.createElement('script');
        script.src = url;
        document.body.appendChild(script);
      }
    }
  }

  function getDateset (data) {
    let logo = data.stream.channel.logo ||
      'https://dummyimage.com/50x50/ecf0e7/5c5457.jpg&text=0x3F';
    let name = data.stream.channel.display_name;
    let url = data.stream.channel.url;
    let game = data.stream.game;
    let status = data.stream.channel.status;
    // console.log(logo, '\n', name, url, '\n', game, status)
    html += '<div class="channel">';
    html += '<img class="logo" src="' + logo + '" alt="logo"/>';
    html += '<span class="name">';
    html += '<a href="' + url + '" target="_blank">' + name + '</a>';
    html += '</span>';
    html += '<span>' + game + ' ' + status + '</span>';
    html += '</div>';
    document.getElementById('channels').innerHTML = html;
  }

  function getDateset2 (data) {
    // console.log('2 --> ', data)
    let name = data.display_name;
    let game = ''; // data.game
    let logo = data.logo ||
      'https://dummyimage.com/50x50/ecf0e7/5c5457.jpg&text=0x3F';
    let url = data.url;
    let status = 'OFFLINE'; // data.status
    if (data.status === 422) {
      name = data.display_name;
      status = 'ACCOUNT CLOSED';
      url = '#';
    }
    html += '<div class="channel">';
    html += '<img class="logo" src="' + logo + '" alt="logo"/>';
    html += '<span class="name">';
    html += '<a href="' + url + '" target="_blank">' + name + '</a>';
    html += '</span>';
    html += '<span>' + game + ' ' + status + '</span>';
    html += '</div>';
    document.getElementById('channels').innerHTML = html;
  }

  return {
    init: init,
    printChannel: printChannel,
    getDateset: getDateset,
    getDateset2: getDateset2
  };
}());

window.addEventListener('load', app.init);
