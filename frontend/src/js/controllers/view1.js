'use strict';

NATApp.controller("View1Controller", function($scope) {
  $scope.timeseries = [];
  /*
  $scope.timeseries = [
      [1477834243000, 24],
      [1477835253000, 45],
      [1478834243000, 50],
      [1479834243000, 56],
      [180834243000, 200]
    ];
/**/

  var ws = new WebSocket("ws://localhost:8080/channel/demochannel/listen");
  ws.onmessage = function(evt) {
    var dataPoint = JSON.parse(evt.data);
    var copy = [];
    copy = angular.copy($scope.timeseries);
    if (copy.length > 2048) {
      copy = copy.slice(1);
    }
    //copy.push([moment(dataPoint.timestamp).toDate(), dataPoint.value])
    copy.push([moment(dataPoint.timestamp).valueOf(), dataPoint.value])
    $scope.timeseries = copy;
    //console.log($scope.timeseries);
    //console.log("Data updated..");
    $scope.$digest();
  };
  /**/
});