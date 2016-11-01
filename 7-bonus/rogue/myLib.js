var bb = {
  getRandomNumber: function (min, max) {
    return Math.floor(Math.random() * (max - min + 1) + min);
  },
  initializeMultiArray: function (cols, rows, value) {
    var array = [];
    for (var i = 0; i < cols; i++) {
      array[i] = [];
      for (var j = 0; j < rows; j++) {
        array[i][j] = value;
      }
    }
    return array;
  },
  //
  // bb.getAjaxData(url, returnData)
  // function returnData (dataset) {
  //   console.log(dataset)
  // }
  getAjaxData: function (urlData, callback) {
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      var DONE = 4;
      var OK = 200;
      if (xhr.readyState === DONE) {
        if (xhr.status === OK) {
          callback(JSON.parse(xhr.responseText));
        } else {
          console.log('Error: ' + xhr.status);
        }
      }
    };
    xhr.open('GET', urlData); // add false to synchronous request
    xhr.send();
  }
};
