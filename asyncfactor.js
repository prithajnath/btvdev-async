const axios = require("axios");
const API_URL = "https://0bdwnj5rj7.execute-api.us-east-1.amazonaws.com/prod"

async function callFactorAPI(num, data){
  console.log(`Factoring ${num}`);
  var result = await axios({
     method: 'get',
     url: API_URL,
     data: data
  });
  return JSON.parse(result.data.body).result
}

(async function(){
 var N = parseInt(process.argv.slice(2)[0]);
 var numsToFactor = Object.keys([...Array(N)]).map(x => Math.floor(Math.random() * 5000) + 5000);
 try{
 var factoredValues = await Promise.all(numsToFactor.map(num => callFactorAPI(num, {number:num})));
 }catch(e){
  console.log(e);
 }
 Object.keys(numsToFactor).forEach(n => {
   console.log(`${numsToFactor[n]} can be factored into ${factoredValues[n]}`);
 });
})();
