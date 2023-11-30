create table if not exists `account` (
    `account_id` int primary key auto_increment,
    `name` varchar(20) not null unique,
    `balance` decimal(20,2) not null,
    `created` datetime not null default current_timestamp,
    `modified` datetime null on update current_timestamp
);


create table if not exists `category` (
    `category_id` int primary key auto_increment,
    `name` varchar(20) not null unique,
    `created` datetime not null default current_timestamp,
    `modified` datetime null on update current_timestamp
    );

create table if not exists `transaction` (
    `transaction_id` bigint primary key auto_increment,
    `account_id` int,
    `category_id` int,
    `description` varchar(250) not null,
    `transaction_type` enum ('Credit', 'Debit') not null,
    `amount` decimal(20,2) not null,
    `charges` decimal(10,2) not null default '0.0',
    `created` datetime not null default current_timestamp,
    `modified` datetime null on update current_timestamp,

    foreign key (`account_id`) references `account`(`account_id`),
    foreign key (`category_id`) references `category`(`category_id`),
    key (`created`)
    );

insert ignore into `account` (`name`) values
('Cash'),
('Mpesa'),
('Co-op'),
('KCB'),
('Stima'),
('CIC');

insert ignore into `category`(`name`) values
('Food'),
('Transport'),
('Entertainment'),
('Debt'),
('Savings'),
('Health'),
('Utilities');