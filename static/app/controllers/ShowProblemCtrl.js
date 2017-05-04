app.controller("ShowProblemCtrl", ['$scope', '$routeParams',
  function($scope, $http ,$routeParams) {
  var getProblems;
  $scope.problems = [];
  getProblems = function() {
    return $http.get('/assignment/'+$routeParams.id).success(function(data) {
      return $scope.problems = data;
    });
  };
  getProblems();
}]);
