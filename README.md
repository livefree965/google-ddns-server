# Google Domain DDNS 转发服务

本项目使用 Golang 实现了对Google Domain DDNS接口的转发，用于解决在国内无法直接访问 Google Domain DDNS 接口的问题。通过这个项目，您可以在国内使用 Google Domain DDNS 服务，实现动态域名解析。

## 项目背景

Google Domain 是 Google 提供的域名注册和管理服务，其中包括了动态域名解析（Dynamic DNS，简称 DDNS）功能，允许用户通过 API 更新域名的 IP 地址，实现动态域名的解析。然而，由于某些原因，Google Domain 在国内的访问受到限制，导致无法直接使用其 DDNS 接口。

为了解决这个问题，本项目使用 Golang 编写了一个简单的转发器，将国内访问转发到 Google Domain DDNS 接口，实现了在国内也能够正常使用 Google Domain DDNS 服务的功能。

## 功能特点

- 将国内无法直接访问的 Google Domain DDNS 接口转发到可访问的地址，解决了无法使用 Google Domain DDNS 服务的问题。
- 使用 Golang 编写，具有高效性和跨平台特性，可以在多种操作系统上运行。
- 代码开源，您可以根据自己的需求进行修改和定制。

## 使用方法

1. 使用前需要在domains.google.com中配置好动态DNS，**注意这里和固定DNS是分开配置的**，解析ip可以随便填，配置完成后会生成随机账号和密码，并在接下来的接口调用时填入。

2. 直接在命令行中执行，如有需要调整端口号即可：
   
   ```bash
   docker run -d -p 8080:8080 livefree965/google_ddns_server
   ```

3. GET请求调用接口，也可以直接在浏览器中请求：
   
   ```
   http://your_ip:your_port/dns/update?hostname=**&ip=**&username=**&password=**
   ```

## 贡献

欢迎对本项目进行贡献！您可以通过以下方式参与项目：

- 提交问题和建议，帮助改进项目。
- 提交代码修复或新增功能，帮助增强项目。
- 分享项目给其他人，帮助更多人解决 Google Domain DDNS 访问受限的问题。

## 许可

本项目使用 MIT 许可证，详细信息请参阅LICENSE文件。

## 帮助和支持

如有任何问题或需要帮助，请在 GitHub 项目页面提交问题，我们会尽快回复并解答您的疑问。

## 免责声明

本项目仅用于解决 Google Domain DDNS 访问受限的问题，不用于违反 Google Domain 或其他相关服务的条款和条件。请合法使用本项目，一切后果由用户自负。