# TG-N5 设备说明


```go
NB1;1234567;1;2;+25.5;00;030;+21;+22
```


电量和信号都是123， 分别代表低中高

断点续传功能，说明之前有因如信号不稳定等因素，造成上传不成功的数据。缓存起来，有上传条件时，一起上传。先进先出原则：即先发送的温度数据为先存储的不成功数据，距离目前时间点较远。每个数据时间间隔为上传周期。上面例子中的就是30分钟