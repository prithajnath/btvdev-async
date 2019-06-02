const axios = require("axios");
const API_URL = "https://0bdwnj5rj7.execute-api.us-east-1.amazonaws.com/prod"

async function callFactorAPI(num, data){
  console.log(`Factoring ${num}`);
  var result = await axios({
     method: 'get',
     url: API_URL,
     data: data
  });
  console.log(`${num} can be factored into ${JSON.parse(result.data.body).result} \n`);
  //console.log(result.data)
}

(async function(){
 var N = parseInt(process.argv.slice(2)[0]);
 var numsToFactor = Object.keys([...Array(N)]).map(x => Math.floor(Math.random() * 5000) + 5000);
 //console.log("numsToFactor : ", numsToFactor)
 try{
 await Promise.all(numsToFactor.map(num => callFactorAPI(num, {number:num})));
 }catch(e){
  console.log(e);
 }
 //var factor = await callFactorAPI(N, {number:N});

})();
