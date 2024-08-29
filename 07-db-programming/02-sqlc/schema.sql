CREATE TABLE `students` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `fname` varchar(50) not null,
    `lname` varchar(50) not null,
    `date_of_birth` datetime not null,
    `email` varchar(50) not null,
    `address` varchar(50) not null,
    `gender` varchar(50) not null,
    PRIMARY KEY (`id`)
);