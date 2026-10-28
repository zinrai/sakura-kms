# sakura-kms

A command-line tool for encrypting and decrypting data using [SAKURA Cloud KMS (Key Management Service)](https://cloud.sakura.ad.jp/products/kms).

## Usage

```bash
$ sakura-kms encrypt < plaintext > ciphertext
$ sakura-kms decrypt < ciphertext > plaintext
```

Both commands read stdin, write the result to stdout and report the operation on stderr. `encrypt` writes the ciphertext string returned by the KMS API as is; `decrypt` takes that string and writes the original bytes.

`sakura-kms version` prints the version.

## Configuration

Set the KMS key resource ID:

```bash
$ export SAKURA_KMS_KEY_ID="110000000000"
```

API credentials and the API endpoint are resolved by [saclient-go](https://github.com/sacloud/saclient-go). Set either static API keys:

```bash
$ export SAKURA_ACCESS_TOKEN="your-api-token"
$ export SAKURA_ACCESS_TOKEN_SECRET="your-api-secret"
```

or service principal credentials, with the private key given as a path or as `SAKURA_PRIVATE_KEY`:

```bash
$ export SAKURA_SERVICE_PRINCIPAL_ID="your-service-principal-id"
$ export SAKURA_SERVICE_PRINCIPAL_KEY_ID="your-key-id"
$ export SAKURA_PRIVATE_KEY_PATH="/path/to/private-key.pem"
```

## Requires

- A KMS key in SAKURA Cloud
- API credentials with access to that key

## License

This project is licensed under the [MIT License](./LICENSE).
