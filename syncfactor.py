from random import randint
import json
import requests as r

N = int(input("How many numbers to factor?: "))
numbers = [randint(5000,10000) for _ in range(N)]
API_URL = "https://0bdwnj5rj7.execute-api.us-east-1.amazonaws.com/prod"

for n in numbers:
    print(f"Factoring {n}...")
    resp = json.loads(r.get(url=API_URL, json={"number":n}).text)
    all_factors = json.loads(resp['body'])['result']
    print(f"{n} can be factored into : {all_factors}")
    
