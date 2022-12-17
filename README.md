# 微博客
> 基于 Go 的个人微博客，一个供你闲言碎语的地方

## 部署教程
### 基于 Docker 进行部署
执行：`docker run -d --restart always -p 3000:3000 -e MB_USERNAME=admin -e MB_PASSWORD=123456 -e TZ=Asia/Shanghai -v /home/ubuntu/data/microblog:/data justsong/microblog`

设置 `SESSION_SECRET` 环境变量后将固定 Session Secret，这样应用重启后 Cookie 也不会失效。

数据将会保存在宿主机的 `/home/ubuntu/data/microblog` 目录，默认用户名为 `admin`，密码为 `123456`。

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
   ./microblog --port 3000 --username admin --password 123456
   ```
3. 访问 [http://localhost:3000/](http://localhost:3000/) 即可开始使用，默认用户名为 `admin`，密码为 `123456`。

其中 `port` 即服务的端口号，`username` 和 `password` 即用于身份验证的用户名和密码。

可以使用 `pm2` 来进行进程守护：`pm2 start ./microblog --name microblog -- --port 3000 --username admin --password 123456`

更加详细的部署教程[参见此处](https://iamazing.cn/page/how-to-deploy-a-website)。

## 主题设置
你可以通过设置 `theme` 命令行参数或者 `MB_THEME` 环境变量来切换主题。

目前可用的内置主题有 `default`。

如果要使用第三方主题，只需要设置 `theme` 命令行参数或者 `MB_THEME` 环境变量为主题的路径即可。

例如：`./microblog --theme ./path/to/theme`。

如果想要自行开发主题，请参考 [theme/default](./theme/default)，你可以自行复制一份，然后修改 `index.html` 以及 `app.css` 即可。

欢迎提交 PR 将你的自定义主题 merge 到本项目。

## 使用教程
点击右下角的加号按钮可以发布微博，在输入框中输入 `delete id` 可以删除指定 id 的微博，微博的 id 可以在对应卡片的右下角找到。

例如删除 id 为 34 的微博：`delete 34`。

如果用户未登录，则会首先弹出登录框，用户名和密码是你在环境变量或者命令行参数中设置的值。

支持 Markdown。

由于微博客本身简洁的设计，你可以将其使用 `iframe` 嵌入到你的主博客里面，例如：https://iamazing.cn/page/nonsense

## 在线示例
我自己的部署版本：https://nonsense.justsong.cn

## 其他
1. 创意参考：https://github.com/daibor/nonsense.fun
2. `v0.1` 版本升级 `v0.2` 版本请使用此[数据库迁移脚本](./bin/migration-v0.1-v0.2.py)。