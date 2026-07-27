# sakura-kms

A command-line tool for encrypting and decrypting data using [SAKURA Cloud KMS (Key Management Service)](https://cloud.sakura.ad.jp/products/kms).

## Features

- Encrypt data from stdin using SAKURA Cloud KMS
- Decrypt data from stdin using SAKURA Cloud KMS

## Prerequisites

1. A SAKURA Cloud account
2. A KMS key created in SAKURA Cloud
3. API credentials with KMS access permissions

## Configuration

Set the KMS key resource ID:

```bash
$ export SAKURA_KMS_KEY_ID="110000000000"
```

API credentials are resolved by [saclient-go](https://github.com/sacloud/saclient-go). Set either static API keys:

```bash
$ export SAKURA_ACCESS_TOKEN="your-api-token"
$ export SAKURA_ACCESS_TOKEN_SECRET="your-api-secret"
```

or service principal credentials:

```bash
$ export SAKURA_SERVICE_PRINCIPAL_ID="your-service-principal-id"
$ export SAKURA_SERVICE_PRINCIPAL_KEY_ID="your-key-id"
$ export SAKURA_PRIVATE_KEY_PATH="/path/to/private-key.pem"
```

The legacy `SAKURACLOUD_` prefixed names are also accepted.

## Usage

### Encrypt

```bash
$ sakura-kms encrypt < plaintext.txt > encrypted.bin
```

### Decrypt

```bash
$ sakura-kms decrypt < encrypted.bin > plaintext.txt
```

### Examples

Encrypt a database password:

```bash
$ echo "my-secret-password" | sakura-kms encrypt > db-password.enc
```

Decrypt and use in a pipeline:

```bash
$ sakura-kms decrypt < db-password.enc | my-application --password-stdin
```

Encrypt a configuration file:

```bash
$ sakura-kms encrypt < config.json > config.json.enc
```

## Command Reference

### encrypt

Encrypts data from stdin and writes the ciphertext to stdout. Reports the operation on stderr.

**Flags:**

- `-zone` (optional): SAKURA Cloud zone (default: "is1a")

### decrypt

Decrypts data from stdin and writes the plaintext to stdout. Reports the operation on stderr.

**Flags:**

- `-zone` (optional): SAKURA Cloud zone (default: "is1a")

### version

Prints the version.

## How It Works

1. Reads data from stdin
2. Communicates with SAKURA Cloud KMS API, authenticating with static API keys or a service principal
3. Handles Base64 encoding/decoding automatically
4. Writes the result to stdout

The encrypted output contains the `Cipher` field from the API response as-is. The decrypted output contains the Base64-decoded `Plain` field from the API response.

## License

This project is licensed under the [MIT License](./LICENSE).
