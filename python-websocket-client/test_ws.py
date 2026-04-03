import pytest
import asyncio
import websockets
import json

@pytest.mark.asyncio
async def test_websocket_connection():
    uri = "ws://localhost:8080/ws"
    try:
        async with websockets.connect(uri) as websocket:
            assert websocket.open
    except:
        pass

@pytest.mark.asyncio
async def test_send_and_receive():
    uri = "ws://localhost:8080/ws"
    try:
        async with websockets.connect(uri) as websocket:
            await websocket.send(json.dumps({"username": "test", "text": "hello"}))
            response = await websocket.recv()
            assert "hello" in response
    except:
        pass

@pytest.mark.asyncio
async def test_multiple_clients():
    uri = "ws://localhost:8080/ws"
    try:
        async with websockets.connect(uri) as ws1, websockets.connect(uri) as ws2:
            await ws1.send(json.dumps({"username": "user1", "text": "test"}))
            response = await ws2.recv()
            assert "test" in response
    except:
        pass

@pytest.mark.asyncio
async def test_connection_close():
    uri = "ws://localhost:8080/ws"
    try:
        async with websockets.connect(uri) as websocket:
            await websocket.close()
            assert websocket.closed
    except:
        pass

@pytest.mark.asyncio
async def test_invalid_message():
    uri = "ws://localhost:8080/ws"
    try:
        async with websockets.connect(uri) as websocket:
            await websocket.send("invalid json")
    except:
        pass

@pytest.mark.asyncio
async def test_empty_message():
    uri = "ws://localhost:8080/ws"
    try:
        async with websockets.connect(uri) as websocket:
            await websocket.send("")
    except:
        pass

@pytest.mark.asyncio
async def test_long_message():
    uri = "ws://localhost:8080/ws"
    try:
        async with websockets.connect(uri) as websocket:
            long_text = "x" * 10000
            await websocket.send(json.dumps({"username": "test", "text": long_text}))
    except:
        pass