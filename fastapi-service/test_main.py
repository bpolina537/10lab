import pytest
from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

def test_ping_endpoint():
    response = client.get("/ping")
    assert response.status_code == 200
    assert response.json() == {"message": "pong from FastAPI"}

def test_call_go_success():
    response = client.post("/call-go", json={"data": "hello"})
    assert response.status_code in [200, 503]

def test_call_go_empty_data():
    response = client.post("/call-go", json={})
    assert response.status_code in [200, 503]

def test_call_go_none_data():
    response = client.post("/call-go", json={"data": None})
    assert response.status_code in [200, 503]

def test_call_go_invalid_json():
    response = client.post("/call-go", data="invalid")
    assert response.status_code == 422

def test_call_go_unicode_data():
    response = client.post("/call-go", json={"data": "привет мир"})
    assert response.status_code in [200, 503]

def test_call_go_large_data():
    large_data = "x" * 10000
    response = client.post("/call-go", json={"data": large_data})
    assert response.status_code in [200, 503]