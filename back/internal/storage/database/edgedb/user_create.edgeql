with user := (insert User {
    username := <str>$username,
    password := <str>$password,
    email := <str>$email
})
    select User {
        username, email
    } filter .id = user.id;