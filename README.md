## GiteeAutoUpdate

推送到 Gitee 自动更新到其它平台的 Webhook 服务端

```
GAU 1.0.0 - Gitee Auto Update is a bot to do something useful for project on Gitee

Usage: Config your config.json and start GAU with command ./gau -c ./config.json &

Args:
        -c      Specify a config json
        -h      Help and quit
        -v      Show current version and quit

```

#### 功能特点
1. 支持同步推送到多个平台
2. 支持 Webhook IP 白名单过滤

#### 设计流程图

![输入图片说明](https://images.gitee.com/uploads/images/2020/1128/224510_ee7e0c4a_62561.png "屏幕截图.png")

### 快速使用（云服务）

1. 提交 PullRequest 到本仓库的`config/syncWhitelist`文件，按照格式将自己的仓库以及目标仓库加入
2. 在对应平台将`GiteeAutoUpdate`用户加到对应仓库，并赋予写权限（目前仅在 Github 上创建了用户）
3. 添加`http://139.162.89.15:6666/sync`到 Gitee 仓库的 Push Hook

### 自行搭建
你可以选择在自己的服务器上搭建一套`GiteeAutoUpdate`

#### 构建说明

```
git clone https://gitee.com/kesin/GiteeAutoUpdate.git
cd GiteeAutoUpdate
go build
./gau -c config/config.json
```

#### 配置说明
```
{
  "Port": 6666,  // 启动端口
  "GiteePrivateToken": "xxxxxxxxxxx", // Gitee 的私人令牌，用于创建Issue和评论，同步结果
  "IpWhitelist": [
    "127.0.0.1" // 白名单过滤
  ],
  "UpdateWhitelistToken": "xxxxxx", // 更新仓库白名单的 Token
  "SyncUser": { // 要推送的平台的用户
    "github.com": {
      "username": "GiteeAutoUpdate",
      "password": "xxxxxx"
    },
    "gitlab.com": {
      "username": "GiteeAutoUpdate",
      "password": "xxxxxxx"
    }
  }
}

```

#### TODO
1. 将推送结果以 Issue 的形式提交到对应仓库

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

欢迎提交更多其他 Hooks 功能的支持