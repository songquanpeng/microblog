# Micro Blog
> 基于 Go 的单可执行文件「微型博客」实现，一个供你闲言碎语的地方

## 部署教程
### 手动部署
1. 从 [GitHub Releases](https://github.com/songquanpeng/microblog/releases/latest) 下载可执行文件或者从源码编译：
   ```shell
   git clone https://github.com/songquanpeng/microblog.git
   go mod download
   go build -ldflags "-s -w" -o microblog
   ````
2. 运行：
   ```shell
   chmod u+x microblog
   ./microblog --port 3000 --token 123456
   ```
3. 访问 [http://localhost:3000/](http://localhost:3000/) 即可开始使用，默认 token 为 `123456`。

其中 `port` 即服务的端口号，`token` 即用于身份验证的令牌。

可以使用 `pm2` 来进行进程守护：`pm2 start ./microblog --name microblog -- --port 3004 --token my_private_token`

更加详细的部署教程[参见此处](https://iamazing.cn/page/how-to-deploy-a-website)。

### 基于 Docker 进行部署
执行：`docker run -d --restart always -p 3000:3000 -e TOKEN=123456 -e TZ=Asia/Shanghai -v /home/ubuntu/data/microblog:/data justsong/microblog`

数据将会保存在宿主机的 `/home/ubuntu/data/microblog` 目录，鉴权 Token 为 `123456`。

## 使用教程
点击右下角的加号按钮可以发布微博，在输入框中输入 `delete id` 可以删除指定 id 的微博，微博的 id 可以在对应卡片的右下角找到。

例如删除 id 为 34 的微博：`delete 34`。

## 在线示例
我自己的部署版本：https://nonsense.justsong.cn/

## 其他
创意参考：https://github.com/daibor/nonsense.fun