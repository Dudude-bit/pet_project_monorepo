DEFINE TABLE users SCHEMAFULL;

DEFINE FIELD username ON TABLE users TYPE string;
DEFINE FIELD email ON TABLE users TYPE string;
DEFINE FIELD password ON TABLE users TYPE string;

DEFINE INDEX usernameUsersIndex ON TABLE users COLUMNS username UNIQUE;
DEFINE INDEX emailUsersIndex ON TABLE users COLUMNS email UNIQUE;