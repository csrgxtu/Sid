### Sid
图形化内容传播路径，可以用来展示微博，twitter，病毒，公司内部创建内容的传播路径。

|  Port |  dev |  prod |
|--:|---|---|
|  9013 |  dev-sid.beautifulreading.com |  sid.beautifulreading.com |

### How To Start
```bash
# First, clone the code
git clone https://github.com/beautifulreading/Sid.git

# Second, change the run mode in conf/app.conf
runmode = dev

# Third, start the server
bee run
```

### Result
一篇文章的传播情况，这里只列出上层的传播情况，因为深度太深，无法在一个页面完整列出来[demo](http://localhost:9013/static/d3.html?vid=77b75ab7527556b208ba55ee127df8ac)。
![result.jpg](https://github.com/csrgxtu/Sid/blob/master/static/result.jpg)
