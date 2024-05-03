# simple-blockchain

A port an old project to golang.

https://github.com/smhmayboudi/simple_blockchain/

## Help

- https://blog.logrocket.com/build-blockchain-with-go/
- https://jeiwan.net/posts/building-blockchain-in-go-part-1/
- https://jeiwan.net/posts/building-blockchain-in-go-part-2/
- https://jeiwan.net/posts/building-blockchain-in-go-part-3/
- https://jeiwan.net/posts/building-blockchain-in-go-part-4/
- https://jeiwan.net/posts/building-blockchain-in-go-part-5/
- https://jeiwan.net/posts/building-blockchain-in-go-part-6/
- https://jeiwan.net/posts/building-blockchain-in-go-part-7/

- https://github.com/Jeiwan/blockchain_go

- https://www.nikitv.ir/blog/post_40_bc1/
- https://www.nikitv.ir/blog/post_41_bc2/
- https://www.nikitv.ir/blog/post_42_bc3/
- https://www.nikitv.ir/blog/post_43_bc4/
- https://www.nikitv.ir/blog/post_44_bc5/
- https://www.nikitv.ir/blog/post_45_bc6/
- https://www.nikitv.ir/blog/post_46_bc7/

- https://medium.com/coinmonks/creating-a-blockchain-part-1-transport-layer-327dc2a85df9
- https://medium.com/coinmonks/creating-a-blockchain-part-2-blocks-and-hashing-ae7dc7549b08?source=user_profile---------8----------------------------
- https://medium.com/coinmonks/creating-a-blockchain-part-3-crypto-key-pairs-signature-1a76d4dd9856?source=user_profile---------7----------------------------
- https://medium.com/coinmonks/creating-a-blockchain-part-4-block-and-tx-signing-and-verification-2d94be9bf2c0?source=user_profile---------6----------------------------
- https://medium.com/coinmonks/creating-a-blockchain-part-5-blockchain-data-structure-b0d91054b40d?source=user_profile---------4----------------------------
- https://medium.com/coinmonks/creating-a-blockchain-part-6-transaction-mempool-and-tx-encoding-a1581479449e?source=user_profile---------3----------------------------
- https://medium.com/@opensiddhu993/creating-a-blockchain-part-7-rpc-protocol-78469a2757b4?source=user_profile---------2----------------------------
- https://medium.com/coinmonks/creating-a-blockchain-part-8-creating-block-7431c593bd18?source=user_profile---------1----------------------------
- https://medium.com/coinmonks/creating-a-blockchain-part-9-broadcast-blocks-73ad5e4de121?source=user_profile---------0----------------------------

## TODO

- Complete the steps
- migration from [BoltDB](https://github.com/boltdb/bolt) to GO [LevelDB](https://github.com/syndtr/goleveldb)

## RUN

```SHELL
$ NODE_ID=3000 go run . createwallet
Your new address: 16PPYFiioDtFEHBYTd81Dcm3vBdTAKLCCj

$ NODE_ID=3000 go run . createwallet
Your new address: 1BNfobn516oN7dB8dcEY9SKkBxetmUn62a

$ NODE_ID=3000 go run . createblockchain -address 16PPYFiioDtFEHBYTd81Dcm3vBdTAKLCCj
45e5d430eb64a03f23feaee9e711b6ad071104c1130d59e930c4c94693d13f3a

Done!

$ NODE_ID=3000 go run . getbalance -address 16PPYFiioDtFEHBYTd81Dcm3vBdTAKLCCj
Balance of '16PPYFiioDtFEHBYTd81Dcm3vBdTAKLCCj': 10

$ NODE_ID=3000 go run . getbalance -address 1BNfobn516oN7dB8dcEY9SKkBxetmUn62a
Balance of '1BNfobn516oN7dB8dcEY9SKkBxetmUn62a': 0

$ NODE_ID=3000 go run . send -from 16PPYFiioDtFEHBYTd81Dcm3vBdTAKLCCj -to 1BNfobn516oN7dB8dcEY9SKkBxetmUn62a -amount 2
```