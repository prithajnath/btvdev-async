import asyncio
import aiohttp
import json
import sys
from random import randint

API_URL = "https://0bdwnj5rj7.execute-api.us-east-1.amazonaws.com/prod"


async def call_factor_API(session,num, data):
    print(f"Factoring {num}...")
    async with session.get(url=API_URL, json=data) as resp:
        result = await resp.text()
        all_factors = json.loads(json.loads(result)['body'])['result']

        all_factors = map(str, all_factors)
        print(f"{num} can be factored into : {', '.join(all_factors)}")

async def main():
    #N = int(input("How many numbers to factor?"))
    N = int(sys.argv[1])
    numbers = [randint(5000, 10000) for _ in range(N)]
    tasks = []
    async with aiohttp.ClientSession() as session:
        for n in numbers:
            tasks.append(asyncio.ensure_future(call_factor_API(session,n,{"number":n})))

        await asyncio.gather(*tasks)

loop = asyncio.get_event_loop()
loop.run_until_complete(main())
loop.close()
