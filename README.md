
# Experimental code for interacting with the Ethereum Network

*Warning*: USE AT YOUR OWN RISK - use of this code can cause loss of funds.

## Getting Started

```
$ make bin/cli
$ export ETHEREUM_SERVER=ws://server.example.net:8546/
$ ./bin/cli client status
Client in sync
```

## Creating an account, listing, and checking passphrase

```
$ export WALLET_KEYSTORE=~/.keystore
$ ./bin/cli account new
Enter Passphrase:
Repeat:
Account: 0xf068fdA9ae6ade96472Ed6b355277B595b84f6AE
```

```
$ ./bin/cli account list
Account: 0xf068fdA9ae6ade96472Ed6b355277B595b84f6AE
```

```
$ export ETHEREUM_ADDRESS=0xf068fdA9ae6ade96472Ed6b355277B595b84f6AE
$ ./bin/cli account unlock
Enter Passphrase: 
Unlocked 0xf068fdA9ae6ade96472Ed6b355277B595b84f6AE
```
