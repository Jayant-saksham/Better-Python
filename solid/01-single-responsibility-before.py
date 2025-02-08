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

    def pay(self, payment_type, security_code):
        if payment_type == "debit":
            print("Processing debit payment", security_code)
        elif payment_type == "credit":
            print("Processing credit payment", security_code)
        else:
            raise Exception(f"Unknown payment type: {payment_type}")

order = Order()
order.add_items("Keyboard", 1, 100, "123")
order.add_items("Mouse", 2, 50, "1234")
order.add_items("Speakers", 3, 150, "1235")

print(order.total_price())
order.pay("debit", "jjdhh134r5tyhbgvfc")

