package darwinia

const MSGLINE_MESSAGER = `[
    {
        "name": "receiveMessage",
        "outputs": [],
        "type": "function",
        "stateMutability": "nonpayable",
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_srcAppChainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "_remoteAppAddress",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_localAppAddress",
                "type": "address"
            },
            {
                "internalType": "bytes",
                "name": "_message",
                "type": "bytes"
            }
        ]
    }
]`

const XTOKEN_ISSUING_NEXT = `[
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_remoteChainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "_originalToken",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_originalSender",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_recipient",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_rollbackAccount",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "_amount",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "_nonce",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "_extData",
                "type": "bytes"
            }
        ],
        "name": "issue",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
	{
	  "inputs": [
		{
		  "internalType": "address",
		  "name": "_xToken",
		  "type": "address"
		},
		{
		  "internalType": "address",
		  "name": "_recipient",
		  "type": "address"
		},
		{
		  "internalType": "address",
		  "name": "_rollbackAccount",
		  "type": "address"
		},
		{
		  "internalType": "uint256",
		  "name": "_amount",
		  "type": "uint256"
		},
		{
		  "internalType": "uint256",
		  "name": "_nonce",
		  "type": "uint256"
		},
		{
		  "internalType": "bytes",
		  "name": "_extData",
		  "type": "bytes"
		},
		{
		  "internalType": "bytes",
		  "name": "_extParams",
		  "type": "bytes"
		}
	  ],
	  "name": "burnAndXUnlock",
	  "outputs": [
		{
		  "internalType": "bytes32",
		  "name": "transferId",
		  "type": "bytes32"
		}
	  ],
	  "stateMutability": "payable",
	  "type": "function"
	}
]`

const XTOKEN_BACKING_NEXT = `[
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_remoteChainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "_originalToken",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_recipient",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_rollbackAccount",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "_amount",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "_nonce",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "_extData",
                "type": "bytes"
            },
            {
                "internalType": "bytes",
                "name": "_extParams",
                "type": "bytes"
            }
        ],
        "name": "lockAndXIssue",
        "outputs": [
            {
                "internalType": "bytes32",
                "name": "transferId",
                "type": "bytes32"
            }
        ],
        "stateMutability": "payable",
        "type": "function"
    },
	{
	  "inputs": [
		{
		  "internalType": "uint256",
		  "name": "_remoteChainId",
		  "type": "uint256"
		},
		{
		  "internalType": "address",
		  "name": "_originalToken",
		  "type": "address"
		},
		{
		  "internalType": "address",
		  "name": "_originSender",
		  "type": "address"
		},
		{
		  "internalType": "address",
		  "name": "_recipient",
		  "type": "address"
		},
		{
		  "internalType": "address",
		  "name": "_rollbackAccount",
		  "type": "address"
		},
		{
		  "internalType": "uint256",
		  "name": "_amount",
		  "type": "uint256"
		},
		{
		  "internalType": "uint256",
		  "name": "_nonce",
		  "type": "uint256"
		},
		{
		  "internalType": "bytes",
		  "name": "_extData",
		  "type": "bytes"
		}
	  ],
	  "name": "unlock",
	  "outputs": [],
	  "stateMutability": "nonpayable",
	  "type": "function"
	}
]`

const WTOKEN_CONVERTOR = `[
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_remoteChainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "_recipient",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_rollbackAccount",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "_amount",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "_nonce",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "_extData",
                "type": "bytes"
            },
            {
                "internalType": "bytes",
                "name": "_extParams",
                "type": "bytes"
            }
        ],
        "name": "lockAndXIssue",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    }
]`

const XTOKEN_CONVERTOR = `[{
    "inputs": [
      {
        "internalType": "address",
        "name": "_recipient",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_rollbackAccount",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_nonce",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "_extData",
        "type": "bytes"
      },
      {
        "internalType": "bytes",
        "name": "_extParams",
        "type": "bytes"
      }
    ],
    "name": "burnAndXUnlock",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
}]`

const ERC20_ABI = `[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "_spender",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "value",
        "type": "uint256"
      }
    ],
	"outputs": [
		{
		  "internalType": "bool",
		  "name": "success",
		  "type": "bool"
		}
	  ],
    "name": "approve",
    "type": "function"
  },
	{
	  "inputs": [
		{
		  "internalType": "address",
		  "name": "owner",
		  "type": "address"
		},
		{
		  "internalType": "address",
		  "name": "spender",
		  "type": "address"
		}
	  ],
	  "name": "allowance",
	  "outputs": [
		{
		  "internalType": "uint256",
		  "name": "",
		  "type": "uint256"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	}

]`
