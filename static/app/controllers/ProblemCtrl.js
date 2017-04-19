app.controller("ProblemCtrl", function($scope, $rootScope, $location , Problemsrv){
  	$scope.problem = {}
  	$scope.problem.statement = ""
  	$scope.problem.test = ""
  	$scope.uniqueId = ""
  	$scope.save = function(problem){
  		Problemsrv.create({statement : problem.statement , test : problem.test , unique_assignment : problem.uniqueId })
  	.success(function(data){
  		$rootScope.$broadcast("success");
  		$location.path('/create/problem');
  	})
  	.error(function(){
  		alert("error-");
  		$location.path('/create/problem');
  	})
  };
});