const frisby = require('frisby');
const Joi = frisby.Joi;

const host = 'http://localhost:3000'
const hostApi = host + '/api/v1'

describe('User', function(){

  it('list user should be ok but without password column', function() {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.get(hostApi +'/users')
      .expect('status', 200)
      .expect('jsonTypesStrict', 'data.*', {
        user_id: Joi.string(),
        user_name: Joi.string(),
        created_at: Joi.string(),
        updated_at: Joi.string(),
      })
  })
  it('create user should be not ok without admin login', function() {
    return frisby.post(hostApi +'/users', {
      user_name: 'test',
      password: 'test'
    })
      .expect('status', 401)
  })
  it('admin user should get token', function() {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.post(host +'/login', {
      username: 'admin',
      password: 'admin'

    })
      .expect('status', 200)
      .expect('jsonTypesStrict', {
        code: Joi.number(),
        expire: Joi.string(),
        token: Joi.string().min(20).max(180)
      })
  })

  it('admin user should create user successfully at first time', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.post(host +'/login', {
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
        var resToken = res.json.token
        //console.log('resToken', resToken)
        return frisby.post(hostApi +'/users', {
          headers: {
            Authorization: "Bearer " + resToken,
          },
          body : {
            user_name: 'test',
            password: 'test',
          }
        })
          .expect('status', 201)
      })
      .done(done)
  })
  it('admin user should create existing user failed at second time', function(done) {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.post(host +'/login', {
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
        var resToken = res.json.token
        //console.log('resToken', resToken)
        return frisby.post(hostApi +'/users', {
          headers: {
            Authorization: "Bearer " + resToken,
          },
          body : {
            user_name: 'test',
            password: 'test',
          }
        })
          .expect('status', 500)
          .expect('bodyContains', 'User name exist:')
      })
      .done(done)
  })
})
