# banking-system

## Register user
Method POST```localhost:8080/api/auth/register```

# -> login (set Authorization token to headers)
Method POST```localhost:8080/api/auth/login```

# -> profile 
Method GET```localhost:8080/api/user/profile```

# -> create account 
Method POST ```localhost:8080/api/account```

# -> get account info
Method GET ```localhost:8080/api/account/:id```

# -> create transaction 
Method POST ```localhost:8080/api/transaction/```

# -> get transaction info by account id
Method GET ```localhost:8080/api/transaction/:id```
### Where id = accountSender id
