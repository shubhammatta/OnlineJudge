app.factory('UserSrv', function($http){
  var obj = {};

  obj.show_profile = function(){
    return $http.get('/users/profile')
  }

  // obj.show_problems = function(){
  //   return $http.post('/users/problems',data)
  // }
  return obj;
})
