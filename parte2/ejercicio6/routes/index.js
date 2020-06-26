var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.send({ 
    name: "Web Service de ejemplo",
    version: "0.0.0"
  });
});

/* GET home page. */
router.get('/date', function(req, res, next) {
  res.send({ 
    date: new Date().toISOString(),
  });
});

module.exports = router;
