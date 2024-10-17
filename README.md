# Introduction

## **Cryptography** vs **Encryption**: 

*Cryptography* is the science of concealing or obscuring transmitted information messages with a secret code. 
This is about studying methods to keep a message secret between two parties (like symmetric and asymmetric keys)

*Encryption* is the way to encrypt and decrypt data, and this is about the process itself.


Modern cryptosystems are far more advanced but still function in similar ways. Most cryptosystems begin with an unencrypted message known as plaintext, which is then *encrypted* into an indecipherable code known as *ciphertext* by using one or more *encryption keys*.

This *ciphertext* is then transmitted to a recipient. If the *ciphertext* is intercepted and the *encryption algorithm* is strong, the *ciphertext* is useless to any unauthorized eavesdroppers because they will not be able to break the code. However, the intended recipient will easily be able to decipher the text, assuming that they have the correct decryption key. 

*In cryptography, a **cipher** is simply a method or process used to take readable information `(called "plaintext")` and turn it into something that looks random and unreadable `(called "ciphertext")`. This helps to keep the information **secret** from anyone who doesn’t have permission to see it.*

**Here's how it works:**

* **Plaintext**: This is the original message you want to protect (like "Hello, World!").
* **Encryption**: A cipher scrambles the message using a special key (kind of like a password). After encryption, the message becomes unreadable (like "Xy9#@!p2%").
* **Ciphertext**: This is the scrambled, unreadable message.
* **Decryption**: When someone with the correct key wants to read the message, they use the cipher again to turn the ciphertext back into plaintext.

So a **cipher** is just a set of *`rules` (or a process)* that helps you *scramble* and *unscramble* messages. Different `ciphers` have different rules, but they all follow the same basic idea.








## 3 categories of encryption

### Symmetric cryptography algorithms

Also known as *private key cryptography*, *secret key cryptography* or *single-key encryption*, symmetric key encryption uses only one key for both the encryption process and decryption process. For these types of systems, each user must have access to the same private key.

Private keys might be shared either through a previously established secure communication channel like a private courier or secured line or, more practically, a secure key exchange method like the Diffie-Hellman key agreement. 

There are 2 types of symmetric key algorithms:

* **Block cipher**: In a block cipher, the cipher algorithm works on a fixed-size block of data. For example, if the block size is eight, eight bytes of plaintext are encrypted at a time. Normally, the user’s interface to the encrypt/decrypt operation handles data longer than the block size by repeatedly calling the low-level cipher function.
* **Stream cipher**: Stream ciphers do not work on a block basis, but rather convert one bit (or one byte) of data at a time. Basically, a stream cipher generates a keystream based on the provided key. The generated keystream is then XORed with the plaintext data.

### Asymmetric cryptography algorithms 


In asymmetric encryption, a pair of keys is used: one secret key and one public key. For this reason, these algorithms are also referred to as public key algorithms. Public key cryptography is considered to be more secure than symmetric encryption techniques because even though one key is publicly available, an encrypted message can only be decrypted with the intended recipient’s private key.

Some examples of asymmetrical cryptography include the following:

* **RSA**: Named for its founders—Rivest, Shamier and Adleman—in 1977, the RSA algorithm is one of the oldest widely used public key cryptosystems used for secure data transmission. 
* **ECC**: Elliptic curve cryptography is an advanced form of asymmetric encryption that uses the algebraic structures of elliptic curves to create strong cryptographic keys. 

### Hash functions. 




# Reference

https://www.ibm.com/think/topics/cryptography-types