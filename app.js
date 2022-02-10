const express = require('express');
const res = require('express/lib/response');

const app = express();

app.use((req,res,next) => {
    var serverName = process.argv[2];
    res.status(200).json({
        "Message" :serverName + " is Running"
    });
});

module.exports = app;