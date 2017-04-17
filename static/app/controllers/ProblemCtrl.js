app.controller("ProblemCtrl", function($scope, $routeParams, $location , UserSrv){
  $scope.user = {}
  UserSrv.show_problems()
    .success(function(data){
      $scope.user = data;
    })
    .error(function(){
      $location.path('/');
    })
});