// Helper function to get customer ID from a form input
function getCustomerId() {
    return document.getElementById('customerId').value;
}

function createTransaction() {
    const customerName = document.getElementById('customerName').value;
    const customerEmail = document.getElementById('customerEmail').value;
    const cartItems = JSON.parse(document.getElementById('cartItems').value);
    const customerId = getCustomerId();

    const data = {
        customer: {
            id: customerId,
            name: customerName,
            email: customerEmail
        },
        cartItems: cartItems
    };

    fetch('http://localhost:8080/create-transaction', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('transactionID').value = data.transactionID;
        document.getElementById('message').textContent = 'Transaction created successfully!';
    })
    .catch(error => {
        document.getElementById('message').textContent = 'Error creating transaction: ' + error;
    });
}

function processPayment() {
    const transactionID = document.getElementById('transactionID').value;
    const paymentForm = JSON.parse(document.getElementById('paymentForm').value);
    const customerId = getCustomerId();

    const data = {
        transactionID: transactionID,
        paymentForm: paymentForm
    };

    fetch('http://localhost:8080/process-payment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('message').textContent = 'Payment processed successfully!';
    })
    .catch(error => {
        document.getElementById('message').textContent = 'Error processing payment: ' + error;
    });
}

function getCustomerTransactions() {
    const customerId = getCustomerId();

    fetch(`http://localhost:8080/transactions?customerId=${customerId}`)
    .then(response => response.json())
    .then(transactions => {
        const tbody = document.getElementById('transactionsTable').querySelector('tbody');
        tbody.innerHTML = '';

        transactions.forEach(transaction => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${transaction.id}</td>
                <td>${transaction.customerName}</td>
                <td>${transaction.customerEmail}</td>
                <td>${transaction.status}</td>
            `;
            tbody.appendChild(row);
        });
    })
    .catch(error => {
        document.getElementById('message').textContent = 'Error retrieving transactions: ' + error;
    });
}

document.addEventListener('DOMContentLoaded', function () {
    // Function to create a new transaction
    window.createTransaction = function() {
        const customerId = document.getElementById('customerId').value;
        const customerName = document.getElementById('customerName').value;
        const customerEmail = document.getElementById('customerEmail').value;
        const cartItems = JSON.parse(document.getElementById('cartItems').value);

        const data = {
            customer: {
                id: customerId,
                name: customerName,
                email: customerEmail
            },
            cartItems: cartItems
        };

        fetch('/create-transaction', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => response.json())
        .then(data => {
            document.getElementById('message').innerText = 'Transaction created with ID: ' + data.transactionID;
        })
        .catch(error => {
            console.error('Error creating transaction:', error);
            document.getElementById('message').innerText = 'Error creating transaction: ' + error;
        });
    };

    // Function to process a payment
    window.processPayment = function() {
        const transactionID = document.getElementById('transactionID').value;
        const customerId = document.getElementById('customerId').value;
        const paymentForm = JSON.parse(document.getElementById('paymentForm').value);

        const data = {
            transactionID: transactionID,
            customerId: customerId,
            paymentForm: paymentForm
        };

        fetch('/process-payment', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => response.json())
        .then(data => {
            document.getElementById('message').innerText = 'Payment processed successfully';
        })
        .catch(error => {
            console.error('Error processing payment:', error);
            document.getElementById('message').innerText = 'Error processing payment: ' + error;
        });
    };

    // Function to get customer transactions
    window.getCustomerTransactions = function() {
        const customerId = document.getElementById('customerId').value;

        fetch('/transactions/' + customerId)
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('transactionsTable').getElementsByTagName('tbody')[0];
            tbody.innerHTML = ''; // Clear existing rows

            data.transactions.forEach(transaction => {
                const row = tbody.insertRow();
                row.insertCell(0).innerText = transaction.id;
                row.insertCell(1).innerText = transaction.customerName;
                row.insertCell(2).innerText = transaction.customerEmail;
                row.insertCell(3).innerText = transaction.status;
            });
        })
        .catch(error => {
            console.error('Error fetching transactions:', error);
            document.getElementById('message').innerText = 'Error fetching transactions: ' + error;
        });
    };
});
