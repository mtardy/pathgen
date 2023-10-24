# Pathgen

Pathgen is a Golang CLI tool for generating arbitrary long paths on Linux

## Installation

You need Golang to be installed in your environment. Pathgen has no external
dependencies and only use the Go standard library.

```bash
go install github.com/mtardy/pathgen@latest
```

## Usage

```text
Usage:
  pathgen [flags] length

Flags:
  -b string
        binary to put at the end of the random path, used when write is enabled
  -c    cleanup everything under "/tmp/pathgen"
  -l    create a symlink at "/tmp/pathgen/exe" to the target file, used when write is enabled (default true)
  -p string
        prefix for the random path, you will need to cleanup manually (default "/tmp/pathgen")
  -s string
        suffix of the random path, name of the copied binary (default "bin")
  -w    write the filepath on the filesystem
```

For example, to generate a path of 512 bytes and put the binary `nc` at the end
of random path under the name `bin` do:

```shell-session
pathgen -w -b $(which nc) 512
```

The output should be similar to:
```
/tmp/pathgen/fmblkzrkyvzbokvvwqxuzghimgzuhshccwetfneqsvvypjqocjnbzpvjdtvuszkjzugvrnhibjmfimniyujhagnkpykxcjrlfuelynkhnwfmpkztuvugynvaooktegdlpbtjumpzdznycpqsunbpeqomvdccokcbqbbgcnjllwncupxfoyiuofuywmiituxtsyzuewagurliurxfrdhnfbwvrizniuxwhajabtnuuvlsdhfgtxebzopisjydaoe/ouiefagfyastdxfjkxexpgfryogkqdrwtxpvhgoutwuqbehzhbmimoezrqsbkxfgjfpamonxvkkkohmmkehmmnxamuvmetgmlaxekvunksaqrtukatkcikbvmearvyvsksnsyrgdmcbkmdelwbuzlqrguzasaiwyjxhjzlotwowqdnfvqavrwvnpuzebwlkrcidxiitjvahpxlsysulhbeadmhntpfwnrifvssbjzzzumey/bin
```

You will find a symlink at `/tmp/pathgen/exe` that points to the binary that you
can use to trigger execution from the arbitrary path.

```shell-session
ls -l /tmp/pathgen
```

The output should be similar to:
```
lrwxrwxrwx 1 mahe mahe  512 Oct 24 17:08 exe -> /tmp/pathgen/fmblkzrkyvzbokvvwqxuzghimgzuhshccwetfneqsvvypjqocjnbzpvjdtvuszkjzugvrnhibjmfimniyujhagnkpykxcjrlfuelynkhnwfmpkztuvugynvaooktegdlpbtjumpzdznycpqsunbpeqomvdccokcbqbbgcnjllwncupxfoyiuofuywmiituxtsyzuewagurliurxfrdhnfbwvrizniuxwhajabtnuuvlsdhfgtxebzopisjydaoe/ouiefagfyastdxfjkxexpgfryogkqdrwtxpvhgoutwuqbehzhbmimoezrqsbkxfgjfpamonxvkkkohmmkehmmnxamuvmetgmlaxekvunksaqrtukatkcikbvmearvyvsksnsyrgdmcbkmdelwbuzlqrguzasaiwyjxhjzlotwowqdnfvqavrwvnpuzebwlkrcidxiitjvahpxlsysulhbeadmhntpfwnrifvssbjzzzumey/bin
drwx------ 3 mahe mahe 4096 Oct 24 17:08 fmblkzrkyvzbokvvwqxuzghimgzuhshccwetfneqsvvypjqocjnbzpvjdtvuszkjzugvrnhibjmfimniyujhagnkpykxcjrlfuelynkhnwfmpkztuvugynvaooktegdlpbtjumpzdznycpqsunbpeqomvdccokcbqbbgcnjllwncupxfoyiuofuywmiituxtsyzuewagurliurxfrdhnfbwvrizniuxwhajabtnuuvlsdhfgtxebzopisjydaoe
```

## License

[MIT](https://choosealicense.com/licenses/mit/)