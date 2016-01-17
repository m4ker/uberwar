# uberwar

2016 Uber Hackathon Project

## 简介

uberwar是一款基于Uber行程和真实世界地理位置的塔防游戏, 当uber用户开始了一次新的行程, 就可以进入游戏加入战斗. 

![Uber War Screenshot](http://tripwar.laoyou.mobi/html/screenshot.png)

第一次进入游戏, 用户会被分配到一个阵营, 与另一个阵营的玩家进行对抗, 在行程中部署防御塔来攻击敌方玩家来获得分数.

## 语言和开发环境

我们将项目分成了三个部分:

1. 服务端, 由golang实现, 负责调用uber api并且给客户端/前端提供数据
2. IOS客户端
3. web版客户端, html+js

我们用到了uber以下一些接口:

1. oauth 用户授权
2. profile 获取用户信息
3. request/current 获取用户当前的行程信息和位置信息
4. request 用来模拟用户行程

## 相关链接

* 演示地址 : [http://tripwar.laoyou.mobi/html/](http://tripwar.laoyou.mobi/html/)
* Github : [https://github.com/m4ker/uberwar](https://github.com/m4ker/uberwar)
* 设计文档 : [https://github.com/m4ker/uberwar/wiki/%E6%8E%A5%E5%8F%A3%E8%AE%BE%E8%AE%A1](https://github.com/m4ker/uberwar/wiki/%E6%8E%A5%E5%8F%A3%E8%AE%BE%E8%AE%A1)
* 接口文档 : [https://github.com/m4ker/uberwar/wiki/%E6%8E%A5%E5%8F%A3%E5%AE%9A%E4%B9%89](https://github.com/m4ker/uberwar/wiki/%E6%8E%A5%E5%8F%A3%E5%AE%9A%E4%B9%89)

## 演讲稿

各位同学大家好, 我们给大家带来了一款游戏, 它起名字叫 优步战争, 优步战争是一款基于Uber行程和真实世界地理位置的塔防游戏, 当uber用户开始了一次新的行程, 就可以进入优步战争的世界加入战斗. 

游戏分成两个阵营, 用户第一次进入游戏, 会被分配到其中一个阵营中, 与另一个阵营的玩家进行对抗, 在行程中部署防御塔来攻击敌方玩家来获得分数. 

每个玩家会有自己的血量和能量, 血量决定你是否能够参与行动, 能量用来建造防御塔和攻击敌方的防御塔, 玩家每开始一个新的行程, 血量和能量会进行初始化, 随着你的行程时间的增长, 你的血量和能量会得到补充. 如果你的身边有其他队友, 还会得到血量和能量的加成. 

接下来我们来进行一个简单的演示:

(点亮手机, 切换屏幕)

由于IOS版本遇到一些技术问题没有办法进行演示, 所以我们用web版本进行演示.

(点击桌面图标)

首先我们看到的是游戏的启动页面, 我们假设用户目前已经在行程中了, 所以我们点击"开始战斗"进行授权并进入游戏页面.

(点击 开始战斗)

可以看到目前我们正沿着苏州街向北行驶, 画面中的蓝色点标代表友军的防御塔, 而红色点标代表敌方防御塔, 我们可以建造自己的防御塔

(点击 建造)

(点击 目标)

建造成功会得一分, 减掉一些能量值, 并且看到一条消息. 我们也可以攻击敌军的防御塔:

(点击 攻击)

(点击 目标)

攻击成功敌军的防御塔会消失, 我方得一分, 减掉一些能量. 

当我们到达了敌方防御塔的攻击范围, 会被攻击, 攻击会减少血量, 敌军会得一分. 

由于时间比较紧迫, 所以我们只设计了少量的游戏元素, 并且展示效果也不是非常好. 如果在这两方面可以再花些时间这个游戏一定会更有趣. 

接下来介绍一下我们的项目: 

1. 我们将源代码托管在github上, 并且用github的wiki进行协作
2. IOS客户端
3. web版客户端, html+js
4. 服务端, 由golang实现, 负责调用uber api并且给客户端/前端提供数据, 我们服务端的同学还顺便开发了uber的go语言sdk.
