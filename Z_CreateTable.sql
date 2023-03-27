create database fastwiki;

use fastwiki;

#用户注册表
create table tb_user(
                        `id` bigint primary key auto_increment,
                        `userid`  bigint unsigned  not null ,
                        `author_id` varchar(64)  character set utf8mb4 COLLATE utf8mb4_general_ci not null comment '用户名' ,
                        `password` varchar(64) character set utf8mb4 COLLATE utf8mb4_general_ci not null comment '密码',
                        `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL comment '邮箱',
                        `phonenumber` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL comment '手机号',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                        unique(author_id,phonenumber)
);
#用户权限表
create table tb_userpermissions(
                                   `id` bigint primary key auto_increment,
                                   `userid` bigint unsigned not null ,
                                   `userpermissions` tinyint unsigned not null default 0
);

#默认生成超级管理员用户
insert into tb_user(userid,author_id,password,email,phonenumber) value (47223192342757376,'admin','313233343536ebcf5293e61ff555c5916a650df38d84','admin@key.com','13379797979');
insert into tb_userpermissions(userid, userpermissions) VALUES (47223192342757376,3);

## 用户相关结束=============================================================================================================================================




#类别表
create table `tb_community`(
                               `community_id` bigint primary key auto_increment comment '文本类型代码',
                               `community_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL comment '文本类型',
                               `introduction` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL comment '文本类型描述',
                               `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

insert into tb_community(community_name, introduction) VALUES('战甲','warfarme'),('近战武器','close combat'),
                                                             ('主武器','Rifle'),('手枪','pistol'),
                                                             ('增幅器','Growth of plant'),('专精','specialization'),
                                                             ('mod','模组'),('赋能','plugin');

#详细类别表
create table `tb_SecCommunity`(
                                  `id` bigint primary key auto_increment,
                                  `community_id` bigint comment '文本类型代码',
                                  `sec_community_id` varchar(10)  comment '次级文本类型代码',
                                  `sec_community_info` varchar(128)  comment '次级分类信息',
                                  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)

insert into tb_SecCommunity(community_id, sec_community_id, sec_community_info) VALUES (1,001,'战甲'),
                                                                                       (2,001,'侍刃'),(2,002,'单手侍刃'),(2,003,'枪刃'),(2,004,'战刃'),(2,005,'棍棒'),(2,006,'巨剑'),
                                                                                       (3,001,'突击步枪'),(3,002,'狙击枪'),(3,004,'弓弩'),
                                                                                       (4,001,'手枪'),(4,002,'手炮'),
                                                                                       (5,001,'增幅器'),
                                                                                       (6,001,'专精'),
                                                                                       (7,001,'模组'),
                                                                                       (8,001,'赋能')

create table tb_document(
                            `id` bigint NOT NULL AUTO_INCREMENT primary key ,
                            `uniposcod` varchar(9) not null ,
                            `titlename` varchar(50) not null ,
                            `titleimage` varchar(100) not null ,
                            `document` text not null ,
                            `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            unique (uniposcod)
);

show create table tb_seccommunity;

CREATE TABLE `tb_seccommunity` (
                                   `id` bigint NOT NULL AUTO_INCREMENT,
                                   `community_id` bigint DEFAULT NULL COMMENT '文本类型代码',
                                   `sec_community_id` int(4) unsigned zerofill DEFAULT NULL COMMENT '次级文本类型代码',
                                   `sec_community_info` varchar(128) DEFAULT NULL COMMENT '次级分类信息',
                                   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


select a.author_id,b.userid,b.userpermissions
from tb_user a,tb_userpermissions b
where a.author_id = ?;