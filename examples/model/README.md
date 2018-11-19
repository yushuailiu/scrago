```sql
create table user (
  `id` int not null auto_increment comment '自增id',
  `user_name` varchar(128) not null,
  `avatar` varchar(512) not null,
  `type` tinyint not null,
  `intro` varchar(1024) not null,
  `follow_count` int not null,
  `fans_count` int not null,
  `user_type` tinyint not null,
  `is_vip` tinyint not null,
  `pubshare_count` int not null,
  `hot_uk` bigint not null,
  `album_count` int not null,
  `updated_at` datetime null default null on update current_timestamp comment '更新时间',
  `created_at` datetime not null default current_timestamp comment '创建时间',
  primary key (`id`)	
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```