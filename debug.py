import time

class Debug:
    def __init__(self, enabled: bool, window_width: int, config: dict):
        self.enabled = enabled
        self.window_width = window_width
        self.start_time = time.time()
        self.conf = config

    def get_info(self, msg_len: int) -> str:
        duration = int((time.time() - self.start_time) * 1000000) 

        debug_msg = f"{duration} μs | {self.conf}"

        shifts = self.window_width - (msg_len % self.window_width) - len(debug_msg)

        return " " * shifts + debug_msg 
