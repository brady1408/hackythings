
# openssl

## Convert PFX to a CRT and RSA key for AWS

note: you will need the password for the PFX to complete these operations.

**Convert the PFX to a PEM**

    openssl pkcs12 -in filename.pfx -out filename.pem

This should get you the cert and the chain within the .pem file but you will need to convert the key to a RSA Key.

**Convert the encrypted key to an RSA key**

    openssl rsa -in filename.pem -out filename.key

This should provide you with a .key file which has the RSA key in it that AWS will accept.
