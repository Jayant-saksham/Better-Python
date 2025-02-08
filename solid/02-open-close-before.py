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

class PaymentProcessor:
    def pay_debit(self, order, security_code):
        print("Processing debit payment", security_code)

    def pay_credit(self, order, security_code):
        print("Processing credit payment", security_code)



order = Order()
order.add_items("Keyboard", 1, 100, "123")
order.add_items("Mouse", 2, 50, "1234")
order.add_items("Speakers", 3, 150, "1235")

print(order.total_price())
processor = PaymentProcessor()
processor.pay_debit(order, "123456")

