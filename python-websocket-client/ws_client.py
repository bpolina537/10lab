import asyncio
import websockets

async def chat():
    uri = "ws://localhost:8080/ws"
    async with websockets.connect(uri) as websocket:
        username = input("Enter your name: ")
        print("Connected to chat. Type messages below:")

        async def receive():
            async for message in websocket:
                print(f"\n{message}")

        async def send():
            while True:
                text = await asyncio.get_event_loop().run_in_executor(None, input)
                await websocket.send(f'{{"username":"{username}","text":"{text}"}}')

        await asyncio.gather(receive(), send())

if __name__ == "__main__":
    asyncio.run(chat())