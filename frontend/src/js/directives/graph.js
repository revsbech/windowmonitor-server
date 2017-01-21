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
