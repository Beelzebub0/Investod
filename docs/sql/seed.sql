INSERT INTO `user`
(`id`, `age`, `name`, `job`)
VALUES
(1, 20, 'Ted', 'SJW'),
(2, 25, 'Sebastian', 'Keyboard Warrior');

INSERT INTO `transaction`
(`id`, `fk_category_id`, `fk_user_id`, `amount`, `name`)
VALUES
(1, 1, 1, 1000, 'Outflow');

INSERT INTO `categories`
(`id`, `name`, `level`, `notes`)
VALUES
(1, 'Food', 1, 'this is notes');

INSERT INTO `deposit`
(`id`, `fk_user_id`, `name`, `amount`, `notes`)
VALUES
(1, 1, 'Bank 1', 1000, 'test');