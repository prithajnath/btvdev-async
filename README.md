# BTV Python - Writing Asynchronous Python

## What these scripts do
Factors random N 4-5 digit integers. The factoring is done through an API call to my Lambda function.
`syncfactor.py` does it synchronously with the `requests` library, and `asyncfactor.py` does it asynchronously with `asyncio` and `aiohttp`. 

Time the synchronous version

```sh
time python3.6 syncfactor.py 500
```
![](https://res.cloudinary.com/dzmp7xptn/image/upload/v1558444938/Screen_Shot_2019-05-21_at_9.16.57_AM_zgcnbp.png)


Time the asynchronous version
```sh
time python3.6 asyncfactor.py 1000
```

![](https://res.cloudinary.com/dzmp7xptn/image/upload/v1558444938/Screen_Shot_2019-05-21_at_9.17.45_AM_hzklvs.png)
