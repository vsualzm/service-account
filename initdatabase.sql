create table account (
	id serial primary key,
	nama varchar(100) not null,
	email varchar(100) not null,
	password_hash varchar(100) not null,
	roles varchar(100) not null,
	nik varchar(100) not null unique,
	no_hp varchar(100) not null unique,
	saldo numeric(15,2) default 0,
	no_rekening varchar(100) not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	update_at TIMESTAMP default CURRENT_TIMESTAMP
);
create table transaction (
	id serial primary key,
	account_id int not null,
	no_rekening_to varchar(100) not null,
	code_transaction varchar(100) not null,
	total_amount numeric(15,2) default 0,
	status varchar(20) not null,
	remark varchar(20) not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	update_at TIMESTAMP default CURRENT_TIMESTAMP,
	CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES account (id) ON DELETE CASCADE
);
