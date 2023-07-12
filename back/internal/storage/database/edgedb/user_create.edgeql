with user := (insert User {
    username := <str>$username,
    password := <str>$password,
    email := <str>$email
})
  select user {
    username,
    email
  }