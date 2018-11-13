# stellarTools
Tools for the Stellar

The samrt contract of the Stellar CreateAccounts.go which bases on the testnet that allows you to create two accounts
and build up the trustline with the MXC token, also create a master key for the second account automatically.

Then you can use CheckBalance.go to check the balance and if it builds up the trulst line with MXC tokcn successfully.

After you created the accounts, try to get some MXC token by GetSomeMXC.go to do the test.

SendToAccount2 allows you to send some token to you second account.

If you run the EmergencyFunction.go, it will send all your MXC tokens to a security account,
this tool is for protect the token been stolen by the hackers.

WebResponse.go is a server that can response 0 or 1, it works with Tx_SecAcc.go,
so when second account tries to send the money, if the response is 0 then the account is able to send the tokens,
in contrast, if the response is 1, then the transaction will be rejected.
