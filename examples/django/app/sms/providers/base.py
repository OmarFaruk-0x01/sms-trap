from abc import ABC, abstractmethod

class SmsProvider(ABC):
    @abstractmethod
    def send(self, phones: list[str], message: str, additional: dict) -> dict:
        pass