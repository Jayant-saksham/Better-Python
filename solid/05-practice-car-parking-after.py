"""
    Design a car parking application

    There are two categories of cars: SUV and hatchback.
    Maintain a count of how many SUV and hatchback cars enter the premise
    Calculate the payment each car has to make based upon the rates as hatchback parking as 10 rupees per hour and SUV being 20 rupees an hour.
    In case if hatchback occupancy is full then hatchback cars can occupy SUV spaces but with hatchback rates.
    During exit there needs to be the system to inform the user how much they have to pay
    Admin can see all the cars which are parked in the system
"""
import time
from abc import ABC, abstractmethod
from datetime import datetime
from random import randint

ledger = {}

class Car:

    @abstractmethod
    def park(self, car_id):
        pass

    @abstractmethod
    def unpark(self, car_id):
        pass

    @abstractmethod
    def pay(self, car_id):
        pass

    @abstractmethod
    def validate_park(self, car_id):
        pass

    # Concreate methods
    def validate_unpark(self, count, car_id, car_type):
        if count <= 0:
            raise Exception(f"No {car_type} Cars Found in parking space")

        if car_id not in ledger:
            raise Exception(f"No {car_type} Cars Found in ledger")

    def admin_display(self):
        print(ledger)

class SUVCar(Car):

    def __init__(self, total_space, rate = 20):
        self.count = 0
        self.total_space = total_space
        self.rate = rate

    def validate_park(self, car_id):
        global ledger
        if self.total_space == 0:
            raise Exception("SUV Space occupied")

        if car_id in ledger:
            raise Exception("SUV Car of this id already exist")

    def park(self, car_id):
        global ledger
        self.validate_park(car_id)

        current_time = datetime.now()
        ledger.update({
            car_id: {"time_in": current_time, "parking_lot": "SUV"}
        })
        self.total_space -=1
        self.count +=1

    def unpark(self, car_id):
        global ledger
        self.validate_unpark(self.count, car_id, "SUV")
        self.count -=1
        self.total_space += 1
        self.pay(car_id)
        del ledger[car_id]

    def pay(self, car_id):
        time_out = datetime.now()
        time_in = ledger[car_id]["time_in"]
        price = ((time_out - time_in).total_seconds() / 3600) * 20
        print(price)

class HatchBackCar(Car):
    def __init__(self, total_space, rate):
        self.count = 0
        self.total_space = total_space
        self.rate = rate
        self.suv = SUVCar(total_space)

    def validate_park(self, car_id):
        if self.total_space == 0 and self.suv.total_space == 0:
            raise Exception("HatchBack Space occupied")

        if car_id in ledger:
            raise Exception("HatchBack Car of this id already exist")

    def park(self, car_id):
        global ledger
        self.validate_park(car_id)

        current_time = datetime.now()
        if self.total_space == 0:
            ledger.update({
                car_id: {"time_in": current_time, "parking_lot": "SUV"}
            })
            self.suv.total_space -= 1
        else:
            ledger.update({
                car_id: {"time_in": current_time, "parking_lot": "HatchBack"}
            })
            self.total_space -= 1

        self.count += 1

    def unpark(self, car_id):
        global ledger
        self.validate_unpark(self.count, car_id, "HatchBack")
        self.count -= 1
        parking_lot = ledger[car_id]["parking_lot"]
        if parking_lot == "SUV":
            self.suv.total_space += 1
        else:
            self.total_space += 1

        self.pay(car_id)
        del ledger[car_id]

    def pay(self, car_id):
        time_out = datetime.now()
        time_in = ledger[car_id]["time_in"]
        price = ((time_out - time_in).total_seconds() / 3600) * 10
        print(price)

if __name__ == "__main__":
    hb1 = HatchBackCar(3, 10)
    suv1 = SUVCar(3, 20)
    rndInt = randint(1, 10)

    hb1.park(rndInt * "d")
    suv1.unpark(rndInt * "d")



