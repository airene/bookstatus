DROP TABLE IF EXISTS `tb_book`;
CREATE TABLE `tb_book` (
  `id` mediumint(8) unsigned NOT NULL auto_increment,
  `bookcode` varchar(20) NOT NULL default '' COMMENT '图书编码',
  `bookname` varchar(100) NOT NULL default '' COMMENT '图书名称',
  `borrowby` varchar(12) NOT NULL default '0' COMMENT '借阅人',
  `borrowtime` datetime NOT NULL default '0000-00-00 00:00:00' COMMENT '借阅时间',
  `status` tinyint(3) NOT NULL default '0' COMMENT '是否借出',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_borrowlog`;
CREATE TABLE `tb_borrowlog` (
  `id` 			mediumint(8) unsigned 	NOT NULL auto_increment,
  `bookid` 		mediumint(8)  	     	NOT NULL default '0' COMMENT '图书id',
  `borrowby` 	varchar(12) 			NOT NULL default '' COMMENT '借阅人',
  `borrowtime` 	datetime NOT NULL default '0000-00-00 00:00:00' COMMENT '借阅时间',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` 			mediumint(8) unsigned 	NOT NULL auto_increment,
  `username` 	varchar(40)  	   	 	NOT NULL default '' COMMENT '用户名',
  `password` 	varchar(40) 			  NOT NULL default '' COMMENT '登录密码',
  `userrole` 	varchar(40)  	   	 	NOT NULL default '' COMMENT '用户角色',
  `regtime`  datetime NOT NULL default '0000-00-00 00:00:00' COMMENT '注册时间',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


