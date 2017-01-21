'use strict';

// Declare app level module which depends on views, and components
var NATApp = angular.module('NATApp', [
  'ui.router',
  'ui.bootstrap',
  'NATApp.version'
]);

NATApp.config(function ($stateProvider, $urlRouterProvider) {
  // UI Routing
  $urlRouterProvider.otherwise("/view1"); // If url not found redirect to dashboard

  $stateProvider
    // Independent Login Template
    .state('view1', {
      url: '/view1',
      templateUrl: 'views/view1.html',
      controller: 'View1Controller',
    })
    .state('view2', {
      url: '/view1',
      templateUrl: 'views/view2.html',
      controller: 'View2Controller',
    })

});

'use strict';

angular.module('NATApp.version.interpolate-filter', [])

.filter('interpolate', ['version', function(version) {
  return function(text) {
    return String(text).replace(/\%VERSION\%/mg, version);
  };
}]);

'use strict';

angular.module('NATApp.version.version-directive', [])

.directive('appVersion', ['version', function(version) {
  return function(scope, elm, attrs) {
    elm.text(version);
  };
}]);

'use strict';

angular.module('NATApp.version', [
  'NATApp.version.interpolate-filter',
  'NATApp.version.version-directive'
])

.value('version', '0.1');

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
'use strict';

NATApp.controller("View2Controller", function($scope) {

});

/**
 * This directive expects the date, equipment and frequency to be in scope!
 *
 * - Device must be a devicemodel with an id attribute
 * - Equipment must be an equipment model with a code attribute
 * - Date must be and object containting minDate and maxDate as moment objects
 * - frequency must be a string "minute", "hour" or "day"
 *
 */

NATApp.directive('graph', function ($window) {
  return {
    restrict: 'E',
    link: function (scope, elem, attrs) {

      var chart = null;

      var options = {
          xaxis: {
              mode: 'time',
              timezone: 'browser'
          },
          yaxes: [
              { position: 'left', labelWidth: 50},
          ],
          tooltip: true,
          grid: {
              hoverable: true,
              clickable: true,
              autoHighlight: true,
              borderColor: '#404041'
          }
      };
      scope.$watchCollection("timeseries", function (timeseries) {
        //console.log("Redrawing...");

        var dataset = [{
          'label': 'Data points',
          'data':  timeseries
        }];

        if (chart) {
          chart.setData(dataset);
          chart.setupGrid();
          chart.draw();
          return;
        }

        chart = $.plot(elem, dataset, options);
        elem.show();

      }, true);

      // Cleanup the chart when it goes out of scope
      // https://groups.google.com/forum/#!topic/angular/5rNNI8ONhYQ
      scope.$on('$destroy', function () {
        if (chart) {
          chart.shutdown();
          //delete chart;
        }
      });
    }
  }
});


function AuthenticationError(message) {
    this.message = (message || "");
}
AuthenticationError.prototype = new Error();

NATApp.factory('APIservice', function ($q, $http) {
  var deviceHiveBaseUrl = '/';

  /*
  var deviceHiveApi = new DeviceHiveApi({
      baseUrl: deviceHiveBaseUrl,
      request: function(request) {
          var deferred = $q.defer();
          var canceler = $q.defer();
          request.cache = false;
          request.headers['Authorization'] = "Bearer " + deviceHiveToken;
          //http://docs.angularjs.org/api/ng.$http#methods_get
          $http(request)
              .success(function(data, status, headers, config) {
                  if (request.parse) {
                      deferred.resolve(request.parse(data));
                      return;
                  }
                  deferred.resolve(data);
              })
              .error(function(data, status, headers, config) {
                  deferred.reject(status);
                  if (status == 401) {
                      throw new AuthenticationError('Could not query the api');
                  }
              })
          ;

          return deferred.promise;
      }
  });
  */

  var serviceAPI = {
    getSamples: function (deviceId, equipmentId, frequency, startTime, endTime) {
      return deviceHiveApi.getSamples(deviceId, equipmentId, frequency, startTime.format(), endTime.format());
    }
  };

  return serviceAPI;
});

