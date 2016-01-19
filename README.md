# Spellrcheck

Is a simple spellchecking webservice. It can check the spelling of a whole text for spellerrors or make suggestions for a specific misspelled word.

# Methods

1. `/check` type **POST**
  * lang : *'en_US', 'de_DE', 'bg_BG'* ISO 639 language code and an optional two letter ISO 3166 country code
  * text : *<long text>*

2. `/suggest` type **POST**
  * lang : *'en_US', 'de_DE', 'bg_BG'* ISO 639 language code and an optional two letter ISO 3166 country code
  * word : *<single_word>*

# Building

To build the server run:

```bash
go build spellrcheck.go
```

Make shure your have [go-aspell](http://github.com/trustmaster/go-aspell) installed and also aspell with desired language packages

# Running

Just execute `./spellrcheck`

# LICENSE

```
	DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
``` TO.
```