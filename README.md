### Ethereum Total Transactions Amount
HTTP service providing sum of transaction amounts for a given ETH (Ethereum) block number.  
[Etherscan](https://etherscan.io/) is a data source.

#### Request:
    The service provides single HTTP endpoint:
        GET /api/block/<block_number>/total
    where <block_number> is a requested block number in decimal format.

    Request example: /api/block/11509797/total
#### Response:
    The response is returned as a JSON document with two fields:
    - “transactions” - total count of transactions in the block
    - “amount” - total amount of transaction values (in Ether)

    Response body example:
    {"transactions":155,"amount":2.285404805647828}

#### To run:
    docker-compose up