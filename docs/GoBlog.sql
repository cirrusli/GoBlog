/*
MySQL Backup
Database: summer_project
Backup Time: 2022-10-14 11:20:12
*/

SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `summer_project`.`articles`;
DROP TABLE IF EXISTS `summer_project`.`comments`;
DROP TABLE IF EXISTS `summer_project`.`likes`;
DROP TABLE IF EXISTS `summer_project`.`replies`;
DROP TABLE IF EXISTS `summer_project`.`users`;
CREATE TABLE `articles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `aid` bigint(20) DEFAULT NULL,
  `author` longtext,
  `uid` bigint(20) DEFAULT NULL,
  `title` longtext,
  `cover` varchar(191) DEFAULT 'https://img-blog.csdnimg.cn/img_convert/7ad86d3945ddf7946e39d179e60c5687.png',
  `summary` longtext NOT NULL,
  `content` longtext,
  `hits` bigint(20) DEFAULT NULL,
  `comments` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_articles_aid` (`aid`),
  KEY `idx_articles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;
CREATE TABLE `comments` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `cid` bigint(20) DEFAULT NULL,
  `comment` longtext,
  `uid` bigint(20) DEFAULT NULL,
  `aid` bigint(20) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `content` longtext,
  `user_id` bigint(20) DEFAULT NULL,
  `user_name` longtext,
  `article_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_comments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
