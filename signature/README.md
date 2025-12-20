- generate private key PKCS#8

$ openssl genpkey -algorithm RSA -out signature.pem -pkeyopt rsa_keygen_bits:4096

- generate public key from private key

$ openssl pkey -in signature.pem -pubout -out signature.pub.pem