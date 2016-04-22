# Counterparty VanityGen
Small script to generate vanity address that is compatible with counterwallet.

Generates a 16 byte random hash and loops to concat a 4 byte nonce to it
and test if `m/0'/0/0` starts with desired (case insensitve) prefix.

Result will be a counterwallet compatible mnemonic.

## Usage
`./cp-vanity-gen -threads=4 1xcp`

Latest compiled binaries for your OS/arch can be found on the releases page: https://github.com/rubensayshi/cp-vanitygen/releases

## Building Binaries
`./build.sh` will build binaries in `./build` for all platforms

## Release
`/release.sh v0.0.1` will tag, push tag, create github release and upload binaries
`./release-upload.sh v0.0.1` is used by `./release.sh`
