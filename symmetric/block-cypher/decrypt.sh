#!/bin/bash


openssl enc -d -aes-256-cbc -in encrypted.dat -out decrypted.txt -K $(cat key.hex) -iv $(cat iv.hex)