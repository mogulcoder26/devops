const express = require('express');
const app = express();

app.get('/',(req,res)=>{
    res.status(200).send(`Hi there,Soubhik!!!Dudee!!jesus`);
})


app.listen(5000 || process.env.PORT,()=>{
    console.log(`5000`);
})

