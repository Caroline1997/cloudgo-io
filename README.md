## 处理 web 程序的输入与输出
* **支持静态文件服务**
```
//在服务器上创建目录，以存放静态内容.
assets(静态文件虚拟根目录)
  |-- js
  |-- images
  +-- css
```
访问静态文件截图如下：

![访问 html 静态网页](https://raw.githubusercontent.com/Caroline1997/cloudgo-io/master/screenshot/static.PNG)
* **支持简单 js 访问**

网页使用javascrip获取下面的信息，程序中用于在访问静态文件时页面使用javascript获取信息显示在网页上

![网页使用 javascript 获取信息](https://raw.githubusercontent.com/Caroline1997/cloudgo-io/master/screenshot/apitest.PNG)
* **提交表单，并输出一个表格**

输入登录的信息并提交，跳转到输出表格的页面

![login](https://raw.githubusercontent.com/Caroline1997/cloudgo-io/master/screenshot/login.PNG)

![input](https://raw.githubusercontent.com/Caroline1997/cloudgo-io/master/screenshot/inputdata.PNG)

![jump](https://raw.githubusercontent.com/Caroline1997/cloudgo-io/master/screenshot/table.PNG)
* **对 /unknown 给出开发中的提示，返回码 5xx**

模仿 NotFound 函数，当访问/unknown 时返回页面提示 501 Not Implemented

![1](https://raw.githubusercontent.com/Caroline1997/cloudgo-io/master/screenshot/501(1).PNG)

![2](https://raw.githubusercontent.com/Caroline1997/cloudgo-io/master/screenshot/501(2).PNG)