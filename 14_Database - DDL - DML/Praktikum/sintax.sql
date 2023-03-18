CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  status SMALLINT,
  dob DATE,
  gender CHAR(1),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE product_type (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE operators (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);


CREATE TABLE payment_method (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  status SMALLINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE product (
  id INT NOT NULL AUTO_INCREMENT,
  product_type_id INT NOT NULL,
  operator_id INT NOT NULL,
  code VARCHAR(25) NOT NULL,
  name VARCHAR(100) NOT NULL,
  price DECIMAL(25.2) NOT NULL,
  stock INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (product_type_id) REFERENCES product_type(id),
  FOREIGN KEY (operator_id) REFERENCES operators(id)
);

CREATE TABLE transaction (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  payment_method_id INT NOT NULL,
  status VARCHAR(255),
  total_qty int(11),
  total_price DECIMAL(10,2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);


CREATE TABLE transaction_detail (
  id INT NOT NULL AUTO_INCREMENT,
  transaction_id INT NOT NULL,
  product_id INT NOT NULL,
  status VARCHAR(10),
  quantity INT(11),
  price DECIMAL(10,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (transaction_id) REFERENCES transaction(id),
  FOREIGN KEY (product_id) REFERENCES product(id)
);

CREATE TABLE kurir (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE payment_method_description (
  id INT NOT NULL AUTO_INCREMENT,
  description VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

ALTER TABLE payment_method 
ADD description_id INT NOT NULL;

ALTER TABLE payment_method 
ADD FOREIGN KEY (description_id) REFERENCES payment_method_description(id)
;

CREATE TABLE alamat (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  alamat VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE users
ADD alamat_id INT NOT NULL
;
ALTER TABLE users
ADD FOREIGN KEY (alamat_id) REFERENCES alamat(id)
;

CREATE TABLE user_payment_method_detail (
  user_id INT NOT NULL,
  payment_method_id INT NOT NULL,
  PRIMARY KEY (user_id, payment_method_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);