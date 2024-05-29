document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('createTransactionBtn').addEventListener('click', createTransaction);
    document.getElementById('addItemBtn').addEventListener('click', addItem);
    document.getElementById('processPaymentBtn')?.addEventListener('click', processPayment);
    document.getElementById('getTransactionsBtn')?.addEventListener('click', getCustomerTransactions);
});

function addItem() {
    const cartItemsContainer = document.getElementById('cartItems');
    const cartItemTemplate = document.getElementById('cartItemTemplate').content.cloneNode(true);
    cartItemsContainer.appendChild(cartItemTemplate);
}

function createTransaction() {
    const customerId = document.getElementById('customerId').value;
    const customerName = document.getElementById('customerName').value;
    const customerEmail = document.getElementById('customerEmail').value;

    const cartItems = [];
    document.querySelectorAll('.cart-item').forEach(item => {
        cartItems.push({
            id: item.querySelector('.itemID').value,
            name: item.querySelector('.itemName').value,
            price: parseFloat(item.querySelector('.itemPrice').value),
            quantity: parseFloat(item.querySelector('.itemQuantity').value)
        });
    });

    fetch('/create-transaction', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            customer: { id: customerId, name: customerName, email: customerEmail },
            cartItems: cartItems
        })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('message').innerText = `Transaction created with ID: ${data.transactionID}`;
        // Clear the form
        document.getElementById('transactionForm').reset();
        document.getElementById('cartItems').innerHTML = '';
    })
    .catch(error => {
        document.getElementById('message').innerText = `Error creating transaction: ${error}`;
    });
}

function processPayment() {
    const transactionID = document.getElementById('transactionID').value;
    const cardNumber = document.getElementById('cardNumber').value;
    const expirationDate = document.getElementById('expirationDate').value;
    const cvv = document.getElementById('cvv').value;
    const name = document.getElementById('name').value;
    const address = document.getElementById('address').value;

    fetch('/process-payment', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            transactionID: transactionID,
            paymentForm: {
                cardNumber: cardNumber,
                expirationDate: expirationDate,
                cvv: cvv,
                name: name,
                address: address
            }
        })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('message').innerText = `Payment status: ${data.status}`;
    })
    .catch(error => {
        document.getElementById('message').innerText = `Error processing payment: ${error}`;
    });
}

function getCustomerTransactions() {
    const customerId = document.getElementById('customerId').value;

    fetch(`/transactions/${customerId}`)
    .then(response => response.json())
    .then(data => {
        const tableBody = document.getElementById('transactionsTable').getElementsByTagName('tbody')[0];
        tableBody.innerHTML = '';
        if (data.transactions && data.transactions.length > 0) {
            data.transactions.forEach(transaction => {
                const row = tableBody.insertRow();
                row.insertCell(0).innerText = transaction.id;
                row.insertCell(1).innerText = transaction.customerName;
                row.insertCell(2).innerText = transaction.customerEmail;
                row.insertCell(3).innerText = transaction.status;
            });
        } else {
            document.getElementById('message').innerText = 'No transactions found for this customer.';
        }
    })
    .catch(error => {
        document.getElementById('message').innerText = `Error fetching transactions: ${error}`;
    });
}
