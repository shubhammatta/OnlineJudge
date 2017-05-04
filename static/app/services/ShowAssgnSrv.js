app.factory('ShowAssgnSrv', function($http){
  var obj = {};

  obj.show = function(data){
    console.log(data)
    return $http.get('/assignment/' + data)
  }
  return obj;
})
