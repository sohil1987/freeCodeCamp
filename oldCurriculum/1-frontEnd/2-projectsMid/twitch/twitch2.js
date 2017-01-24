var channels = ['OgamingSC2', 'freecodecamp', 'storbeck', 'terakilobyte',
  'habathcx', 'RobotCaleb', 'thomasballinger', 'noobs2ninjas', 'beohoff',
  'brunofin', 'comster404', 'test_channel', 'cretetion', 'sheevergaming',
  'TR7K', 'ESL_SC2'
];

var type;

function init () {
  btnAll.addEventListener('click', listChannels);
  btnOnline.addEventListener('click', listChannels);
  btnOffline.addEventListener('click', listChannels);
}

function listChannels (ev) {
  document.getElementById('results').innerHTML = '';
  type = ev.target.id;
  // console.log('CLICK --> ', type)
  for (var i = 0; i < channels.length; i++) {
    var script = document.createElement('script');
    var cb = '?callback=printChannel';
    var url = 'https://api.twitch.tv/kraken/streams/' + channels[i] + cb;
    script.src = url;
    // document.body.appendChild(script)
    document.head.appendChild(script);
  }
}

function printChannel (data) {
  if (data.stream !== null && data.stream !== undefined) {
    if (type === 'btnOnline' || type === 'btnAll') {
      getHtml(data);
    }
  } else if (data.stream === null) {
    if (type === 'btnOffline' || type === 'btnAll') {
      var channel = data._links.channel;
      var script = document.createElement('script');
      var cb = '?callback=getHtml2';
      var url = channel + cb;
      script.src = url;
      // document.body.appendChild(script)
      document.head.appendChild(script);
    }
  } else if (data.stream === undefined) {
    if (type === 'btnOffline' || type === 'btnAll') {
      data.display_name = data.message;
      getHtml2(data);
    }
  }
}

function getHtml2 (data) {
  var name = data.display_name;
  var game = ''; // data.game
  var logo = data.logo ||
    'https://dummyimage.com/50x50/ecf0e7/5c5457.jpg&text=0x3F';
  var url = data.url;
  var status = 'OFFLINE'; // data.status
  if (data.status === 422) {
    name = data.display_name;
    status = 'ACCOUNT CLOSED';
    url = '#';
  }

  var html = document.getElementById('results').innerHTML;
  html += '<div class="row text-xs-center no">';
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
  html += '<div class="space"></div>';
  document.getElementById('results').innerHTML = html;
}

function getHtml (data) {
  var name = data.stream.channel.display_name;
  var game = data.stream.game;
  var logo = data.stream.channel.logo ||
    'https://dummyimage.com/50x50/ecf0e7/5c5457.jpg&text=0x3F';
  var url = data.stream.channel.url;
  var status = data.stream.channel.status;
  var html = document.getElementById('results').innerHTML;
  html += '<div class="row text-xs-center yes">';
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
  html += '<div class="space"></div>';
  document.getElementById('results').innerHTML = html;
}

addEventListener('load', init);
