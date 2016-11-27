package main



var Toolbar = `
var app = angular.module('Calc', ['ngMaterial'], function($interpolateProvider) {
    $interpolateProvider.startSymbol('[[');
    $interpolateProvider.endSymbol(']]');
});
var toa;
//var math=JSON.parse('{"app":[{"x":"10","y":"20","o":"+","v":"50"},{"x":"5","y":"3","o":"+","v":"8"}]}'); //外部的全局变量访问不到？

app.controller('AppCtrl', function($scope,$rootScope, $timeout, $mdSidenav, $mdUtil, $log, $mdToast) {
    $scope.toggleLeft = buildToggler('left');
    $scope.toggleRight = buildToggler('right');
    toa = $mdToast;

    function buildToggler(navID) {
      var debounceFn =  $mdUtil.debounce(function(){
            $mdSidenav(navID)
              .toggle()
              .then(function () {
                $log.debug("toggle " + navID + " is done");
              });
          },300);
      return debounceFn;
    }
    
    var conn = new WebSocket("ws://localhost:8888/ws");
    conn.onclose = function(e){
      $mdToast.show(
        $mdToast.simple()
          .content('Disconnected!')
          .hideDelay(3000)
      );      
    };
    conn.onopen = function(e) {
      $mdToast.show(
        $mdToast.simple()
          .content('Connected!')
          .hideDelay(3000)
      );
    };
    conn.onmessage = function(e) {
      $rootScope.fat = JSON.parse(e.data);
      $mdToast.show(
        $mdToast.simple()
          .content($rootScope.fat.app[0].x)
          .hideDelay(3000)
      );
      $rootScope.$apply();
    };  

  });
app.controller('LeftCtrl', function ($scope, $timeout, $mdSidenav, $log) {
    $scope.close = function () {
      $mdSidenav('left').close()
        .then(function () {
          $log.debug("close LEFT is done");
        });
    };
  });
app.controller('RightCtrl', function ($scope, $timeout, $mdSidenav, $log) {
    $scope.close = function () {
      $mdSidenav('right').close()
        .then(function () {
          $log.debug("close RIGHT is done");
        });
    };
  });  


app.controller('tm_ctl',function ($scope) {
  $scope.clock = {
    now: new Date()
  };
  var updateClock = function(){
    $scope.clock.now = new Date();
  };
  setInterval(function () {
    $scope.$apply(updateClock);
  },1000);
  updateClock();  
});

function Get_value($index, $scope, $mdToast){
 toa.show(
    toa.simple()
      .content($scope.fat.app[$index].v)
      .hideDelay(1000)
  );
};


`;