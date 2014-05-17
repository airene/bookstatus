CREATE TABLE `tb_book` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `bookcode` varchar(20) NOT NULL DEFAULT '' COMMENT '图书编码',
  `bookname` varchar(100) NOT NULL DEFAULT '' COMMENT '图书名称',
  `booktype` varchar(45) NOT NULL DEFAULT '' COMMENT '图书类型',
  `borrowby` varchar(12) NOT NULL DEFAULT '0' COMMENT '借阅人',
  `borrowtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '借阅时间',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否借出',
  `borrowcount` mediumint(8) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
CREATE TABLE `tb_borrowlog` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `bookid` mediumint(8) NOT NULL DEFAULT '0' COMMENT '图书id',
  `borrowby` varchar(12) NOT NULL DEFAULT '' COMMENT '借阅人',
  `borrowtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '借阅时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
CREATE TABLE `tb_user` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(40) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(40) NOT NULL DEFAULT '' COMMENT '登录密码',
  `userrole` varchar(40) NOT NULL DEFAULT '' COMMENT '用户角色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
