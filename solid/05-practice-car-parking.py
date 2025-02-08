"""
    Design a car parking application

    There are two categories of cars: SUV and hatchback.
    Maintain a count of how many SUV and hatchback cars enter the premise
    Calculate the payment each car has to make based upon the rates as hatchback parking as 10 rupees per hour and SUV being 20 rupees an hour.
    In case if hatchback occupancy is full then hatchback cars can occupy SUV spaces but with hatchback rates.
    During exit there needs to be the system to inform the user how much they have to pay
    Admin can see all the cars which are parked in the system
"""
import datetime

class Car:
    hatchBackCount : int
    SUVCount : int
    InitialCapSUV: int
    InitialCapHatchBack: int

    def __init__(self):
        self.SUVCount = 0
        self.hatchBackCount = 0
        self.InitialCapHatchBack = 3
        self.InitialCapSUV = 3
        self.cars = {}

    def admin_display(self):
        print("Total SUV", self.SUVCount )
        print("Total HatchBack", self.hatchBackCount)


    def park(self, car_type, car_id):
        if car_type == "SUV":
            if self.InitialCapSUV <= 0:
                raise Exception("Space occupied")

            self.SUVCount += 1
            self.InitialCapSUV -= 1
            self.cars[car_id] = {"time": datetime.datetime.now(), "car_type": car_type}

        elif car_type == "HatchBack":
            if self.InitialCapHatchBack <= 0:
                if self.InitialCapSUV <= 0:
                    raise Exception("Space occupied SUV and HatchBack")

                self.InitialCapSUV -= 1

            else:
                self.InitialCapHatchBack -= 1
            self.hatchBackCount += 1
            self.cars[car_id] = {"time": datetime.datetime.now(), "car_type": car_type}

    def unpark(self, car_type, car_id):
        if car_id in self.cars:

            if car_type == "SUV":
                if self.SUVCount < 1:
                    raise Exception("There is no car parked")
                else:
                    self.SUVCount -= 1
                    self.InitialCapSUV += 1

                    start_time = self.cars[car_id]['time']
                    end_time = datetime.datetime.now()
                    time_difference = ((end_time - start_time).total_seconds())
                    self.cars[car_id]['time_difference'] = time_difference
                    self.pay(car_type, car_id)

            elif car_type == "HatchBack":
                if self.hatchBackCount < 1:
                    raise Exception("There is no car parked")
                else:
                     if self.InitialCapHatchBack <= 0:
                         self.InitialCapSUV += 1
                     else:
                         self.InitialCapHatchBack += 1

                     self.hatchBackCount -= 1


                     start_time = self.cars[car_id]['time']
                     end_time = datetime.datetime.now()
                     time_difference = ((end_time - start_time).total_seconds())
                     self.cars[car_id]['time_difference'] = time_difference
                     self.pay(car_type, car_id)

    def pay(self, car_type, car_id):
        if car_type == "SUV":
            print((self.cars[car_id]['time_difference'] / 60) * 20)
        else:
            print((self.cars[car_id]['time_difference'] / 60) * 10)

    def get_count_suv(self):
        return self.SUVCount

    def get_count_hatchback(self):
        return self.hatchBackCount


c = Car()
# c.park("SUV", "123")
# c.park("SUV", "123")
# c.park("SUV", "123")
# c.park("HatchBack", "321")
# c.park("HatchBack", "321")
# c.park("HatchBack", "321")
c.unpark("SUV", "123")
c.unpark("SUV", "123")
c.unpark("SUV", "123")
c.admin_display()