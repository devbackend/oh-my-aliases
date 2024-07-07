import time

class Debug:
    def __init__(self, enabled: bool):
        self.enabled = enabled
        self.start_time = time.time()

    def get_info(self) -> str:
        duration = int((time.time() - self.start_time) * 1000000) 

        return f"(took: {duration} μs)"
