"""
Imagine you are creating an application that shows the data about all different types of vehicles present.
It takes the data from APIs of different vehicle organizations in XML format and then displays the information.
But suppose at some time you want to upgrade your application with a Machine Learning algorithms
that work beautifully on the data and fetch the important data only. But there is a problem, it takes data in
JSON format only. It will be a really poor approach to make changes in Machine Learning Algorithm so that
it will take data in XML format.
"""

from abc import ABC, abstractmethod

class Vehicle(ABC):

    @abstractmethod
    def  get_data(self):
        pass

class Truck(Vehicle):
    def get_data(self):
        return "XML truck data"

class Scotter(Vehicle):
    def get_data(self):
        return "XML Scotter data"

class Bike(Vehicle):
    def get_data(self):
        return "XML Bike data"


class XMLToJSONAdapter:
    def get_json(self, xml):
        return f"JSON Data from XML {xml}"

class MachineLearning:
    def model_train(self, json_data):
        truck = Truck()
        truck.get_data()
        print(f"using {json_data} for ML")


if __name__ == "__main__":
    truck = Truck()
    xml_data = truck.get_data()
    json_data = XMLToJSONAdapter().get_json(xml_data)
    ml = MachineLearning()
    ml.model_train(json_data)

