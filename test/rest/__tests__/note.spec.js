const frisby = require('frisby');
const Joi = frisby.Joi;

const host = 'http://localhost:3000'
const hostApi = host + '/api/v1/noteapp'

describe('Category', function(){
  var categoryId=""
  var noteId=""
  it('user should able to create category', function() {
    return frisby.post(hostApi +'/categories', {
      body : {
        name: 'category01',
      }
    })
      .expect('status', 201)
      .expect('jsonTypesStrict', {
        data: Joi.string(),
      }).then(function(res){
        categoryId = res.json.data
      })
  })
  it('user should be list all categories', function() {
    // Return the Frisby.js Spec in the 'it()' (just like a promise)
    return frisby.get(hostApi +'/categories')
      .expect('status', 200)
      .expect('jsonTypesStrict', 'data.*', {
        id: Joi.string(),
        name: Joi.string(),
        created_at: Joi.string(),
        updated_at: Joi.string(),
      })
  })
  it('user should able to create first note with category_id', function() {
    return frisby.post(hostApi +'/notes', {
      body : {
        title: 'note01',
        body: 'note01',
        category_id: categoryId,
      }
    })
      .expect('status', 201)
      .expect('jsonTypesStrict', {
        data: Joi.string(),
      }).then(function(res){
        noteId = res.json.data
      })
  })
  it('user should able to create second note with category_id', function() {
    return frisby.post(hostApi +'/notes', {
      body:{
        title: 'note02',
        body: 'note02',
        category_id: categoryId,
      }
    })
      .expect('status', 201)
  })
  it('user should able to list 2 notes with category id', function() {
    return frisby.get(hostApi +'/categories/' +  categoryId)
      .expect('status', 200)
      .expect('jsonTypesStrict', 'data.notes.*', {
        id: Joi.string(),
        title: Joi.string(),
        body: Joi.string(),
        created_at: Joi.string(),
        updated_at: Joi.string(),
      })
      .then(function(res){
        return expect(res.json.data.notes).toHaveLength(2)
      })
  })
  it('user should able to update note with id', function(done) {
    return frisby.put(hostApi +'/notes/' +  noteId, {
      body: {
        title: 'note03',
        body: 'note03',
        category_id: categoryId,
      }
    })
      .expect('status', 200)
      .then(function(){
        return frisby.get(hostApi +'/notes/' +  noteId)
        .expect('status', 200)
        .expect('bodyContains', 'note03')
        .done(done)
      })
  })
  it('user should able to delete note with id', function() {
    return frisby.delete(hostApi +'/notes/' +  noteId)
      .expect('status', 204)
  })
  it('user should able to update category with id', function(done) {
    return frisby.put(hostApi +'/categories/' +  categoryId, {
      body: {
        name: 'category02',
      }
    })
      .expect('status', 200)
      .then(function(){
        return frisby.get(hostApi +'/categories/' +  categoryId)
        .expect('status', 200)
        .expect('bodyContains', 'category02')
        .done(done)
      })
  })
  it('user should able to delete category with id', function(done) {
    return frisby.delete(hostApi +'/categories/' +  categoryId)
      .expect('status', 204)
      .then(function(res){
        return frisby.get(hostApi +'/categories')
          .expect('status', 200)
          .then(function(res){
            return expect(res.json.data).toHaveLength(0)
          })
        .done(done)
      })
  })
  it('user should able to delete category with id caused CASCADE notes delete', function(done) {
    return frisby.get(hostApi +'/notes')
      .expect('status', 200)
      .then(function(res){
        return expect(res.json.data).toHaveLength(1)
      })
      .done(done)
  })
})

