# make-shadow [![Build Status](https://travis-ci.com/tsuty/make-shadow.svg?branch=master)](https://travis-ci.com/tsuty/make-shadow)

The tool of making `/etc/shadow`

```
Usage:
  make-shadow [options] [name]

Application Options:
      --min=days           The minimum password age
      --max=days           The maximum password age
      --warning=days       The number of days before a password is going to expire
      --inactivity=days    The number of days after a password has expired
      --expiration=days    The date of expiration of the account, expressed as the number of days since Jan 1, 1970
      --md5                MD5
      --sha256             SHA-256
      --sha512             SHA-512 (default)
      --only-encrypt       Only encrypt password
  -h, --help               Show this help
```

Execute `make-shadow` command. And then enter the password. it outputs fields to stdout.

```bash
$ ./make-shadow tsuty
Enter Password:
tusty:$6$XpfOLr2VPR5tlYBB$kLwFV6RTFn7vXaPrr3YrTNY/iiQDmOYCuK4gNrAawljLTNQOR2m549niokSnHoTbSA6ZZZFNa8DlaevwkXe7v1:17862::::::

$ ./make-shadow --only-encrypt
Enter Password:
$6$hunZRG/CqxZJU0wm$KJP1KYP0No5m3NPRn8zgKdQM1td8qe.lCmgN1HoUzBWQExpIxygJguRNQswjfxGW6UVs3PiyK4cbJnJspj/Jz0
```

Not so good, but useful way. You can enter the password from stdin.

```bash
$ echo "Password" | make-shadow tusty
tusty:$6$Y2IhxUgxLQY5LUoA$TyCEEdYWaNJJQ5OGSiq6oy7FTGYyTuupDfcBANZqF6aAkRvmUnXmLBlxNQtwNgwpVmq2QH.u21FS.fCBXa8G40:17862::::::

$ echo "Password" | make-shadow --only-encrypt
$6$fnsEXhv/cs38Hf0.$BTAQe9jiTcbeBZ5Gild5WDmldFDFYUsH7NUaZqLFC.kqdHMstW.3Ije6ddBhioqUMVLnhGxIOOYYFn2JfEaVP.
```
