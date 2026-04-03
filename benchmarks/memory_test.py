import psutil
import subprocess
import time
import json
import os
import signal
import sys


def get_memory_usage(pid):
    try:
        proc = psutil.Process(pid)
        return proc.memory_info().rss / 1024 / 1024  # MB
    except:
        return None


def start_service(service_type):
    if service_type == "gin":
        return subprocess.Popen(
            ["go", "run", "main.go"],
            cwd=os.path.join(os.path.dirname(__file__), "..", "go-gin-service"),
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL,
            creationflags=subprocess.CREATE_NEW_PROCESS_GROUP if sys.platform == "win32" else 0
        )
    elif service_type == "fastapi":
        return subprocess.Popen(
            [sys.executable, "-m", "uvicorn", "main:app", "--port", "8001"],
            cwd=os.path.join(os.path.dirname(__file__), "..", "fastapi-service"),
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL,
            creationflags=subprocess.CREATE_NEW_PROCESS_GROUP if sys.platform == "win32" else 0
        )


def run_load_test():
    """Запускает ab нагрузочный тест"""
    try:
        ab_path = "C:/Apache24/bin/ab.exe"
        result = subprocess.run(
            [ab_path, "-n", "500", "-c", "10", "http://localhost:8080/ping"],
            capture_output=True, text=True, timeout=30
        )
        return result.returncode == 0
    except:
        return False


def main():
    results = {}

    for service in ["gin", "fastapi"]:
        print(f"\nТестируем {service}...")

        # Запускаем сервис
        proc = start_service(service)
        time.sleep(3)  # Даём время на запуск

        # Замер памяти до нагрузки
        mem_before = get_memory_usage(proc.pid)
        time.sleep(1)

        # Запускаем нагрузку (только для Gin, для FastAPI отдельно)
        port = 8080 if service == "gin" else 8000
        ab_path = "C:/Apache24/bin/ab.exe"
        subprocess.run([ab_path, "-n", "200", "-c", "10", f"http://localhost:{port}/ping"],
                       capture_output=True, timeout=30)

        # Замер памяти после нагрузки
        time.sleep(2)
        mem_after = get_memory_usage(proc.pid)

        # Останавливаем сервис
        if sys.platform == "win32":
            proc.terminate()
        else:
            proc.send_signal(signal.SIGTERM)
        time.sleep(2)

        results[service] = {
            "memory_before_mb": mem_before,
            "memory_after_mb": mem_after,
            "delta_mb": mem_after - mem_before if mem_before and mem_after else None
        }

        print(f"  Память до: {mem_before:.2f} MB" if mem_before else "  Память до: N/A")
        print(f"  Память после: {mem_after:.2f} MB" if mem_after else "  Память после: N/A")

    # Сохраняем результаты
    with open("memory_benchmark_results.json", "w") as f:
        json.dump(results, f, indent=2)

    print("\nРезультаты сохранены в memory_benchmark_results.json")


if __name__ == "__main__":
    main()