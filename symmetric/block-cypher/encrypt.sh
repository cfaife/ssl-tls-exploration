#!/bin/bash

openssl enc -aes-256-cbc -in plaintext.txt -out encrypted.dat -K $(cat key.hex) -iv $(cat iv.hex)