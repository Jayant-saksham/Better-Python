from abc import ABC, abstractmethod

class Order:

    def __init__(self):
        self.items = {}

    def add_items(self, name, quantity, price, id):
        self.items[id] = {
            "name": name,
            "quantity": quantity,
            "price": price,
        }

    def total_price(self):
        total = 0
        for i in self.items.items():
            key, value = i
            total +=(value['price'] * value['quantity'])
        return total

class PaymentProcessor(ABC):

    @abstractmethod
    def pay(self, order):
        pass

class PaymentProcessorSMS(PaymentProcessor):
    @abstractmethod
    def auth_sms(self, code):
        pass


class DebitPaymentProcessor(PaymentProcessorSMS):
    def __init__(self, security_code):
        self.security_code = security_code

    def auth_sms(self, code):
        print("Debit Card Payment SMS verfied")

    def pay(self,order):
        print("Processing debit payment", self.security_code)

class CreditPaymentProcessor(PaymentProcessor):
    def pay(self,order):
        print("Processing credit payment")

class WalletPaymentProcessor(PaymentProcessorSMS):
    def __init__(self, email):
        self.email = email

    def auth_sms(self, code):
        print("Wallet Payment Auth SMS verified")

    def pay(self,order):
        print("Processing wallet payment", self.email)



order = Order()
order.add_items("Keyboard", 1, 100, "123")
order.add_items("Mouse", 2, 50, "1234")
order.add_items("Speakers", 3, 150, "1235")

print(order.total_price())
debit = DebitPaymentProcessor("12345")
debit.pay(order)


wallet = WalletPaymentProcessor("jayant2410@gmail.com")
wallet.auth_sms("code")
wallet.pay(order)