CREATE TABLE `likes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `type_id` bigint(20) DEFAULT NULL,
  `type` tinyint(1) DEFAULT NULL,
  `uid` bigint(20) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `like_type` tinyint(1) DEFAULT NULL,
  `to_uid` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_likes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
CREATE TABLE `replies` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `rid` bigint(20) DEFAULT NULL,
  `reply_type` tinyint(1) DEFAULT NULL,
  `content` longtext,
  `uid` bigint(20) DEFAULT NULL,
  `reply_cid` bigint(20) DEFAULT NULL,
  `to_uid` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_replies_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_name` longtext,
  `password` longtext,
  `avatar_url` varchar(191) DEFAULT 'https://image-1302243118.cos.ap-beijing.myqcloud.com/public/img/BLUE_GLASSES_GOPHER-1616727843503.png',
  `uid` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_m_users_deleted_at` (`deleted_at`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 STATS_PERSISTENT=1;
BEGIN;
LOCK TABLES `summer_project`.`articles` WRITE;
DELETE FROM `summer_project`.`articles`;
INSERT INTO `summer_project`.`articles` (`id`,`created_at`,`updated_at`,`deleted_at`,`aid`,`author`,`uid`,`title`,`cover`,`summary`,`content`,`hits`,`comments`,`likes`,`is_deleted`,`user_id`) VALUES (1, '2022-08-18 21:57:10.000', '2022-08-18 21:59:27.000', NULL, 123123, 'admin', 666666, 'MySQL事务隔离级别', 'https://img-blog.csdnimg.cn/img_convert/7ad86d3945ddf7946e39d179e60c5687.png', '持续创作，加速成长！这是我参与「掘金日新计划 · 6 月更文挑战」的第13天，点击查看活动详情', '持续创作，加速成长！这是我参与「掘金日新计划 · 6 月更文挑战」的第13天，点击查看活动详情\r\n\r\nMySQL 事务隔离级别是为了解决并发事务互相干扰的问题的，MySQL 事务隔离级别总共有以下 4 种：\r\n\r\nREAD UNCOMMITTED：读未提交。\r\nREAD COMMITTED：读已提交。\r\nREPEATABLE READ：可重复读。\r\nSERIALIZABLE：序列化。\r\n1.四种事务隔离级别\r\n1.1 READ UNCOMMITTED\r\n读未提交，也叫未提交读，该隔离级别的事务可以看到其他事务中未提交的数据。该隔离级别因为可以读取到其他事务中未提交的数据，而未提交的数据可能会发生回滚，因此我们把该级别读取到的数据称之为脏数据，把这个问题称之为脏读。\r\n\r\n1.2 READ COMMITTED\r\n读已提交，也叫提交读，该隔离级别的事务能读取到已经提交事务的数据，因此它不会有脏读问题。但由于在事务的执行中可以读取到其他事务提交的结果，所以在不同时间的相同 SQL 查询中，可能会得到不同的结果，这种现象叫做不可重复读。\r\n\r\n1.3 REPEATABLE READ\r\n可重复读，MySQL 默认的事务隔离级别。可重复读可以解决“不可重复读”的问题，但还存在幻读的问题。所谓的幻读指的是，在同一事务的不同时间使用相同 SQL 查询时，会产生不同的结果。例如，一个 SELECT 被执行了两次，但是第二次返回了第一次没有返回的一行，那么这一行就是一个“幻像”行。\r\n\r\n注意：幻读和不可重复读的侧重点是不同的，不可重复读侧重于数据修改，两次读取到的同一行数据不一样；而幻读侧重于添加或删除，两次查询返回的数据行数不同。\r\n\r\n1.4 SERIALIZABLE\r\n序列化，事务最高隔离级别，它会强制事务排序，使之不会发生冲突，从而解决了脏读、不可重复读和幻读问题，但因为执行效率低，所以真正使用的场景并不多。', 651, 32, 49, 0, NULL),(2, '2022-08-19 11:50:50.284', '2022-08-19 11:50:50.284', NULL, 2537860315094265856, 'admin', 666666, 'Docker详细命令大全', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '前面文章讲解了Docker的安装以及基本命令的使用，该文章将介绍Docker镜像命令，容器命令，容器数据卷，Dockerfile的构建，发布镜像，Docker网络等知识', '<h2>Docker的常用命令</h2>\n<hr>\n<h3>帮助命令</h3>\n<hr>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker version #显示docker的版本信息\ndocker info #显示docker的系统命令\ndocker 命令 --help #万能\n</code></pre>\n<p>中文文档：[http://c.biancheng.net/docker/(http://c.biancheng.net/docke)</p>\n<hr>\n<h3>镜像命令</h3>\n<hr>\n<h4>docker images #查看所有镜像</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">[root@192 wfeng]# docker images\nREPOSITORY   TAG       IMAGE ID   CREATED   SIZE\n# 解释\nREPOSITORY 镜像的仓库源\nTAG		   镜像的标签\nIMAGE ID   镜像的id\nCERATED    镜像创建的时间\nSIZE	   镜像的大小\n\n#可选项\n-a, --all	# 列出所有镜像\n-q, --quiet # 只显示镜像的id\n</code></pre>\n<h4>docker search # 镜像搜索</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">[root@192 wfeng]# docker search mysql\nNAME                              DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED\nmysql                             MySQL is a widely used, open-source relation…   11610     [OK]\nmariadb                           MariaDB Server is a high performing open sou…   4417      [OK]\nmysql/mysql-server                Optimized MySQL Server Docker images. Create…   857                  [OK]\n\n# 可选项\n-f,--filter # -f=STARS=3000 搜索镜像STARS&gt;3000的\n</code></pre>\n<h4>docker pull 下载镜像</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\"># 下载镜像 docker pull 镜像名[:tag]\n[root@192 wfeng]# docker pull hello-world\nUsing default tag: latest # 如果不写tag,默认就是latest\nlatest: Pulling from library/hello-world\n93288797bd35: Pull complete  #分层下载，docker image的核心 联合文件系统\n93287127bd35: Pull complete\nDigest: sha256:37a0b92b08d4919615c3ee023f7ddb012b8387475d64c622ac30f45c29c51\nStatus: Downloaded newer image for hello-world:latest\ndocker.io/library/hello-world:latest #真实地址\n\n等价于它\ndocker pull hello-world\ndocker pull docker.io/library/hello-world:latest\n</code></pre>\n<h4>docker rmi 删除镜像</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">[root@192 wfeng]# docker rmi -f ecac195d15af   #指定id删除\n[root@192 wfeng]# docker rmi -f id1 id2        #多个删除\n[root@192 wfeng]# docker rmi -f $(docker images -aq)    #全部删除\n</code></pre>\n<hr>\n<h3>容器命令</h3>\n<hr>\n<p>说明：有了镜像才可以创建容器</p>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker pull centos\n</code></pre>\n<h4>新建容器并启动</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker run [可选参数] image\n\n#参数说明\n--name=\"name\"  容器名字，用来区分容器\n-d             后台方式运行\n-it            使用交互方式运行，进入容器查看内容\n-p			   指定容器的端口 -p 8080:8080\n	-p ip:主机端口:容器端口\n	-p 主机端口:容器端口 （常用）\n	-p 容器端口\n	容器端口\n-p			   随机指定端口\n\n# 测试，启动并进入容器\n[root@192 wfeng]# docker run -it centos /bin/bash\n[root@0ec9062b3f77 /]# \n# 退出\n[root@0ec9062b3f77 /]# exit\n</code></pre>\n<h4>列出所有的运行的容器</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\"># docekr ps 命令\n		# 列出当前正在运行的容器\n-a  	# 列出当前正在运行的容器+带出历史运行过的容器\n-n=?	# 显示最近的创建的容器\n-q		# 只显示容器的编号\n</code></pre>\n<h4>退出容器</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">exit  # 直接退出容器并停止\nctrl + p + q  #不停止退出\n</code></pre>\n<h4>删除容器</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker rm 容器id # 删除指定容器，不能删除运行中的容器，参数 -f 强制删除\ndocker rm -f $(docker ps -aq) # 删除所有容器 \ndocker ps -a -q | xargs docker rm # 删除所有容器\n</code></pre>\n<h4>启动和停止容器的操作</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker start 容器id   #启动\ndocker restart 容器id #重启\ndocker stop 容器id    #停止\ndocker kill 容器id    #强制停止\n</code></pre>\n<h3>常用其他命令</h3>\n<hr>\n<h4>后台启动容器</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">[root@192 wfeng]# docker run -d centos\n# 问题docker ps,发现 centos 停止了\n\n# 常见的坑：docker 容器使用后台运行，就必须要有一个前台进程，docker发现没有应用，就会自动停止\n# nginx,容器启动后，发现自己没有提供服务，就会立刻停止，就是没有程序了\n</code></pre>\n<h4>查看日志</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker logs -tf --tail 10 容器id\n-tf 			# 显示日志\n--tail number   # 显示日志条数\n</code></pre>\n<h4>查看容器中进程信息ps</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker top 容器id\n</code></pre>\n<h4>查看容器元数据</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker inspect 容器id\n</code></pre>\n<h4>进入当前正在运行的容器</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\"># 命令\ndocker exec -it 容器id bashShell # 进入容器打开一个新的终端\ndocker attach 容器id # 进入正在运行的一个终端\n</code></pre>\n<h4>从容器内拷贝文件到主机上</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker cp 容器id:容器内路径  目的的主机路径\n\n[root@57721483c31b home]# touch test.java\n[root@57721483c31b home]# exit\nexit\n[root@192 wfeng]# docker ps\nCONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES\n[root@192 wfeng]# docker cp 57721483c31b:/home/test.java test.java\n[root@192 wfeng]# ls\ntest.java\n</code></pre>\n<h4>小结</h4>\n<p><img src=\"https://img-blog.csdn.net/2018071915491757?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE2MjkwNzkx/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70\" alt=\"img\"></p>\n<h3>容器数据卷</h3>\n<h4>挂载目录</h4>\n<p>实现主机与容器文件共享</p>\n<pre class=\"lang-shell\"><code data-language=\"shell\"># 指定目录\ndocker run -v 主机目录:容器目录 镜像\n\n# 匿名挂载\ndocker run -v /容器目录 镜像\n\n# 具名挂载\ndocker run -v 卷名:/容器目录 镜像\n\n[root@192 wfeng]# docker run -d -p 3310:3306 -v /home/edk/mysql/conf:/etc/mysql/conf.d -v /home/edk/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 --name mysql5.7 ibex/debian-mysql-server-5.7\n\n\n#### 未指定目录的情况下都是在：/var/lib/docker/volume/xxx/_data  大多数使用具名挂载\n</code></pre>\n<p>拓展</p>\n<pre class=\"lang-shell\"><code data-language=\"shell\"># ro  readonly 只读，只能通过主机来操作，容器内部无法操作\n# rw  read/write  可读可写\ndocker run -d -P --name nginx02 -v nginx02:/etc/nginx:ro nginx\ndocker run -d -P --name nginx02 -v nginx02:/etc/nginx:rw nginx\n</code></pre>\n<h3>数据卷容器</h3>\n<p>--volume-from 容器</p>\n<h3>Dockerfile</h3>\n<h4>Dockerfile介绍</h4>\n<p>dockerfile 是用来构建docker镜像的文件！命令参数脚本！<br>\n步骤：</p>\n<ul>\n<li>编写一个 dockerfile 文件</li>\n<li>docker build 构建成为一个镜像</li>\n<li>docker run 运行镜像</li>\n<li>docker push 发布镜像（DockerHub,阿里仓库）</li>\n</ul>\n<hr>\n<h4>Dockerfile构建过程</h4>\n<p><em><strong>基础知识</strong></em></p>\n<ul>\n<li>命令必须是大写字母</li>\n<li>执行顺序从上到下</li>\n<li>#表示注释</li>\n<li>每个指令都会创建提交一个新的镜像层，并提交！</li>\n</ul>\n<h4>Dockerfile 指令</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">FROM            # 基础镜像\nMAINTAINER      # 创建者\nRUN             # 镜像构建的时候需要运行的命令\nADD             # 添加镜像内容,tar类型文件会自动解压(网络压缩资源不会被解压)，可以访问网络资源，类似wget\nWORKDIR         # 镜像的工作目录,类似于cd命令\nVOLUME          # 挂载的目录\nEXPOSE          # 对外暴露的端口\nCMD             # 指定容器启动的时候运行的命令，只有最后一个会成效，可被替代\nENTRYOPOINT     # 指定容器启动的时候运行的命令，可以追加\nONBUILD         # 当构建一个被继承 DockerFile 这个时候就会运行 ONBUILD 的指令，触发指令\nCOPY            # 类似ADD,将我们的文件拷贝到镜像中,但是是不会自动解压文件，也不能访问网络资源\nENV             # 构建的时候设置环境变量\n</code></pre>\n<h4>DockerFile 实战</h4>\n<p>Docker Hub 中99%都是从这个基础镜像过来的 FROM scratch,然后配置需要的软件和配置来进行构建</p>\n<blockquote>\n<p>测试创建一个centos镜像</p>\n</blockquote>\n<pre class=\"lang-shell\"><code data-language=\"shell\"># 1.编写dockerfile文件\nFROM centos\nMAINTAINER wfeng&lt;1224971566@qq.com&gt;\n\nENV MYPATH /usr/local\nWORKDIR $MYPATH\n\n\nRUN yum -y install vim\nRUM yum -y install net-tools\n\nEXPOSE 80\n\nCMD echo $MYPATH\nCMD echo \"---end---\"\nCMD /bin/bash\n\n# 2.通过文件构建镜像\n# 命令： docker build -f 文件 -t 镜像名:版本 .\n</code></pre>\n<h4>发布镜像</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\"># 1.登录 dockerhub\n[root@192 dockerfile]# docker login -u fengziy\nPassword:\nWARNING! Your password will be stored unencrypted in /root/.docker/config.json.\nConfigure a credential helper to remove this warning. See\nhttps://docs.docker.com/engine/reference/commandline/login/#credentials-store\n\nLogin Succeeded\n\n# 2.push hub\n# 规范镜像名才能push  docker push 用户名/镜像名[:tag]\n[root@192 dockerfile]# docker push fengziy/mycentos:0.1\nThe push refers to repository [docker.io/library/mycentos]\n591b3cb2ac87: Preparing\nc4651d2a50d0: Preparing\nd871dadfb37b: Preparing\n</code></pre>\n<h5>发布阿里云镜像仓库</h5>\n<ol>\n<li>登录阿里云Docker Registry</li>\n</ol>\n<pre class=\"lang-shell\"><code data-language=\"shell\">[root@192 wfeng]# docker login --username=182*****507 registry.cn-shenzhen.aliyuncs.com\nPassword:\nWARNING! Your password will be stored unencrypted in /root/.docker/config.json.\nConfigure a credential helper to remove this warning. See\nhttps://docs.docker.com/engine/reference/commandline/login/#credentials-store\n\nLogin Succeeded\n</code></pre>\n<ol start=\"2\">\n<li>pull/push</li>\n</ol>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker pull registry.cn-shenzhen.aliyuncs.com/fengziy/镜像:[镜像版本号]\ndocker push registry.cn-shenzhen.aliyuncs.com/fengziy/镜像:[镜像版本号]\n</code></pre>\n<hr>\n<h3>Docker网络（容器编排 集群）</h3>\n<p>#veth-pair技术 一对虚拟设备接口<br>\n容器之间通信通过 docker0 网桥转发</p>\n<h4>--link 连接容器,通过容器名</h4>\n<p>本质：修改了 hosts 映射</p>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker run -d -P --name tomcat01 tomcat\ndocker run -d -P --name tomcat02 --link tomcat01 tomcat\ndocker exec -it tomcat02 ping tomcat01\n</code></pre>\n<h4>自定义网络</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">#网卡列表\n[root@192 wfeng]# docker network ls\nNETWORK ID     NAME      DRIVER    SCOPE\n5b37dfde19ba   bridge    bridge    local\n76afc652c9c4   host      host      local\n1c29823208de   none      null      local\n\n#网络模式\n#bridge：桥接 docker 默认\n#none: 不配置\n#host：和主机共享\n\n# 默认网络参数是：--net bridge\n[root@192 wfeng]# docker run -d -P --name tomcat01 --net bridge tomcat\n\n\n#创建网络\n[root@192 wfeng]# docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet\n800c94b5ff39de6874ce89dc12a39b3c66dd0aca328138dfd671f348501f5485\n[root@192 wfeng]# docker run -d -it --name centos02 --net mynet centos\n379a4ff7d36c8e79ad6f83a3aadb1036dfab5f4fc2855af88d52e7e349ed071f\n[root@192 wfeng]# docker run -d -it --name centos01 --net mynet centos\nc65f745c64509541c193c2f3d2676a6d68438e6e1fd28c801bd1eddf64e2e0d4\n\n#测试ping容器名称\n[root@192 wfeng]# docker exec -it centos01 ping centos02\nPING centos02 (192.168.0.2) 56(84) bytes of data.\n64 bytes from centos02.mynet (192.168.0.2): icmp_seq=1 ttl=64 time=0.656 ms\n</code></pre>\n<h4>容器连接一个网络</h4>\n<pre class=\"lang-shell\"><code data-language=\"shell\">docker network connet 网络 容器\n</code></pre>', 0, 0, 0, 0, NULL),(3, '2022-08-19 18:18:08.700', '2022-08-19 18:55:48.294', '2022-08-19 18:59:10.327', 2537957784075448320, 'admin', 666666, 'last try', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '前面文章讲解了Docker的安装以及基本命令的使用，该文章将介绍Docker镜像命令，容器命令，容器数据卷，Dockerfile的构建，发布镜像，Docker网络等知识', 'test update the article content and status', 0, 0, 0, 1, NULL),(4, '2022-08-19 18:20:53.361', '2022-08-19 18:20:53.361', NULL, 2537958474717933568, 'admin', 666666, 'Docker详细命令大全', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '前面文章讲解了Docker的安装以及基本命令的使用，该文章将介绍Docker镜像命令，容器命令，容器数据卷，Dockerfile的构建，发布镜像，Docker网络等知识', 'test', 0, 0, 0, 1, NULL),(6, '2022-08-19 18:29:06.406', '2022-08-19 18:29:06.406', NULL, 2537957784075448322, 'admin', 666666, 'Docker详细命令大全', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '前面文章讲解了Docker的安装以及基本命令的使用，该文章将介绍Docker镜像命令，容器命令，容器数据卷，Dockerfile的构建，发布镜像，Docker网络等知识', 'test', 0, 0, 0, 0, NULL),(7, '2022-08-19 18:29:34.982', '2022-08-19 18:29:34.982', NULL, 234234, 'admin', 666666, 'Docker详细命令大全', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '前面文章讲解了Docker的安装以及基本命令的使用，该文章将介绍Docker镜像命令，容器命令，容器数据卷，Dockerfile的构建，发布镜像，Docker网络等知识', 'test', 0, 0, 0, 0, NULL),(8, '2022-08-19 18:30:26.799', '2022-08-19 18:30:26.799', '2022-08-19 19:14:26.067', 345345, 'admin', 666666, 'Docker详细命令大全', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '前面文章讲解了Docker的安装以及基本命令的使用，该文章将介绍Docker镜像命令，容器命令，容器数据卷，Dockerfile的构建，发布镜像，Docker网络等知识', 'test update the article content and status', 0, 0, 0, 1, NULL),(9, '2022-08-19 18:33:36.115', '2022-08-19 18:33:36.115', NULL, 456456, 'test', 123123, 'Docker详细命令大全', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '前面文章讲解了Docker的安装以及基本命令的使用，该文章将介绍Docker镜像命令，容器命令，容器数据卷，Dockerfile的构建，发布镜像，Docker网络等知识', 'test update the article content and status', 0, 0, 0, 1, NULL),(10, '2022-08-26 19:52:24.610', '2022-08-26 19:52:24.610', NULL, 2540518221740584960, 'chengxu', 2540423949712437248, 'Vue.js介绍', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架。', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架。与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合。另一方面，当与现代化的工具链以及各种支持类库结合使用时，Vue 也完全能够为复杂的单页应用提供驱动。如果你想在深入学习 Vue 之前对它有更多了解，我们制作了一个视频，带您了解其核心概念和一个示例工程。如果你已经是有经验的前端开发者，想知道 Vue 与其它库/框架有哪些区别，请查看对比其它框架。', 0, 0, 1, NULL, NULL),(11, '2022-08-26 19:57:09.032', '2022-08-26 19:57:09.032', NULL, 2540519414697111552, 'chengxu', 2540423949712437248, 'Vue.js介绍——第二部分', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架。', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架。与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合。另一方面，当与现代化的工具链以及各种支持类库结合使用时，Vue 也完全能够为复杂的单页应用提供驱动。如果你想在深入学习 Vue 之前对它有更多了解，我们制作了一个视频，带您了解其核心概念和一个示例工程。如果你已经是有经验的前端开发者，想知道 Vue 与其它库/框架有哪些区别，请查看对比其它框架。', 0, 0, 0, NULL, NULL),(12, '2022-08-30 08:35:50.704', '2022-08-30 08:35:50.704', NULL, 5222127940625349, 'chengxu', 2540423949712437248, 'Vue.js介绍——第三部分', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架。与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合。另一方面，当与现代化的工具链以及各种支持类库结合使用时，Vue 也完全能够为复杂的单页应用提供驱动。如果你想在深入学习 Vue 之前对它有更多了解，我们制作了一个视频，带您了解其核心概念和一个示例工程。如果你已经是有经验的前端开发者，想知道 Vue 与其它库/框架有哪些区别，请查看对比其它框架。', 0, 0, 0, NULL, NULL),(13, '2022-08-30 08:36:40.973', '2022-08-30 08:36:40.973', NULL, 5222131235054533, 'chengxu', 2540423949712437248, 'Vue.js介绍——第四部分', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '想知道 Vue 与其它库/框架有哪些区别，请查看对比其它框架。', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架。与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合。另一方面，当与现代化的工具链以及各种支持类库结合使用时，Vue 也完全能够为复杂的单页应用提供驱动。如果你想在深入学习 Vue 之前对它有更多了解，我们制作了一个视频，带您了解其核心概念和一个示例工程。如果你已经是有经验的前端开发者，想知道 Vue 与其它库/框架有哪些区别，请查看对比其它框架。', 0, 0, 0, NULL, NULL),(14, '2022-09-16 16:12:04.918', '2022-09-16 16:12:04.918', NULL, 5320181213914053, 'chengxu', 2540423949712437248, 'Vue.js介绍——第四部分', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', '想知道 Vue 与其它库/框架有哪些区别，请查看对比其它框架。', 'Vue (读音 /vjuː/，类似于 view) 是一套用于构建用户界面的渐进式框架。与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合。另一方面，当与现代化的工具链以及各种支持类库结合使用时，Vue 也完全能够为复杂的单页应用提供驱动。如果你想在深入学习 Vue 之前对它有更多了解，我们制作了一个视频，带您了解其核心概念和一个示例工程。如果你已经是有经验的前端开发者，想知道 Vue 与其它库/框架有哪些区别，请查看对比其它框架。', 0, 0, 0, NULL, NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `summer_project`.`comments` WRITE;
DELETE FROM `summer_project`.`comments`;
INSERT INTO `summer_project`.`comments` (`id`,`created_at`,`updated_at`,`deleted_at`,`cid`,`comment`,`uid`,`aid`,`likes`,`is_deleted`,`content`,`user_id`,`user_name`,`article_id`) VALUES (1, '2022-08-25 16:24:49.079', '2022-08-25 16:24:49.079', NULL, 2540103591591489536, NULL, 666666, 123123, 0, NULL, 'test post comment', NULL, NULL, NULL),(2, '2022-08-26 13:42:24.906', '2022-08-26 13:42:24.906', '2022-08-27 18:23:21.010', 2540425109437493248, NULL, 2540423949712437248, 2540424791253397500, 0, NULL, 'test data back', NULL, NULL, NULL),(3, '2022-08-27 09:50:22.743', '2022-08-27 09:50:22.743', NULL, 2540729103514939392, NULL, 2540423949712437248, 2540424791253397500, 0, NULL, '数据一致性', NULL, NULL, NULL),(4, '2022-08-27 09:53:29.364', '2022-08-27 09:53:29.364', NULL, 2540729886260146176, NULL, 666666, 2540424791253397500, 2, NULL, '必须先保存才可以吗？', NULL, NULL, NULL),(5, '2022-09-04 17:36:12.348', '2022-09-04 17:36:12.348', '2022-09-04 17:36:23.039', 5252564277486533, NULL, 666666, 2540424791253397500, 0, NULL, '测试删除评论', NULL, NULL, NULL),(6, '2022-09-30 14:48:50.195', '2022-09-30 14:48:50.195', NULL, 5399126225347525, NULL, 666666, 2540424791253397500, 0, NULL, '测试删除评论', NULL, NULL, NULL),(7, '2022-09-30 14:49:45.751', '2022-09-30 14:49:45.751', NULL, 5399129866265541, NULL, 666666, 2540424791253397500, 0, NULL, '测试1', NULL, NULL, NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `summer_project`.`likes` WRITE;
DELETE FROM `summer_project`.`likes`;
INSERT INTO `summer_project`.`likes` (`id`,`created_at`,`updated_at`,`deleted_at`,`type_id`,`type`,`uid`,`status`,`like_type`,`to_uid`) VALUES (1, '2022-09-09 10:39:50.267', '2022-09-09 10:39:50.267', NULL, 2540518221740584960, NULL, 666666, 1, 0, 2540423949712437248),(2, '2022-09-09 10:40:24.180', '2022-09-09 10:40:24.180', NULL, 2540518221740584960, NULL, 666666, 1, 0, 2540423949712437248),(3, '2022-09-09 10:41:26.576', '2022-09-09 10:41:26.576', NULL, 2540518221740584960, NULL, 666666, 1, 0, 2540423949712437248),(4, '2022-09-09 11:17:16.254', '2022-09-09 11:17:16.254', NULL, 2540518221740584960, NULL, 666666, 0, 0, 2540423949712437248),(5, '2022-09-09 11:17:39.572', '2022-09-09 11:17:39.572', NULL, 2540518221740584960, NULL, 666666, 1, 0, 2540423949712437248),(6, '2022-09-09 11:18:18.879', '2022-09-09 11:18:18.879', NULL, 2540729886260146176, NULL, 666666, 1, 1, 2540423949712437248),(7, '2022-09-09 11:18:54.916', '2022-09-09 11:18:54.916', NULL, 2540729886260146176, NULL, 2540533250644848640, 1, 1, 2540423949712437248);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `summer_project`.`replies` WRITE;
DELETE FROM `summer_project`.`replies`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `summer_project`.`users` WRITE;
DELETE FROM `summer_project`.`users`;
INSERT INTO `summer_project`.`users` (`id`,`created_at`,`updated_at`,`deleted_at`,`user_name`,`password`,`avatar_url`,`uid`) VALUES (1, '2022-08-18 14:45:29.735', '2022-08-18 16:15:38.142', NULL, 'admin', 'dc4325d7130abf0ce0d272c53db39d619b40287fdc58df7ae048c4d7c2a2464c', 'https://image-1302243118.cos.ap-beijing.myqcloud.com/public/img/BLUE_GLASSES_GOPHER-1616727843503.png', 666666),(2, '2022-08-18 16:38:41.164', '2022-08-18 16:38:41.164', NULL, 'testUid', '6d2d9fc610337f813a1b85869ec214129940860543ad04308d87357f6c0133f6', 'https://image-1302243118.cos.ap-beijing.myqcloud.com/public/img/BLUE_GLASSES_GOPHER-1616727843503.png', 2537570366549733376),(3, '2022-08-26 13:37:48.407', '2022-08-26 13:37:48.407', NULL, 'chengxu', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'https://image-1302243118.cos.ap-beijing.myqcloud.com/public/img/BLUE_GLASSES_GOPHER-1616727843503.png', 2540423949712437248),(4, '2022-08-26 20:00:05.107', '2022-08-26 20:00:05.107', NULL, 'lixuanyu', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'https://image-1302243118.cos.ap-beijing.myqcloud.com/public/img/BLUE_GLASSES_GOPHER-1616727843503.png', 2540520153204994048),(5, '2022-08-26 20:52:07.780', '2022-08-26 20:52:07.780', NULL, 'lzq', 'dc4325d7130abf0ce0d272c53db39d619b40287fdc58df7ae048c4d7c2a2464c', 'https://image-1302243118.cos.ap-beijing.myqcloud.com/public/img/BLUE_GLASSES_GOPHER-1616727843503.png', 2540533250644848640);
UNLOCK TABLES;
COMMIT;
