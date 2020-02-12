# kanayomi
A tool to convert English words to katakana(Japanese alphabets) reading.

# Installation from Source

You need to clone the repository and compile it.

```
git clone https://github.com/ironsand/kanayomi.git
cd kanayomi
go build -o kanayomi main.go bep_dic.go
mv kanayomi /usr/local/bin
```
Change `/usr/local/bin` suitable for your environments.

# How to use

The command takes a english word as an argument and returns katakana. If corresponding katakana doesn't exist it returns nothing.

```
$ kanayomi english
イングリッシュ
$ kanayomi word_not_exist
$ 
```

# Licence

Originally the code is written by [@metropolis](https://ja.stackoverflow.com/users/16894/metropolis) in [Stackoverlow](https://ja.stackoverflow.com/a/62939/3271)

The translation data `bep-eng.dic` is from [Bilingual Emacspeak Project](http://www.argv.org/bep/)