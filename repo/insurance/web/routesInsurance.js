//SPDX-License-Identifier: Apache-2.0

var tuna = require('./controllerInsurance.js');
console.log(tuna);
module.exports = function(app){

  app.get('/get_person/:id', function(req, res){
    console.log(">>>> router get_person 1");
    tuna.get_person(req, res);
  });
  app.get('/get_claim/:id', function(req, res){
    console.log(">>>> router get_claim 1");
    tuna.get_claim(req, res);
  });
  app.get('/get_claim/:id', function(req, res){
    console.log(">>>> router get_claim 1");
    tuna.get_claim(req, res);
  });
  app.get('/create_claim/:holder', function(req, res){
    console.log(">> create_claim ...1 ");
    tuna.create_claim(req, res);
  });
  app.get('/approve_claim/:holder', function(req, res){
    console.log(">> approve_claim ...1 ");
    tuna.approve_claim(req, res);
  });
  app.get('/reject_claim/:holder', function(req, res){
    console.log(">> reject_claim ...1 ");
    tuna.reject_claim(req, res);
  });
  app.get('/add_policy', function(req, res){
    tuna.add_policy(req, res);
  });
  app.get('/add_claim', function(req, res){
    tuna.add_claim(req, res);
  });
  app.get('/approve_claim', function(req, res){
    tuna.approve_claim(req, res);
  });
  
  app.get('/list_person', function(req, res){
    tuna.list_person(req, res);
  });
  app.get('/change_holder/:holder', function(req, res){
    tuna.change_holder(req, res);
  });
}
