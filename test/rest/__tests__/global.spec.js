const frisby = require('frisby');
const Joi = frisby.Joi;

const host = 'http://localhost:3000'
const hostApi = host + '/api/v1'

it('health should be ok', function(done) {
  // Return the Frisby.js Spec in the 'it()' (just like a promise)
  return frisby.get(hostApi+'/health')
    .expect('status', 200)
    .expect('json', 'status', 'success')
    .expect('jsonTypesStrict', {
      status: Joi.string()
    })
    .done(done)
})
