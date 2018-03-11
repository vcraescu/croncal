# Cron Calendar

![alt text](https://user-images.githubusercontent.com/941660/37251106-d4e43758-2511-11e8-86ff-cf73770a244d.png "Screenshot")

### Usage:

Export your crontab to file:

```
sudo crontab -l >> crontab.txt
```

Then run croncal app passing the exported crontab file from above:

```
./croncal crontab.txt
```

By default, the UI is will bind to `0.0.0.0:8080` interface. 

