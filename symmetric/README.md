## Block Cipher

To generate a `block cipher key` using *OpenSSL*, you can use the **openssl rand command**, which generates a random key of a specified length. The key length depends on the block cipher algorithm you're planning to use, such as AES-128, AES-192, or AES-256.

#### Example: Generating a 256-bit AES key
To generate a 256-bit key (32 bytes) for AES, use:

    openssl rand -hex 32

This will output a 32-byte (64 hexadecimal characters) key suitable for AES-256 encryption.

#### Example: Generating a 128-bit AES key
For a 128-bit AES key (16 bytes), use:

    openssl rand -hex 16

The same applies to other key lengths, just adjust the byte size for your required algorithm.

To encrypt and decrypt data using a block cipher like AES with OpenSSL, you can follow these steps. I'll demonstrate how to use `AES-256 in CBC (Cipher Block Chaining)` mode, which is a commonly used block cipher mode.


Generate a random key and IV (Initialization Vector).
### 1. Generating a Key and IV
Generate a 256-bit AES key (32 bytes)

    openssl rand -hex 32 > aes_key.hex

#### Generate a 128-bit IV (16 bytes)

    openssl rand -hex 16 > aes_iv.hex

### 2. Encrypting Data
#### Encrypt with AES-256-CBC
Use the openssl enc command to encrypt a plaintext file. Replace plaintext.txt with your input file, encrypted.dat with your desired output file, and use the generated key and IV.

    openssl enc -aes-256-cbc -in plaintext.txt -out encrypted.dat -K $(cat aes_key.hex) -iv $(cat aes_iv.hex)

##### Parameters:
* `-aes-256-cbc`: Specifies the AES-256 encryption in CBC mode.
* `-in`: Input file to encrypt.
* `-out`: Output file for encrypted data.
* `-K`: Specifies the key in hexadecimal format.
* `-iv`: Specifies the IV in hexadecimal format.
### 3. Decrypting Data
#### Decrypt with AES-256-CBC
To decrypt the data, use the **same** `key` and `IV`:

    openssl enc -d -aes-256-cbc -in encrypted.dat -out decrypted.txt -K $(cat aes_key.hex) -iv $(cat aes_iv.hex)

##### Parameters:

* -d: Specifies decryption mode.

#### Example
Assuming you have a file called plaintext.txt that you want to encrypt:

#### Generate the key and IV:

    openssl rand -hex 32 > aes_key.hex
    openssl rand -hex 16 > aes_iv.hex

#### Encrypt the file:

    openssl enc -aes-256-cbc -in plaintext.txt -out encrypted.dat -K $(cat aes_key.hex) -iv $(cat aes_iv.hex)

#### Decrypt the file:

    openssl enc -d -aes-256-cbc -in encrypted.dat -out decrypted.txt -K $(cat aes_key.hex) -iv $(cat aes_iv.hex)

**Notes** Ensure that the key and IV are kept secret and secure. If anyone obtains them, they can decrypt your data.
*The `IV` does not need to be secret but should be unique for each encryption operation with the same key.*



## Stream cipher

Here are examples for some common stream ciphers:

### RC4 Stream Cipher Key
RC4 is a common stream cipher that uses a variable key length, typically between 40 and 2048 bits. Here’s how to generate a 128-bit (16-byte) key for RC4:

    openssl rand -hex 16

This will generate a 16-byte (32 hexadecimal characters) key for RC4.

### ChaCha20 Stream Cipher Key
ChaCha20, another popular stream cipher, requires a 256-bit (32-byte) key. You can generate this as follows:

    openssl rand -hex 32

This will give you a 32-byte (64 hexadecimal characters) key suitable for ChaCha20.

### Saving the Key to a File (Optional)
You can save the generated key to a file for later use:

    openssl rand 32 > chacha20.key

This will save a raw 32-byte key for ChaCha20 in the file chacha20.key.

### Encrynting and Descrypting
To encrypt and decrypt data using stream cipher keys (e.g., for RC4 or ChaCha20) with OpenSSL, you can use specific commands. I'll show you how to use RC4 and ChaCha20 for encryption and decryption.

