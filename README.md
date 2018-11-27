# make-shadow

The tool of making `/etc/shadow` 

```
Usage:
  make-shadow [options] name

Application Options:
      --min=days           The minimum password age
      --max=days           The maximum password age
      --warning=days       The number of days before a password is going to expire
      --inactivity=days    The number of days after a password has expired
      --expiration=days    The date of expiration of the account, expressed as the number of days since Jan 1, 1970
      --md5                MD5
      --sha256             SHA-256
      --sha512             SHA-512 (default)
  -h, --help               Show this help
```

Execute `make-shadow tsuty` command. And then enter the password. it outputs fields to stdout.

```bash
$ make-shadow tsuty
Enter Password: 
tusty:$6$XpfOLr2VPR5tlYBB$kLwFV6RTFn7vXaPrr3YrTNY/iiQDmOYCuK4gNrAawljLTNQOR2m549niokSnHoTbSA6ZZZFNa8DlaevwkXe7v1:17862::::::
```

Not so good, but useful way. You can enter the password from stdin.

```bash
$ echo "Password" | ./make-shadow tusty
tusty:$6$Y2IhxUgxLQY5LUoA$TyCEEdYWaNJJQ5OGSiq6oy7FTGYyTuupDfcBANZqF6aAkRvmUnXmLBlxNQtwNgwpVmq2QH.u21FS.fCBXa8G40:17862::::::
```

