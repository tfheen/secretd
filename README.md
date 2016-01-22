
Build instructions:

- install gb, `apt-get install gb`

- update vendored source, `gb vendor update --all`

- build, `gb build all`

You should now have binaries in `bin/`

- create a postgres database for secretd, the default is `secrets`,
`createdb secrets`.

- load various fixtures into the database:

```
psql secrets < doc/users.sql
psql secrets < doc/tables.sql
psql secrets < doc/functions.sql
psql secrets < doc/views.sql
psql secrets < doc/data.sql
```
