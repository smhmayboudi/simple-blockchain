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

go run . createblockchain -address 15AedVzQJqeeS5UDcBnAtKfwe8o7gCsvJL
go run . getbalance -address 15AedVzQJqeeS5UDcBnAtKfwe8o7gCsvJL
go run . getbalance -address 1AweBuGtrBZf446nZmPAdaMj5Vk3kiEiyP
go run . send -from 15AedVzQJqeeS5UDcBnAtKfwe8o7gCsvJL -to 1AweBuGtrBZf446nZmPAdaMj5Vk3kiEiyP -amount 2

15AedVzQJqeeS5UDcBnAtKfwe8o7gCsvJL
1AweBuGtrBZf446nZmPAdaMj5Vk3kiEiyP
19QLTKXnqqyFXwP2taC5vQ2NHeqZzW4NgU