#### RC4 Stream Cipher
##### Encrypt with RC4
To encrypt data using an RC4 key, use the openssl rc4 command. You must provide the key you generated and the plaintext file.

Example of encryption using a 128-bit (16-byte) key:

    openssl rc4 -in plaintext.txt -out encrypted.dat -K <your_key_in_hex>
Where:

* plaintext.txt is the file containing your plain text.
* encrypted.dat is the output encrypted file.
* <your_key_in_hex> is the key in hexadecimal format (without spaces).
Decrypt with RC4

To decrypt the file, you use the same command but swap the input/output files and use the same key:

    openssl rc4 -d -in encrypted.dat -out decrypted.txt -K <your_key_in_hex>

#### ChaCha20 Stream Cipher

ChaCha20 is available via OpenSSL 1.1.0 or later. The command structure is a bit different because it requires a nonce (IV) along with the key.

Encrypt with ChaCha20
First, generate a 96-bit nonce (12 bytes), which is required for ChaCha20:


    openssl rand -hex 12

Now, encrypt the data with ChaCha20, providing both the key and the nonce:

    openssl enc -chacha20 -in plaintext.txt -out encrypted.dat -K <your_key_in_hex> -iv <your_nonce_in_hex>
Where:

* plaintext.txt is the file containing your plain text.
* encrypted.dat is the output encrypted file.
* <your_key_in_hex> is the 32-byte (64 characters) key.
* <your_nonce_in_hex> is the 12-byte (24 characters) nonce.

##### Decrypt with ChaCha20
To decrypt, provide the same key and nonce used during encryption:

    openssl enc -d -chacha20 -in encrypted.dat -out decrypted.txt -K <your_key_in_hex> -iv <your_nonce_in_hex>


# Note

The difference between `block` and `stream` ciphers isn't in the key generation itself but in how the keys are applied during encryption.

Here’s how it breaks down:

## 1. Key Generation (Shared Process)
* Both block and stream ciphers rely on cryptographic keys that need to be securely generated. The length of the key typically depends on the encryption algorithm.
* Command for key generation:
    * For a 256-bit key (common for many modern ciphers): openssl rand -hex 32
    * For a 128-bit key: openssl rand -hex 16

In both cases, OpenSSL generates a random key, but the way the key is used differs based on whether you're encrypting with a block or stream cipher.

## 2. Block Cipher Usage
Block ciphers (like AES) break the input data into fixed-size blocks (e.g., 128 bits for AES) and encrypt each block, often using techniques like CBC (Cipher Block Chaining) or GCM (Galois/Counter Mode) to ensure that blocks are linked in some way (i.e., so that identical plaintext blocks don't produce identical ciphertext blocks).

### Encryption example (AES-256 CBC mode):

    openssl enc -aes-256-cbc -in plaintext.txt -out encrypted.dat -K <your_key_in_hex> -iv <your_iv_in_hex>

* Key length: 32 bytes (256 bits for AES-256).

* IV: Initialization Vector (IV) is required, which is unique for each encryption but not secret.

## 3. Stream Cipher Usage
Stream ciphers (like RC4 or ChaCha20) work by encrypting the data bit by bit or byte by byte, generating a pseudorandom keystream that is XOR'd with the plaintext.

* Encryption example (ChaCha20):

    openssl enc -chacha20 -in plaintext.txt -out encrypted.dat -K <your_key_in_hex> -iv <your_nonce_in_hex>

* Key length: 32 bytes (256 bits for ChaCha20).

* Nonce: A 12-byte (96-bit) nonce is required for ChaCha20 but must not be reused with the same key.

## Summary
* Key generation (openssl rand -hex 32) is common for both block and stream ciphers if you're working with 256-bit keys.
* Encryption process differs:
    * Block ciphers use fixed-size blocks and may require an IV (e.g., AES-256-CBC).
    * Stream ciphers generate a pseudorandom keystream and require a nonce (e.g., ChaCha20).
