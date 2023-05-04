CREATE DATABASE IF NOT EXISTS go_im_service DEFAULT CHARACTER SET utf8;
use go_im_service;


CREATE TABLE `im_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `appid` bigint(20) DEFAULT NULL COMMENT '应用id',
  `master` bigint(20) DEFAULT NULL COMMENT '群主',
  `super` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是超级群 1=是',
  `name` varchar(255) DEFAULT NULL COMMENT '群名称',
  `notice` varchar(255) DEFAULT NULL COMMENT '公告',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除标志',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `im_group_member` (
  `group_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '群id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `timestamp` int(11) DEFAULT NULL COMMENT '入群时间,单位：秒',
  `nickname` varchar(255) DEFAULT NULL COMMENT '群内昵称',
  `mute` tinyint(1) DEFAULT '0' COMMENT '群内禁言',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除标志',
  PRIMARY KEY (`group_id`,`uid`),
  KEY `idx_group_member_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `im_friend` (
  `appid` bigint(20) NOT NULL DEFAULT '0' COMMENT '应用id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `friend_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '好友id',
  `timestamp` int(11) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`appid`, `uid`,`friend_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='好友关系';


CREATE TABLE `im_blacklist` (
  `appid` bigint(20) NOT NULL DEFAULT '0' COMMENT '应用id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `friend_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '好友id',
  `timestamp` int(11) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`appid`, `uid`,`friend_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='黑名单';


