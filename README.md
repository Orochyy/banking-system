# banking-system

## Register user

Method POST```localhost:8080/api/auth/register```

### For register need send next parameters
Example:
```
   name:        test1
   email:       test1@gmail.com
   password:    test1
 ```

# -> login (set Authorization token to headers)

Method POST```localhost:8080/api/auth/login```

### For login need send next parameters
Example:
```
    name:       test1
    password:   test1
```

# -> profile

Method GET```localhost:8080/api/user/profile```

# -> create account

Method POST ```localhost:8080/api/account```

### For create account need send next parameters
Example:
```
    currency:   MXN
    amount:     13000
```

# -> get account info

Method GET ```localhost:8080/api/account/:id```

# -> create transaction

Method POST ```localhost:8080/api/transaction```

### For create transaction need send next parameters
Example:
```
    AccountSender:          {account:id}
    AccountRecipient:       {account:id}
    currency:               MXN
    amount:                 13
    type:                   transfer
```

# -> get transaction info by account id

Method GET ```localhost:8080/api/transaction/:id```

### Where id = accountSender id
