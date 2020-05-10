CREATE TABLE `sys_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(100) DEFAULT NULL COMMENT '密码',
  `salt` varchar(20) DEFAULT NULL COMMENT '盐',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `mobile` varchar(100) DEFAULT NULL COMMENT '手机号',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态  0：禁用   1：正常',
  `dept_id` int(11) DEFAULT NULL COMMENT '部门ID',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=107 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='系统用户'


INSERT INTO `sys_user` VALUES ('1', 'administrator', '5f9c50b9d370e553b076ecf20870baab6dff1d061fb15868b62ca17f04b70a16', 'YzcmCZNvbXocrsz9dm8e', 'administrator@hq.com', 'admin', '1', '1', '2016-11-11 11:11:11', 'https://hqjy-coursetest.oss-cn-shenzhen.aliyuncs.com/knowledge_base/image/portrait/20190117/b769083fb1a64073a59fa0fe3cc033c4.jpeg');
INSERT INTO `sys_user` VALUES ('2', 'zengqi1', 'b9438621eb4d1e51200dd3c1ba99e2009b0db2c934637eaf012dad120c8a1050', '', '', '17620712093', '0', '32', '2018-11-22 17:09:16', '');
INSERT INTO `sys_user` VALUES ('3', '462877@qq.ocm', '5f9c50b9d370e553b076ecf20870baab6dff1d061fb15868b62ca17f04b70a16', 'YzcmCZNvbXocrsz9dm8e', 'admin@123.com', '18148770353', '1', '1', '2016-11-11 11:11:11', 'https://hqjy-coursetest.oss-cn-shenzhen.aliyuncs.com/image/md/20190102/073a99864ab5459c80469306fbc61af2.jpg');
INSERT INTO `sys_user` VALUES ('5', '秦天柱', '', '', '', '', '1', '0', '2018-11-29 15:22:58', '');
INSERT INTO `sys_user` VALUES ('8', '忘川河', 'admin', '', 'admin', '1111111', '1', '0', '2018-11-29 15:23:08', '');
INSERT INTO `sys_user` VALUES ('62', 'zengqi', 'd57082ad5a37f6bfdba3628ce2424073bb03fb57f688a110b7fddebd38ba862f', 'nTVds41IXyjsYSYDJgc8', '446', '15011133345', '1', '34', '2018-09-19 17:29:45', 'https://hqjy-coursetest.oss-cn-shenzhen.aliyuncs.com/image/md/20181229/3941e5c9119d424eab8eb6aad796b4a0.jpg');
INSERT INTO `sys_user` VALUES ('64', '测试', '2382aa24bd21419dd6b2e6f32a220de2ab2f0fa3c781d93b0ee3ba7515e5c593', 'ozESb0U7VANsT1bi8Hlx', '4444@qq.com', '15622308421', '0', '1', '2018-09-19 18:00:23', '');
INSERT INTO `sys_user` VALUES ('81', '陆挺', 'be5f10dae970c4dfda8f43ab8834a4d81c242597dc16fd459f35451aae30799b', 'zDZPzmrPK2cJS3KdVXU1', 'yaya@123,com', '18148770321', '1', '29', '2018-12-29 15:46:24', 'https://hqjy-coursetest.oss-cn-shenzhen.aliyuncs.com/knowledge_base/image/portrait/20190107/3e2cfd43981f4cb1b7b38a494aea7f36.jpg');
INSERT INTO `sys_user` VALUES ('88', '陈小小', '1f04ab269a64457413af789fc142b3526987175c1ae35108f01ee0ad176280a3', 'C91keVcTDLmqX4F3cNNz', '', '15622758019', '1', '32', '2019-05-09 17:36:14', '');
INSERT INTO `sys_user` VALUES ('89', 'ABC', '5ef63a85cd7f97f46a3e034b9dcd2d58032efbf466b71fea7e76e394590439dc', '4qbATzkEWKX87yLmwZgU', '', '15013349143', '0', '28', '2019-05-10 14:39:12', '');
INSERT INTO `sys_user` VALUES ('92', '陈达达', '123456', 'DVcyqxoSP4M5oDqsjtYf', '110@136.com', '15622758011', '1', '1', '2019-05-10 15:17:06', '');
INSERT INTO `sys_user` VALUES ('93', 'SDEWWEWE', '0bd8a2601ea73261e36ff1bc84bbf0e2b736fcb7636944345a79129845b9351a', 'iUe60erqZ0HekDzFVEQx', '', '15013349147', '1', '34', '2019-05-16 09:13:49', '');
INSERT INTO `sys_user` VALUES ('95', '91黄先生', '23456', 'H9KtgCxAUefcIhuFi3gp', '12306@qq.com', '18776038733', '1', '1', '2019-12-21 22:39:40', '');
INSERT INTO `sys_user` VALUES ('96', 'fufulong', '5f9c50b9d370e553b076ecf20870baab6dff1d061fb15868b62ca17f04b70a16', 'YzcmCZNvbXocrsz9dm8e', '1303038078', '18819153276', '1', '1', '2019-09-29 15:34:06', 'shenzhen.aliyuncs.com/knowledge_base/image/portrait/20190117/b769083fb1a64073a59fa0fe3cc033c4.jpeg');
INSERT INTO `sys_user` VALUES ('97', 'testliu', 'a0075921a2a02616ebdb3b9a9f7edbf64a58aea26f34dd539bc6102cb3034588', 'oWj3gYBObwiAQhIeeHqX', '', '17688391209', '0', '1', '2019-11-11 14:12:16', '');
INSERT INTO `sys_user` VALUES ('98', 'aaa', 'admin', '', 'admin', 'aaa', '0', '0', '2019-12-21 11:52:05', '');
INSERT INTO `sys_user` VALUES ('99', '10092', 'admin', '', 'admin', '11', '0', '0', '2019-12-21 11:52:46', '');
INSERT INTO `sys_user` VALUES ('100', 'qwqw', '123456', '', 'qwqw', 'wqwq', '0', '0', '2019-12-21 12:10:27', '');
INSERT INTO `sys_user` VALUES ('101', '23', 'admin', '', 'admin', '2323', '0', '0', '2019-12-21 12:13:09', '');
INSERT INTO `sys_user` VALUES ('102', '23232', 'admin', '', 'admin', '32323', '0', '0', '2019-12-21 12:14:24', '');
INSERT INTO `sys_user` VALUES ('103', 'asdf', 'admin', '', 'admin', 'aasdfasd', '0', '0', '2019-12-21 12:16:10', '');
INSERT INTO `sys_user` VALUES ('104', 'sfg', 'admin', '', 'admin', 'fgsdfgs', '0', '0', '2019-12-21 12:16:57', 'https://hqjy-coursetest.oss-cn-shenzhen.aliyuncs.com/knowledge_base/image/portrait/20190107/3e2cfd43981f4cb1b7b38a494aea7f36.jpg');
INSERT INTO `sys_user` VALUES ('105', '欧安', 'admin', '', 'admin', '110', '1', '0', '2019-12-21 21:17:49', 'https://hqjy-coursetest.oss-cn-shenzhen.aliyuncs.com/knowledge_base/image/portrait/20190107/3e2cfd43981f4cb1b7b38a494aea7f36.jpg');
INSERT INTO `sys_user` VALUES ('106', '小二', 'admin', '', '15600@qq.com', '15600', '1', '94', '2019-12-21 22:14:00', 'https://hqjy-coursetest.oss-cn-shenzhen.aliyuncs.com/knowledge_base/image/portrait/20190107/3e2cfd43981f4cb1b7b38a494aea7f36.jpg');