app.controller('AssignmentCtrl', function($scope, $routeParams, $location , AssgnSrv){
    $scope.assignment = {}
    $scope.assignment.name = ""
    $scope.assignment.uniqueId = ""
    $scope.save = function(assignment){
        AssgnSrv.create({name : assignment.name , unique_assignment : assignment.uniqueId})
            .success(function(data){
            $rootScope.$broadcast("successful:Creation of assignment: "+assignment.name);
            $location.path('/');
        })
        .error(function(){
            alert("error");
            $location.path('/');
        })  
};
