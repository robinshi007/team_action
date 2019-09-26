const frisby = require('frisby');
const Joi = frisby.Joi;

const host = 'http://localhost:3000'
const hostApi = host + '/api/v1'

describe('User', function(){
  var token = ""
  var tetUserId=""

  it('list user should be ok but without password column', function() {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.get(hostApi +'/users')
      .expect('status', 200)
      .expect('jsonTypesStrict', 'data.*', {
        user_id: Joi.string(),
        username: Joi.string(),
        created_at: Joi.string(),
        updated_at: Joi.string(),
        is_active: Joi.number().integer().default(0),
        last_login_at: Joi.string(),
      })
  })
  it('create user should be not ok without admin login', function() {
    return frisby.post(hostApi +'/users', {
      username: 'test',
      password: 'test'
    })
      .expect('status', 401)
  })
  it('admin user should get token', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.post(hostApi +'/login', {
      username: 'admin',
      password: 'admin'

    })
      .expect('status', 200)
      .expect('jsonTypesStrict', {
        code: Joi.number(),
        expire: Joi.string(),
        token: Joi.string().min(20).max(180)
      })
      .then(function(res){
        token = res.json.token
      })
      .done(done)
  })

  it('admin user should create user successfully at first time', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.post(hostApi +'/users', {
      headers: {
        Authorization: "Bearer " + token,
      },
      body : {
        username: 'test',
        password: 'test',
      }
    })
      .expect('status', 201)
      .expect('jsonTypesStrict', {
        data: Joi.string()
      })
      .then(function(res){
        testUserId = res.json.data
      })
      .done(done)
  })
  it('admin user should create existing user failed at second time', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.post(hostApi +'/users', {
      headers: {
        Authorization: "Bearer " + token,
      },
      body : {
        username: 'test',
        password: 'test',
      }
    })
      .expect('status', 500)
      .expect('bodyContains', 'User name exist:')
      .done(done)
  })
  it('admin user should able to get the test user', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.get(hostApi +'/users/' + testUserId, {
      headers: {
        Authorization: "Bearer " + token,
      }
    })
      .expect('status', 200)
      .expect('bodyContains', 'test')
      .done(done)
  })
  it('admin user should able to update the test user', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.put(hostApi +'/users/' + testUserId, {
      headers: {
        Authorization: "Bearer " + token,
      },
      body : {
        username: 'test1',
        password: 'test1',
      }
    })
      .expect('status', 200)
      .then(function(res){
        return frisby.get(hostApi +'/users/' + testUserId, {
          headers: {
            Authorization: "Bearer " + token,
          }
        })
          .expect('status', 200)
          .expect('bodyContains', 'test1')
          .done(done)
      })
  })
  it('admin user should able to update password the test user', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.put(hostApi +'/users/' + testUserId + '/update_password', {
      headers: {
        Authorization: "Bearer " + token,
      },
      body : {
        password: 'test11',
      }
    })
      .expect('status', 200)
      .then(function(res){
        return frisby.post(hostApi +'/login', {
          body : {
            username: 'test1',
            password: 'test11',
          }
        })
          .expect('status', 200)
          .done(done)
      })
  })
  it('admin user should able to update last_login_at the test user', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.put(hostApi +'/users/' + testUserId + '/update_last_login', {
      headers: {
        Authorization: "Bearer " + token,
      },
    })
      .expect('status', 200)
      .then(function(res){
        return frisby.get(hostApi +'/users/' + testUserId, {
          headers: {
            Authorization: "Bearer " + token,
          }
        })
          .expect('status', 200)
          .expect('json','data.last_login_at', new RegExp((new Date()).getFullYear()))
          .done(done)
      })
  })
  it('admin user should able to delete the test user', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.delete(hostApi +'/users/' + testUserId, {
      headers: {
        Authorization: "Bearer " + token,
      }
    })
      .expect('status', 204)
      .done(done)
  })
})
