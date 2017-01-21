'use strict';

angular.module('NATApp.version', [
  'NATApp.version.interpolate-filter',
  'NATApp.version.version-directive'
])

.value('version', '0.1');
