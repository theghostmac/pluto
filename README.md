# Pluto is a blockchain implementation in Go.
What is a blockchain? [Here](https://en.wikipedia.org/wiki/Blockchain).

![Pluto is active](bootloaderNodeWorks.png)
I asked ChatGPT what is happening because I don't actually understand. Response:
```markdown
The output you provided indicates that the program is running and generating log messages. The repeated log messages you see are likely due to the loop that simulates sending messages between your node1 and node2 transport nodes. These log messages show that the same message is being sent repeatedly from node2 to node1.

Additionally, you have implemented a loop that interacts with the server and the blockchain every 500 milliseconds. This loop repeatedly sends a message to the server, which in turn processes the messages and logs their content.
```

![Tests Pass](testspass.png)

