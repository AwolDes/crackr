# crackr
Crackr - A CLI to crack hashes using a dictionary attack

## Exmaple Usage

### Crack 1 password against 1 dictionary:

`crackr -f passwords.txt -h 5f4dcc3b5aa765d61d8327deb882cf99`

### Crack multiple passwords against 1 dictionary:

`crackr -f passwords.txt -hf hashes.txt`

### Crack 1 password against multiple dictionaries:

`crackr -h 5f4dcc3b5aa765d61d8327deb882cf99 -f dictionaries`

### Crack multiple passwords against multiple dictionaries

`crackr -hf hashes.txt -d dicitonaries`



