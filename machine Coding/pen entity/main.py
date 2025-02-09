# Design a pen
# Pen have its name, brand, price etc
# 1. Pen can be of different types: Ball, Gel, Fountain
# 2. Ink can be of different colors
# 3. Body of pen could be different
# 4. Nib size could be different
# 5. Every refill will have nib
# some ball and gel pen cab be refilled and some cant
from abc import ABC, abstractmethod


class Pen(ABC):
    name: str
    brand: str
    price: float

    @abstractmethod
    def write(self):
        pass

# Plastic refillable and not direct ink filling
class Refillable(ABC):

    @abstractmethod
    def refill(self):
        pass

class Ink:
    def __init__(self, color: str, features: list):
        self._color = color
        self._features = features

class Nib:
    def __init__(self, radius: float):
        self._radius = radius


class BallPenNonReUsable(Pen):

    def __init__(self, name, brand, price):
        self._name = name
        self._brand = brand
        self._price = price


    def write(self):
        print("Writing with ball pen which is non-reusable")

class BallPenReUsable(Pen, Refillable):

    def __init__(self, name, brand, price):
        self._name = name
        self._brand = brand
        self._price = price

    def write(self):
        print("Writing with ball pen which is re-usable")

    def refill(self, Ink, Nib):
        print(f"Refill Ball pen with ink as {Ink._color} and nib as {Nib._radius} with features {Ink._features} ")


class GelPenNonReUsable(Pen):

    def __init__(self, name, brand, price):
        self._name = name
        self._brand = brand
        self._price = price

    def write(self):
        print("Writing with gel pen which is non-reusable")

class GelPenReUsable(Pen, Refillable):

    def __init__(self, name, brand, price):
        self._name = name
        self._brand = brand
        self._price = price

    def write(self):
        print("Writing with gel pen which is re-usable")

    def refill(self):
        print(f"Refill Ball pen with ink as {Ink} and nib as {Nib} with features {Ink._features}")


class FountainPen(Pen):

    def __init__(self, name, brand, price):
        self._name = name
        self._brand = brand
        self._price = price

    def write(self):
        print("Writing with fountain pen")

if __name__ == "__main__":
    pass
