SELECT User {
    id, username, email, password
} FILTER .username = <str>$username;