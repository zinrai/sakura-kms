# sakura-kms

A command-line tool for encrypting and decrypting data using [SAKURA Cloud KMS (Key Management Service)](https://cloud.sakura.ad.jp/products/kms).

## Features

- Encrypt data from stdin using SAKURA Cloud KMS
- Decrypt data from stdin using SAKURA Cloud KMS

## Installation

```bash
$ go install github.com/zinrai/sakura-kms@latest
```

## Prerequisites

1. A SAKURA Cloud account
2. A KMS key created in SAKURA Cloud
3. API credentials with KMS access permissions

## Configuration

Set the following environment variables:

```bash
export SAKURACLOUD_ACCESS_TOKEN="your-api-token"
export SAKURACLOUD_ACCESS_TOKEN_SECRET="your-api-secret"
export SAKURACLOUD_ZONE="is1a"  # or tk1a, tk1b, etc.
```

## Usage

### Encrypt

```bash
$ cat plaintext.txt | sakura-kms encrypt -output encrypted.bin -resource-id 110000000000
```

### Decrypt

```bash
$ cat encrypted.bin | sakura-kms decrypt -output plaintext.enc.txt -resource-id 110000000000
```

### Examples

Encrypt a database password:

```bash
$ echo "my-secret-password" | sakura-kms encrypt -output db-password.enc -resource-id 110000000000
```

Decrypt and use in a pipeline:

```bash
$ cat db-password.enc | sakura-kms decrypt -output /dev/stdout -resource-id 110000000000 | my-application --password-stdin
```

Encrypt a configuration file:

```bash
$ cat config.json | sakura-kms encrypt -output config.json.enc -resource-id 110000000000
```

## Command Reference

### encrypt

Encrypts data from stdin and writes the ciphertext to the specified output file.

**Flags:**
- `-output` (required): Output file path
- `-resource-id` (required): KMS key resource ID

### decrypt

Decrypts data from stdin and writes the plaintext to the specified output file.

**Flags:**
- `-output` (required): Output file path
- `-resource-id` (required): KMS key resource ID

## How It Works

1. Reads data from stdin
2. Communicates with SAKURA Cloud KMS API using Basic Authentication
3. Handles Base64 encoding/decoding automatically
4. Writes the output to the specified file

The encrypted output file contains the `Cipher` field from the API response as-is. The decrypted output file contains the Base64-decoded `Plain` field from the API response.

## License

This project is licensed under the [MIT License](./LICENSE).
