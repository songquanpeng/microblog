# Go Nonsense
> 基于 Go 的单可执行文件「废话胶囊」实现，一个供你发泄情绪的地方

## 部署教程

仅单可执行文件，启动命令：

```bash
./go-nonsense -port 3000 -token 123456  
```

其中 port 即服务的端口号，token 即用于身份验证的令牌。

可以使用 pm2 来进行进程守护：`pm2 start ./go-nonsense --name go-nonsense -- -port 3004 -token my_private_token`

## 使用教程

点击右下角的加号按钮可以发布废话，在输入框中输入 `delete id` 可以删除指定 id 的废话，废话的 id 可以在对应卡片的右下角找到。

例如删除 id 为 34 的废话：`delete 34`。

## 在线示例
我自己的部署版本：https://nonsense.justsong.cn/

Heroku 试用版本（token 就是 token）：https://go-nonsense.herokuapp.com/

## 其他

创意参考：https://github.com/daibor/nonsense.fun