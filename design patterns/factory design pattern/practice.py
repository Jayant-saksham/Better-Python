"""
Without Factory Method
"""

class SavingsAccount:
    def __init__(self, holder_name):
        self.holder_name = holder_name
        self.account_type = "Savings"

    def get_details(self):
        return f"{self.holder_name} has a {self.account_type} account."


class CurrentAccount:
    def __init__(self, holder_name):
        self.holder_name = holder_name
        self.account_type = "Current"

    def get_details(self):
        return f"{self.holder_name} has a {self.account_type} account."


# Without Factory Pattern, we create objects directly
savings = SavingsAccount("Alice")
current = CurrentAccount("Bob")

print(savings.get_details())
print(current.get_details())


"""
With Factory Method
"""


class SavingsAccount:
    def __init__(self, holder_name):
        self.holder_name = holder_name
        self.account_type = "Savings"

    def get_details(self):
        return f"{self.holder_name} has a {self.account_type} account."


class CurrentAccount:
    def __init__(self, holder_name):
        self.holder_name = holder_name
        self.account_type = "Current"

    def get_details(self):
        return f"{self.holder_name} has a {self.account_type} account."

class AccountExporter:
    def get_saving_exporter(self, holder_name):
        return SavingsAccount(holder_name)

    def get_current_exporter(self, holder_name):
        return CurrentAccount(holder_name)



# Without Factory Pattern, we create objects directly
savings = AccountExporter().get_saving_exporter("Alice")
current =AccountExporter().get_current_exporter("Bob")

print(savings.get_details())
print(current.get_details())
