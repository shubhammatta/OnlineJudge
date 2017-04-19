app.factory('AssgnSrv', function($http){
  var obj = {};

  obj.create = function(data){
    console.log(data)
    return $http.post('/create/assignment', data)
  }
  return obj;
})
