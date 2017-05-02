app.factory('Problemsrv', function($http){
  var obj = {};

  obj.create = function(data){
    console.log(data)
    return $http.post('/creates/problem', data)
  }
  return obj;
})
